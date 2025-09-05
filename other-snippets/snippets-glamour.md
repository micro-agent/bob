## Topics Covered [Glamour]
Examples and snippets for glamour markdown rendering
- Basic markdown rendering
- Built-in themes and styles
- Custom renderer configuration
- Word wrapping and width control
- Color profile settings
- Code block syntax highlighting
- Auto style detection
- Environment-based configuration
- Emoji rendering
- Table formatting
- File-based rendering
- Bubbletea integration

----------

## Available Functions
Main rendering functions and configuration
```go
// Basic rendering
glamour.Render(markdown, "dark")           // Render with theme
glamour.RenderBytes([]byte, "light")       // Render bytes
glamour.RenderWithEnvironmentConfig(text)  // Use env config

// Custom renderer
renderer, _ := glamour.NewTermRenderer(
    glamour.WithAutoStyle(),
    glamour.WithWordWrap(80),
    glamour.WithEmoji(),
)

// Style options
glamour.WithStandardStyle("dark")    // Built-in theme
glamour.WithAutoStyle()             // Auto detect theme
glamour.WithColorProfile(profile)   // Color settings
glamour.WithChromaFormatter("terminal256")
```

----------

## Basic Markdown Rendering
Simple markdown to terminal output
```go
package main

import (
    "fmt"
    "log"
    "github.com/charmbracelet/glamour"
)

func main() {
    markdown := `# Hello Glamour

This is **bold text** and this is *italic text*.

- List item 1
- List item 2
- List item 3`

    out, err := glamour.Render(markdown, "dark")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Print(out)
}
```

----------

## Different Themes
Using built-in themes
```go
package main

import (
    "fmt"
    "log"
    "github.com/charmbracelet/glamour"
)

func main() {
    markdown := `# Theme Examples

**Dark theme** looks great on dark terminals.
*Light theme* works well on light backgrounds.`

    themes := []string{"dark", "light", "notty", "ascii"}

    for _, theme := range themes {
        fmt.Printf("=== %s Theme ===\n", theme)
        out, err := glamour.Render(markdown, theme)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Print(out)
        fmt.Println()
    }
}
```

----------

## Custom Renderer
Configure renderer with options
```go
package main

import (
    "fmt"
    "log"
    "github.com/charmbracelet/glamour"
)

func main() {
    renderer, err := glamour.NewTermRenderer(
        glamour.WithAutoStyle(),
        glamour.WithWordWrap(80),
        glamour.WithEmoji(),
    )
    if err != nil {
        log.Fatal(err)
    }

    markdown := `# Custom Renderer :sparkles:

This text will be wrapped at 80 characters and emojis will be rendered!

Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.`

    out, err := renderer.Render(markdown)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Print(out)
}
```

----------

## Word Wrapping
Control text wrapping width
```go
package main

import (
    "fmt"
    "log"
    "github.com/charmbracelet/glamour"
)

func main() {
    longText := `# Word Wrapping Demo

This is a very long line of text that will demonstrate how word wrapping works with different width settings in the glamour markdown renderer.`

    widths := []int{40, 60, 80, 120}

    for _, width := range widths {
        fmt.Printf("=== Width: %d ===\n", width)
        
        renderer, err := glamour.NewTermRenderer(
            glamour.WithStandardStyle("dark"),
            glamour.WithWordWrap(width),
        )
        if err != nil {
            log.Fatal(err)
        }

        out, err := renderer.Render(longText)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Print(out)
        fmt.Println()
    }
}
```

----------

## Code Syntax Highlighting
Render code blocks with syntax highlighting
```go
package main

import (
    "fmt"
    "log"
    "github.com/charmbracelet/glamour"
)

func main() {
    markdown := `# Code Examples

## Go Code
` + "```go" + `
func main() {
    fmt.Println("Hello, World!")
    for i := 0; i < 10; i++ {
        fmt.Printf("Count: %d\n", i)
    }
}
` + "```" + `

## JavaScript Code
` + "```javascript" + `
const greeting = (name) => {
    return \`Hello, \${name}!\`;
};

console.log(greeting("World"));
` + "```" + `

## Python Code
` + "```python" + `
def fibonacci(n):
    if n <= 1:
        return n
    return fibonacci(n-1) + fibonacci(n-2)

print(fibonacci(10))
` + "```"

    renderer, err := glamour.NewTermRenderer(
        glamour.WithStandardStyle("dark"),
        glamour.WithChromaFormatter("terminal256"),
    )
    if err != nil {
        log.Fatal(err)
    }

    out, err := renderer.Render(markdown)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Print(out)
}
```

----------

## Auto Style Detection
Automatically detect terminal background
```go
package main

import (
    "fmt"
    "log"
    "github.com/charmbracelet/glamour"
)

func main() {
    markdown := `# Auto Style Detection

This will automatically choose:
- **Dark theme** for dark terminals
- **Light theme** for light terminals

No need to manually specify the theme!`

    renderer, err := glamour.NewTermRenderer(
        glamour.WithAutoStyle(),
    )
    if err != nil {
        log.Fatal(err)
    }

    out, err := renderer.Render(markdown)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Print(out)
}
```

----------

## Environment Configuration
Use environment variables for styling
```go
package main

import (
    "fmt"
    "log"
    "os"
    "github.com/charmbracelet/glamour"
)

func main() {
    // Set environment variable for glamour style
    os.Setenv("GLAMOUR_STYLE", "dark")

    markdown := `# Environment Config

This rendering uses the style defined in the GLAMOUR_STYLE environment variable.

**Current style:** dark`

    out, err := glamour.RenderWithEnvironmentConfig(markdown)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Print(out)
}
```

----------

## Tables and Lists
Render complex markdown structures
```go
package main

import (
    "fmt"
    "log"
    "github.com/charmbracelet/glamour"
)

func main() {
    markdown := `# Tables and Lists

## Shopping List
1. Apples :apple:
2. Bananas :banana:
3. Oranges :orange:

## Product Comparison

| Product | Price | Rating |
|---------|-------|--------|
| MacBook | $1299 | ⭐⭐⭐⭐⭐ |
| ThinkPad | $899  | ⭐⭐⭐⭐ |
| Surface | $1199 | ⭐⭐⭐⭐ |

## Nested Lists
- Fruits
  - Citrus
    - Orange
    - Lemon
  - Berries
    - Strawberry
    - Blueberry`

    renderer, err := glamour.NewTermRenderer(
        glamour.WithStandardStyle("dark"),
        glamour.WithEmoji(),
        glamour.WithWordWrap(80),
    )
    if err != nil {
        log.Fatal(err)
    }

    out, err := renderer.Render(markdown)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Print(out)
}
```

----------

## File Rendering
Render markdown from file
```go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "github.com/charmbracelet/glamour"
)

func main() {
    // Read markdown from file
    content, err := ioutil.ReadFile("README.md")
    if err != nil {
        log.Printf("Could not read file: %v", err)
        // Use sample content instead
        content = []byte(`# File Rendering Example
        
This demonstrates reading and rendering markdown from a file.

**Features:**
- File I/O handling
- Error management
- Fallback content`)
    }

    renderer, err := glamour.NewTermRenderer(
        glamour.WithStandardStyle("dark"),
        glamour.WithWordWrap(100),
    )
    if err != nil {
        log.Fatal(err)
    }

    out, err := renderer.RenderBytes(content)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Print(out)
}
```

----------

## Color Profile Settings
Configure color output capabilities
```go
package main

import (
    "fmt"
    "log"
    "github.com/charmbracelet/glamour"
)

func main() {
    markdown := `# Color Profile Demo

This text demonstrates different color profile settings:

**TrueColor:** Full 24-bit color support
**ANSI256:** 256-color palette
**ANSI:** Basic 16-color palette
**ASCII:** No colors, text only`

    profiles := []glamour.ColorProfile{
        glamour.TrueColor,
        glamour.ANSI256,
        glamour.ANSI,
        glamour.ASCII,
    }

    profileNames := []string{"TrueColor", "ANSI256", "ANSI", "ASCII"}

    for i, profile := range profiles {
        fmt.Printf("=== %s Profile ===\n", profileNames[i])
        
        renderer, err := glamour.NewTermRenderer(
            glamour.WithStandardStyle("dark"),
            glamour.WithColorProfile(profile),
        )
        if err != nil {
            log.Fatal(err)
        }

        out, err := renderer.Render(markdown)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Print(out)
        fmt.Println()
    }
}
```

----------

## Emoji Support
Enable emoji rendering in markdown
```go
package main

import (
    "fmt"
    "log"
    "github.com/charmbracelet/glamour"
)

func main() {
    markdown := `# Emoji Support :tada:

## Weather Today :sun_with_face:
- Sunny :sunny:
- Temperature: 75°F :thermometer:
- Wind: 5mph :dash:

## Development Status :computer:
- Tests passing :white_check_mark:
- Build successful :green_heart:
- Deploy ready :rocket:

## Team Mood :busts_in_silhouette:
- Alice :smile:
- Bob :thumbsup:
- Charlie :heart_eyes:`

    renderer, err := glamour.NewTermRenderer(
        glamour.WithStandardStyle("dark"),
        glamour.WithEmoji(),
        glamour.WithWordWrap(60),
    )
    if err != nil {
        log.Fatal(err)
    }

    out, err := renderer.Render(markdown)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Print(out)
}
```

----------

## Bubbletea Integration
Use glamour in bubbletea applications
```go
package main

import (
    "fmt"
    "strings"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/glamour"
)

type model struct {
    content  string
    renderer *glamour.TermRenderer
    err      error
}

func initialModel() model {
    renderer, err := glamour.NewTermRenderer(
        glamour.WithAutoStyle(),
        glamour.WithWordWrap(80),
    )

    markdown := `# Bubbletea + Glamour

This is a **bubbletea** application that renders markdown using **glamour**.

## Features
- Dynamic content rendering
- Responsive to terminal size
- Beautiful markdown display

Press 'q' to quit!`

    return model{
        content:  markdown,
        renderer: renderer,
        err:      err,
    }
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q":
            return m, tea.Quit
        }
    case tea.WindowSizeMsg:
        // Update renderer width when window resizes
        m.renderer, m.err = glamour.NewTermRenderer(
            glamour.WithAutoStyle(),
            glamour.WithWordWrap(msg.Width-4),
        )
    }
    return m, nil
}

func (m model) View() string {
    if m.err != nil {
        return fmt.Sprintf("Error: %v\n", m.err)
    }

    rendered, err := m.renderer.Render(m.content)
    if err != nil {
        return fmt.Sprintf("Render error: %v\n", err)
    }

    return strings.TrimSpace(rendered) + "\n"
}

func main() {
    p := tea.NewProgram(initialModel(), tea.WithAltScreen())
    p.Run()
}
```

----------

## Markdown Viewer
Complete markdown file viewer application
```go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "github.com/charmbracelet/glamour"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatal("Usage: viewer <markdown-file>")
    }

    filename := os.Args[1]
    content, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatalf("Could not read file %s: %v", filename, err)
    }

    renderer, err := glamour.NewTermRenderer(
        glamour.WithAutoStyle(),
        glamour.WithWordWrap(100),
        glamour.WithEmoji(),
    )
    if err != nil {
        log.Fatal(err)
    }

    out, err := renderer.RenderBytes(content)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("=== %s ===\n\n", filename)
    fmt.Print(out)
}
```