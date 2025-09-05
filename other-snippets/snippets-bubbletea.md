## Topics Covered [Bubbletea]
Examples and snippets for bubbletea TUI applications
- Basic model structure
- Interactive counter
- Timer with countdown
- Text input form
- Selection list
- Loading spinner
- Multi-view navigation
- Progress bar
- Data table
- Styled interface with lipgloss
- Multi-field form
- Responsive window size handling
- External command execution

----------

## Basic Model Structure
Minimal bubbletea application with model, init, update, view
```go
package main

import (
    "fmt"
    tea "github.com/charmbracelet/bubbletea"
    "os"
)

type model struct{}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        if msg.String() == "ctrl+c" {
            return m, tea.Quit
        }
    }
    return m, nil
}

func (m model) View() string {
    return "Hello World! Press Ctrl+C to exit.\n"
}

func main() {
    p := tea.NewProgram(model{})
    if _, err := p.Run(); err != nil {
        fmt.Printf("Error: %v", err)
        os.Exit(1)
    }
}
```

----------

## Simple Counter
Interactive counter with keyboard controls
```go
package main

import (
    "fmt"
    tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    counter int
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "up":
            m.counter++
        case "down":
            m.counter--
        case "q", "ctrl+c":
            return m, tea.Quit
        }
    }
    return m, nil
}

func (m model) View() string {
    return fmt.Sprintf(
        "Counter: %d\n\nPress ↑/↓ to change, q to quit\n",
        m.counter,
    )
}
```

----------

## Countdown Timer
Timer application with automatic updates
```go
package main

import (
    "fmt"
    "time"
    tea "github.com/charmbracelet/bubbletea"
)

type model int

type tickMsg time.Time

func tick() tea.Cmd {
    return func() tea.Msg {
        time.Sleep(time.Second)
        return tickMsg{}
    }
}

func (m model) Init() tea.Cmd {
    return tick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        if msg.String() == "q" || msg.String() == "ctrl+c" {
            return m, tea.Quit
        }
    case tickMsg:
        m--
        if m <= 0 {
            return m, tea.Quit
        }
        return m, tick()
    }
    return m, nil
}

func (m model) View() string {
    return fmt.Sprintf("Countdown: %d\nPress q to quit\n", m)
}

func main() {
    p := tea.NewProgram(model(5))
    p.Run()
}
```

----------

## Text Input Form
Single text input with validation
```go
package main

import (
    "fmt"
    "github.com/charmbracelet/bubbles/textinput"
    tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    textInput textinput.Model
    err       error
}

func initialModel() model {
    ti := textinput.New()
    ti.Placeholder = "Enter your name..."
    ti.Focus()
    ti.CharLimit = 156
    ti.Width = 20

    return model{
        textInput: ti,
        err:       nil,
    }
}

func (m model) Init() tea.Cmd {
    return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd

    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.Type {
        case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
            return m, tea.Quit
        }

    case error:
        m.err = msg
        return m, nil
    }

    m.textInput, cmd = m.textInput.Update(msg)
    return m, cmd
}

func (m model) View() string {
    return fmt.Sprintf(
        "What's your name?\n\n%s\n\n%s",
        m.textInput.View(),
        "(esc to quit)",
    ) + "\n"
}
```

----------

## Simple List Selection
Menu with selectable items
```go
package main

import (
    "fmt"
    tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    choices  []string
    cursor   int
    selected map[int]struct{}
}

func initialModel() model {
    return model{
        choices:  []string{"Option 1", "Option 2", "Option 3", "Option 4"},
        selected: make(map[int]struct{}),
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
        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }
        case "down", "j":
            if m.cursor < len(m.choices)-1 {
                m.cursor++
            }
        case "enter", " ":
            if _, ok := m.selected[m.cursor]; ok {
                delete(m.selected, m.cursor)
            } else {
                m.selected[m.cursor] = struct{}{}
            }
        }
    }
    return m, nil
}

func (m model) View() string {
    s := "Select items:\n\n"

    for i, choice := range m.choices {
        cursor := " "
        if m.cursor == i {
            cursor = ">"
        }

        checked := " "
        if _, ok := m.selected[i]; ok {
            checked = "x"
        }

        s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
    }

    s += "\nPress q to quit.\n"
    return s
}
```

----------

## Loading Spinner
Animated loading spinner
```go
package main

import (
    "github.com/charmbracelet/bubbles/spinner"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

type model struct {
    spinner spinner.Model
}

func initialModel() model {
    s := spinner.New()
    s.Spinner = spinner.Dot
    s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
    return model{spinner: s}
}

func (m model) Init() tea.Cmd {
    return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "q", "esc", "ctrl+c":
            return m, tea.Quit
        default:
            return m, nil
        }

    default:
        var cmd tea.Cmd
        m.spinner, cmd = m.spinner.Update(msg)
        return m, cmd
    }
}

func (m model) View() string {
    return fmt.Sprintf("\n\n   %s Loading...press q to quit\n\n", m.spinner.View())
}
```

----------

## Multi-View Navigation
Application with multiple screens/views
```go
package main

import (
    "fmt"
    tea "github.com/charmbracelet/bubbletea"
)

type sessionState int

const (
    menuView sessionState = iota
    detailView
)

type model struct {
    state   sessionState
    choices []string
    cursor  int
}

func initialModel() model {
    return model{
        state:   menuView,
        choices: []string{"Home", "Profile", "Settings", "About"},
    }
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch m.state {
        case menuView:
            switch msg.String() {
            case "ctrl+c", "q":
                return m, tea.Quit
            case "up":
                if m.cursor > 0 {
                    m.cursor--
                }
            case "down":
                if m.cursor < len(m.choices)-1 {
                    m.cursor++
                }
            case "enter":
                m.state = detailView
            }
        case detailView:
            switch msg.String() {
            case "ctrl+c", "q":
                return m, tea.Quit
            case "esc":
                m.state = menuView
            }
        }
    }
    return m, nil
}

func (m model) View() string {
    switch m.state {
    case menuView:
        s := "Choose an option:\n\n"
        for i, choice := range m.choices {
            cursor := " "
            if m.cursor == i {
                cursor = ">"
            }
            s += fmt.Sprintf("%s %s\n", cursor, choice)
        }
        s += "\nPress enter to select, q to quit.\n"
        return s

    case detailView:
        return fmt.Sprintf("You selected: %s\n\nPress esc to go back, q to quit.\n", m.choices[m.cursor])

    default:
        return ""
    }
}
```

----------

## Progress Bar
Animated progress indicator
```go
package main

import (
    "fmt"
    "time"
    "github.com/charmbracelet/bubbles/progress"
    tea "github.com/charmbracelet/bubbletea"
)

const (
    padding  = 2
    maxWidth = 80
)

type tickMsg time.Time

type model struct {
    progress progress.Model
    percent  float64
}

func initialModel() model {
    return model{
        progress: progress.New(progress.WithDefaultGradient()),
        percent:  0.0,
    }
}

func tick() tea.Cmd {
    return func() tea.Msg {
        time.Sleep(time.Millisecond * 100)
        return tickMsg{}
    }
}

func (m model) Init() tea.Cmd {
    return tick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        return m, tea.Quit

    case tickMsg:
        if m.percent >= 1.0 {
            return m, tea.Quit
        }

        m.percent += 0.01
        return m, tick()

    case tea.WindowSizeMsg:
        m.progress.Width = msg.Width - padding*2 - 4
        if m.progress.Width > maxWidth {
            m.progress.Width = maxWidth
        }
        return m, nil

    default:
        return m, nil
    }
}

func (m model) View() string {
    pad := strings.Repeat(" ", padding)
    return "\n" +
        pad + m.progress.ViewAs(m.percent) + "\n\n" +
        pad + fmt.Sprintf("%.0f%%", m.percent*100) + "\n\n" +
        pad + "Press any key to quit" + "\n"
}
```

----------

## Table Data Display
Simple data table rendering
```go
package main

import (
    "fmt"
    "github.com/charmbracelet/bubbles/table"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
    BorderStyle(lipgloss.NormalBorder()).
    BorderForeground(lipgloss.Color("240"))

type model struct {
    table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "esc", "q", "ctrl+c":
            return m, tea.Quit
        case "enter":
            return m, tea.Batch(
                tea.Printf("Selected: %s", m.table.SelectedRow()[1]),
            )
        }
    }
    m.table, cmd = m.table.Update(msg)
    return m, cmd
}

func (m model) View() string {
    return baseStyle.Render(m.table.View()) + "\n"
}

func main() {
    columns := []table.Column{
        {Title: "Name", Width: 10},
        {Title: "Age", Width: 4},
        {Title: "City", Width: 12},
    }

    rows := []table.Row{
        {"Alice", "25", "New York"},
        {"Bob", "30", "London"},
        {"Charlie", "35", "Tokyo"},
    }

    t := table.New(
        table.WithColumns(columns),
        table.WithRows(rows),
        table.WithFocused(true),
        table.WithHeight(7),
    )

    m := model{t}
    if _, err := tea.NewProgram(m).Run(); err != nil {
        fmt.Println("Error running program:", err)
        os.Exit(1)
    }
}
```

----------

## Styled Interface
Using lipgloss for styling
```go
package main

import (
    "fmt"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

var (
    titleStyle = lipgloss.NewStyle().
        Bold(true).
        Foreground(lipgloss.Color("#FAFAFA")).
        Background(lipgloss.Color("#7D56F4")).
        PaddingTop(2).
        PaddingLeft(4).
        Width(22)

    statusMessageStyle = lipgloss.NewStyle().
        Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
        Render
)

type model struct {
    status string
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "q", "ctrl+c":
            return m, tea.Quit
        case "enter":
            m.status = "You pressed enter!"
        }
    }
    return m, nil
}

func (m model) View() string {
    title := titleStyle.Render("Styled App")
    status := statusMessageStyle(m.status)
    
    return fmt.Sprintf(
        "%s\n\n%s\n\nPress enter for status, q to quit\n",
        title,
        status,
    )
}
```

----------

## Multiple Text Inputs
Form with multiple input fields
```go
package main

import (
    "fmt"
    "github.com/charmbracelet/bubbles/textinput"
    tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    inputs  []textinput.Model
    focused int
    err     error
}

func initialModel() model {
    inputs := make([]textinput.Model, 3)
    
    inputs[0] = textinput.New()
    inputs[0].Placeholder = "Name"
    inputs[0].Focus()
    inputs[0].CharLimit = 32
    inputs[0].Width = 20
    
    inputs[1] = textinput.New()
    inputs[1].Placeholder = "Email"
    inputs[1].CharLimit = 64
    inputs[1].Width = 20
    
    inputs[2] = textinput.New()
    inputs[2].Placeholder = "Age"
    inputs[2].CharLimit = 3
    inputs[2].Width = 20

    return model{
        inputs:  inputs,
        focused: 0,
    }
}

func (m model) Init() tea.Cmd {
    return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "esc":
            return m, tea.Quit
        case "tab", "shift+tab", "enter", "up", "down":
            s := msg.String()

            if s == "enter" && m.focused == len(m.inputs)-1 {
                return m, tea.Quit
            }

            if s == "up" || s == "shift+tab" {
                m.focused--
            } else {
                m.focused++
            }

            if m.focused > len(m.inputs)-1 {
                m.focused = 0
            } else if m.focused < 0 {
                m.focused = len(m.inputs) - 1
            }

            for i := 0; i <= len(m.inputs)-1; i++ {
                if i == m.focused {
                    m.inputs[i].Focus()
                } else {
                    m.inputs[i].Blur()
                }
            }
        }
    }

    cmd := m.updateInputs(msg)
    return m, cmd
}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
    cmds := make([]tea.Cmd, len(m.inputs))

    for i := range m.inputs {
        m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
    }

    return tea.Batch(cmds...)
}

func (m model) View() string {
    var b strings.Builder

    b.WriteString("User Registration\n\n")

    for i := range m.inputs {
        b.WriteString(m.inputs[i].View())
        if i < len(m.inputs)-1 {
            b.WriteRune('\n')
        }
    }

    button := &blurredButton
    if m.focused == len(m.inputs) {
        button = &focusedButton
    }
    fmt.Fprintf(&b, "\n\n%s\n\n", *button)

    b.WriteString("(tab/shift+tab to navigate)")
    return b.String()
}

var (
    focusedButton = lipgloss.NewStyle().
        Foreground(lipgloss.Color("205")).
        Background(lipgloss.Color("235")).
        Padding(0, 3).
        MarginTop(1).
        Render("Submit")

    blurredButton = lipgloss.NewStyle().
        Foreground(lipgloss.Color("240")).
        Background(lipgloss.Color("236")).
        Padding(0, 3).
        MarginTop(1).
        Render("Submit")
)
```

----------

## Window Size Handling
Responsive layout based on terminal size
```go
package main

import (
    "fmt"
    "strings"
    tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    width  int
    height int
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        if msg.String() == "ctrl+c" || msg.String() == "q" {
            return m, tea.Quit
        }
    case tea.WindowSizeMsg:
        m.height = msg.Height
        m.width = msg.Width
    }
    return m, nil
}

func (m model) View() string {
    if m.width == 0 {
        return "Loading..."
    }

    title := "Terminal Size Demo"
    content := fmt.Sprintf("Width: %d, Height: %d", m.width, m.height)
    
    // Center the content
    titlePadding := (m.width - len(title)) / 2
    contentPadding := (m.width - len(content)) / 2
    
    verticalPadding := (m.height - 4) / 2

    s := strings.Repeat("\n", verticalPadding)
    s += strings.Repeat(" ", titlePadding) + title + "\n\n"
    s += strings.Repeat(" ", contentPadding) + content + "\n\n"
    s += "Press q to quit"
    
    return s
}
```

----------

## Command Execution
Running external commands with feedback
```go
package main

import (
    "fmt"
    "os/exec"
    tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    output string
    err    error
}

type commandFinished struct {
    output string
    err    error
}

func runCommand() tea.Cmd {
    return func() tea.Msg {
        cmd := exec.Command("ls", "-la")
        output, err := cmd.Output()
        return commandFinished{string(output), err}
    }
}

func (m model) Init() tea.Cmd {
    return runCommand()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q":
            return m, tea.Quit
        case "r":
            return m, runCommand()
        }
    case commandFinished:
        m.output = msg.output
        m.err = msg.err
    }
    return m, nil
}

func (m model) View() string {
    s := "Command Output:\n\n"
    
    if m.err != nil {
        s += fmt.Sprintf("Error: %v\n", m.err)
    } else if m.output != "" {
        s += m.output
    } else {
        s += "Running command...\n"
    }
    
    s += "\nPress 'r' to rerun, 'q' to quit\n"
    return s
}
```