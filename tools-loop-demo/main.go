package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
	"github.com/openai/openai-go/v2/shared"
	"github.com/openai/openai-go/v2/shared/constant"
)

func main() {
	// Step 1: Initialize context for request management
	ctx := context.Background()

	// ======================================================
	// MCP Client
	// ======================================================
	mcpClient, err := client.NewStreamableHttpClient(
		"http://localhost:9011/mcp",
	)
	//defer mcpClient.Close()
	if err != nil {
		panic(err)
	}
	// Start the connection to the server
	err = mcpClient.Start(ctx)
	if err != nil {
		panic(err)
	}

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "micro agent",
		Version: "0.0.0",
	}
	_, err = mcpClient.Initialize(ctx, initRequest)
	if err != nil {
		panic(err)
	}
	//fmt.Println("Streamable HTTP client connected & initialized with server!", result)
	//ui.Println(ui.Yellow, "Streamable HTTP client connected & initialized with server!")

	toolsRequest := mcp.ListToolsRequest{}
	mcpTools, err := mcpClient.ListTools(ctx, toolsRequest)
	if err != nil {
		panic(err)
	}

	for _, tool := range mcpTools.Tools {
		fmt.Printf("Tool: %s - %s - %s\n", tool.Name, tool.Description, tool.InputSchema)
	}

	openAITools := ConvertMCPToolsToOpenAITools(mcpTools)

	// ======================================================
	// DMR Client
	// ======================================================
	chatURL := "http://localhost:12434/engines/llama.cpp/v1/"
	model := "hf.co/menlo/jan-nano-gguf:q4_k_m"

	client := openai.NewClient(
		option.WithBaseURL(chatURL),
		option.WithAPIKey(""), // No API key needed for local deployment
	)

	// This will trigger multiple sequential tool calls from the AI
	userQuestion := openai.UserMessage(`
		Find rust snippet about error handling.
		Find go snippet about structure.
	`)

	// Initialize loop control variables
	stopped := false           // Controls the conversation loop
	finishReason := ""         // Tracks why AI stopped responding
	results := []string{}      // Stores tool execution results
	lastAssistantMessage := "" // Final AI message

	// Initialize conversation history with user's question
	messages := []openai.ChatCompletionMessageParamUnion{
		userQuestion,
	}

	// Configure chat completion parameters
	// ParallelToolCalls set to false for sequential execution
	params := openai.ChatCompletionNewParams{
		ParallelToolCalls: openai.Bool(false), // Execute tools one by one
		Tools:             openAITools,        // Available tools for AI
		Model:             model,
		Temperature:       openai.Opt(0.0), // Deterministic responses
	}

	// Main conversation loop - continues until AI says "stop"
	for !stopped {

		// Update parameters with current conversation history
		params.Messages = messages

		// Send request to AI model and get response
		completion, err := client.Chat.Completions.New(ctx, params)
		if err != nil {
			panic(err)
		}
		// Extract finish reason to determine next action
		finishReason = completion.Choices[0].FinishReason

		// Step 15: Handle AI response based on finish reason
		switch finishReason {
		case "tool_calls":
			// AI wants to use tools - extract tool calls
			detectedToolCalls := completion.Choices[0].Message.ToolCalls

			if len(detectedToolCalls) > 0 {

				// Convert tool calls to proper message format
				// WHY: When AI decides to use tools, it returns toolCalls in its response.
				// We must convert these into ChatCompletionMessageToolCallUnionParam format
				// to add them to conversation history. Without this conversion, AI would
				// lose context of what it requested.
				toolCallParams := make([]openai.ChatCompletionMessageToolCallUnionParam, len(detectedToolCalls))

				for i, toolCall := range detectedToolCalls {
					toolCallParams[i] = openai.ChatCompletionMessageToolCallUnionParam{
						OfFunction: &openai.ChatCompletionMessageFunctionToolCallParam{
							ID:   toolCall.ID,
							Type: constant.Function("function"),
							Function: openai.ChatCompletionMessageFunctionToolCallFunctionParam{
								Name:      toolCall.Function.Name,
								Arguments: toolCall.Function.Arguments,
							},
						},
					}
				}

				// Create assistant message with tool calls using proper union type
				// WHY: We need to create an "assistant" message containing the tool calls
				// for conversation history. This is like saying: "AI said: 'I want to call
				// these functions with these parameters'". This message will be added to
				// history before executing tools, so AI remembers what it requested.
				assistantMessage := openai.ChatCompletionMessageParamUnion{
					OfAssistant: &openai.ChatCompletionAssistantMessageParam{
						ToolCalls: toolCallParams,
					},
				}

				// Add the assistant message with tool calls to the conversation history
				messages = append(messages, assistantMessage)

				// Process each detected tool call sequentially
				for _, toolCall := range detectedToolCalls {
					functionName := toolCall.Function.Name
					functionArgs := toolCall.Function.Arguments

					// Execute the requested function
					fmt.Printf("▶️ Executing function: %s with args: %s\n", functionName, functionArgs)

					var args map[string]any
					args, _ = JsonStringToMap(functionArgs)

					// Call the MCP tool with the arguments
					request := mcp.CallToolRequest{}
					request.Params.Name = functionName
					request.Params.Arguments = args

					// Call the tool using the MCP client
					toolResponse, err := mcpClient.CallTool(ctx, request)
					resultContent := toolResponse.Content[0].(mcp.TextContent).Text

					// Handle function execution errors
					if err != nil {
						resultContent = fmt.Sprintf(`{"error": "Function execution failed: %s"}`, err)
					}

					// Store result for potential later use
					results = append(results, resultContent)

					// Add tool execution result to conversation history
					// WHY: After executing each tool, we must tell the AI what the result was.
					// This is like a conversation:
					// - AI: "I want to call sayHello with name='Jean-Luc'"
					// - System: "Result: 'Hello Jean-Luc'"
					// AI needs these results to: 1) Know tool executed successfully,
					// 2) Use results for final response, 3) Decide if more tools needed.
					// Without this step, AI would have no idea what happened after requesting
					// tool execution and couldn't generate the requested final report.
					messages = append(
						messages,
						openai.ToolMessage(
							resultContent,
							toolCall.ID,
						),
					)
					fmt.Println("✅ ResultContent", resultContent)
				}

			} else {
				// Step 25: Handle unexpected case with no tool calls
				fmt.Println("😢 No tool calls found in response")
			}

		case "stop":
			// AI has finished - no more tools needed
			fmt.Println("🟥 Stopping due to 'stop' finish reason.")
			stopped = true
			lastAssistantMessage = completion.Choices[0].Message.Content

			// Add final assistant message to conversation history
			messages = append(messages, openai.AssistantMessage(lastAssistantMessage))
			fmt.Print(strings.Repeat("=", 5), "[Last Assistant Message]", strings.Repeat("=", 51), "\n")
			fmt.Println(lastAssistantMessage)
			fmt.Println(strings.Repeat("=", 80))

		default:
			// Handle unexpected finish reasons
			fmt.Printf("🔴 Unexpected response: %s\n", finishReason)
			stopped = true

		}

	}

}

// JsonStringToMap converts a JSON string to a Go map
// Used to parse function arguments from AI tool calls
func JsonStringToMap(jsonString string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func ConvertMCPToolsToOpenAITools(tools *mcp.ListToolsResult) []openai.ChatCompletionToolUnionParam {
	openAITools := make([]openai.ChatCompletionToolUnionParam, len(tools.Tools))
	for i, tool := range tools.Tools {

		openAITools[i] = openai.ChatCompletionFunctionTool(shared.FunctionDefinitionParam{
			Name:        tool.Name,
			Description: openai.String(tool.Description),
			Parameters: shared.FunctionParameters{
				"type":       "object",
				"properties": tool.InputSchema.Properties,
				"required":   tool.InputSchema.Required,
			},
		},
		)
	}
	return openAITools
}
