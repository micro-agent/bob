## Topics Covered [Bubbles]
Examples and snippets for bubbles UI components
- Basic text input and password masking
- Multi-line text area
- Simple list with custom delegate
- Animated loading spinner
- Progress bar indicator
- Interactive data table
- Scrollable content viewport
- Paginator for navigation
- File picker for file selection
- Countdown timer
- Stopwatch time tracking
- Auto-generated help menu

----------

## Basic Text Input
Simple text input field with placeholder
```go
package main

import (
    "fmt"
    "github.com/charmbracelet/bubbles/textinput"
    tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    textInput textinput.Model
}

func initialModel() model {
    ti := textinput.New()
    ti.Placeholder = "Enter text..."
    ti.Focus()
    ti.CharLimit = 50
    ti.Width = 20

    return model{textInput: ti}
}

func (m model) Init() tea.Cmd {
    return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd

    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.Type {
        case tea.KeyCtrlC, tea.KeyEsc:
            return m, tea.Quit
        }
    }

    m.textInput, cmd = m.textInput.Update(msg)
    return m, cmd
}

func (m model) View() string {
    return fmt.Sprintf("Input: %s\n(esc to quit)", m.textInput.View())
}
```

----------

## Password Input
Text input with masked password mode
```go
package main

import (
    "fmt"
    "github.com/charmbracelet/bubbles/textinput"
    tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    textInput textinput.Model
}

func initialModel() model {
    ti := textinput.New()
    ti.Placeholder = "Password"
    ti.EchoMode = textinput.EchoPassword
    ti.EchoCharacter = '•'
    ti.Focus()

    return model{textInput: ti}
}

func (m model) Init() tea.Cmd {
    return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd

    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.Type {
        case tea.KeyCtrlC, tea.KeyEsc:
            return m, tea.Quit
        case tea.KeyEnter:
            return m, tea.Printf("Password: %s", m.textInput.Value())
        }
    }

    m.textInput, cmd = m.textInput.Update(msg)
    return m, cmd
}

func (m model) View() string {
    return fmt.Sprintf("Password: %s\n(enter to submit)", m.textInput.View())
}
```

----------

## Text Area
Multi-line text input area
```go
package main

import (
    "fmt"
    "github.com/charmbracelet/bubbles/textarea"
    tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    textarea textarea.Model
}

func initialModel() model {
    ta := textarea.New()
    ta.Placeholder = "Enter your message..."
    ta.Focus()
    ta.CharLimit = 300
    ta.SetWidth(50)
    ta.SetHeight(5)

    return model{textarea: ta}
}

func (m model) Init() tea.Cmd {
    return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd

    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.Type {
        case tea.KeyCtrlC:
            return m, tea.Quit
        case tea.KeyEsc:
            if m.textarea.Focused() {
                m.textarea.Blur()
            } else {
                m.textarea.Focus()
            }
        }
    }

    m.textarea, cmd = m.textarea.Update(msg)
    return m, cmd
}

func (m model) View() string {
    return fmt.Sprintf(
        "Message:\n%s\n\n(ctrl+c to quit, esc to focus/blur)",
        m.textarea.View(),
    )
}
```

----------

## Simple List
Basic list with custom items
```go
package main

import (
    "fmt"
    "github.com/charmbracelet/bubbles/list"
    tea "github.com/charmbracelet/bubbletea"
    "io"
)

type item string

func (i item) FilterValue() string { return string(i) }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
    i, ok := listItem.(item)
    if !ok {
        return
    }

    str := fmt.Sprintf("%d. %s", index+1, i)

    fn := itemStyle.Render
    if index == m.Index() {
        fn = func(s ...string) string {
            return selectedItemStyle.Render("> " + strings.Join(s, " "))
        }
    }

    fmt.Fprint(w, fn(str))
}

type model struct {
    list list.Model
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch keypress := msg.String(); keypress {
        case "ctrl+c":
            return m, tea.Quit
        case "enter":
            i, ok := m.list.SelectedItem().(item)
            if ok {
                return m, tea.Printf("Selected: %s", string(i))
            }
        }
    }

    var cmd tea.Cmd
    m.list, cmd = m.list.Update(msg)
    return m, cmd
}

func (m model) View() string {
    return "\n" + m.list.View()
}

func main() {
    items := []list.Item{
        item("Apple"),
        item("Banana"),
        item("Cherry"),
    }

    l := list.New(items, itemDelegate{}, 20, 14)
    l.Title = "Fruits"
    l.SetShowStatusBar(false)
    l.SetFilteringEnabled(false)

    m := model{list: l}
    tea.NewProgram(m).Run()
}
```

----------

## Spinner Loading
Animated spinner for loading states
```go
package main

import (
    "fmt"
    "github.com/charmbracelet/bubbles/spinner"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

type model struct {
    spinner spinner.Model
    loading bool
}

func initialModel() model {
    s := spinner.New()
    s.Spinner = spinner.Dot
    s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
    return model{
        spinner: s,
        loading: true,
    }
}

func (m model) Init() tea.Cmd {
    return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "q", "ctrl+c":
            return m, tea.Quit
        case "s":
            m.loading = !m.loading
            if m.loading {
                return m, m.spinner.Tick
            }
        }
        return m, nil

    default:
        var cmd tea.Cmd
        m.spinner, cmd = m.spinner.Update(msg)
        return m, cmd
    }
}

func (m model) View() string {
    var str string
    if m.loading {
        str = fmt.Sprintf("%s Loading...", m.spinner.View())
    } else {
        str = "Done! Press 's' to restart spinner."
    }
    return fmt.Sprintf("%s\n\nPress 'q' to quit", str)
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

func tickCmd() tea.Cmd {
    return func() tea.Msg {
        time.Sleep(time.Millisecond * 100)
        return tickMsg{}
    }
}

func (m model) Init() tea.Cmd {
    return tickCmd()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        return m, tea.Quit

    case tickMsg:
        if m.percent >= 1.0 {
            return m, tea.Quit
        }
        m.percent += 0.02
        return m, tickCmd()

    case tea.WindowSizeMsg:
        m.progress.Width = msg.Width - 4
        if m.progress.Width > 80 {
            m.progress.Width = 80
        }
        return m, nil

    default:
        return m, nil
    }
}

func (m model) View() string {
    return fmt.Sprintf(
        "\n%s\n\n%.0f%% complete\n\nPress any key to quit",
        m.progress.ViewAs(m.percent),
        m.percent*100,
    )
}
```

----------

## Table Display
Interactive data table
```go
package main

import (
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
            return m, tea.Printf("You selected: %s", m.table.SelectedRow()[1])
        }
    }
    m.table, cmd = m.table.Update(msg)
    return m, cmd
}

func (m model) View() string {
    return baseStyle.Render(m.table.View()) + "\n  " + m.table.HelpView() + "\n"
}

func main() {
    columns := []table.Column{
        {Title: "Rank", Width: 4},
        {Title: "City", Width: 10},
        {Title: "Country", Width: 10},
        {Title: "Population", Width: 10},
    }

    rows := []table.Row{
        {"1", "Tokyo", "Japan", "37,274,000"},
        {"2", "Delhi", "India", "32,065,000"},
        {"3", "Shanghai", "China", "28,516,000"},
        {"4", "Dhaka", "Bangladesh", "22,478,000"},
    }

    t := table.New(
        table.WithColumns(columns),
        table.WithRows(rows),
        table.WithFocused(true),
        table.WithHeight(7),
    )

    s := table.DefaultStyles()
    s.Header = s.Header.
        BorderStyle(lipgloss.NormalBorder()).
        BorderForeground(lipgloss.Color("240")).
        BorderBottom(true).
        Bold(false)
    s.Selected = s.Selected.
        Foreground(lipgloss.Color("229")).
        Background(lipgloss.Color("57")).
        Bold(false)
    t.SetStyles(s)

    m := model{t}
    tea.NewProgram(m).Run()
}
```

----------

## Viewport Scrolling
Scrollable content viewer
```go
package main

import (
    "fmt"
    "strings"
    "github.com/charmbracelet/bubbles/viewport"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

const useHighPerformanceRenderer = false

var (
    titleStyle = func() lipgloss.Style {
        b := lipgloss.RoundedBorder()
        b.Right = "├"
        return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
    }()

    infoStyle = func() lipgloss.Style {
        b := lipgloss.RoundedBorder()
        b.Left = "┤"
        return titleStyle.Copy().BorderStyle(b)
    }()
)

type model struct {
    content  string
    ready    bool
    viewport viewport.Model
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd

    switch msg := msg.(type) {
    case tea.KeyMsg:
        if k := msg.String(); k == "ctrl+c" || k == "q" || k == "esc" {
            return m, tea.Quit
        }

    case tea.WindowSizeMsg:
        headerHeight := lipgloss.Height(m.headerView())
        footerHeight := lipgloss.Height(m.footerView())
        verticalMarginHeight := headerHeight + footerHeight

        if !m.ready {
            m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
            m.viewport.YPosition = headerHeight
            m.viewport.HighPerformanceRendering = useHighPerformanceRenderer
            m.viewport.SetContent(m.content)
            m.ready = true

            m.viewport.YPosition = headerHeight + 1
        } else {
            m.viewport.Width = msg.Width
            m.viewport.Height = msg.Height - verticalMarginHeight
        }

        if useHighPerformanceRenderer {
            cmds := []tea.Cmd{viewport.Sync(m.viewport)}
            cmd = tea.Batch(cmds...)
        }
    }

    m.viewport, cmd = m.viewport.Update(msg)

    return m, cmd
}

func (m model) View() string {
    if !m.ready {
        return "\n  Initializing..."
    }
    return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.viewport.View(), m.footerView())
}

func (m model) headerView() string {
    title := titleStyle.Render("Mr. Pager")
    line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
    return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m model) footerView() string {
    info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
    line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info)))
    return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    content := strings.Repeat("Line content here\n", 100)
    
    m := model{content: content}
    tea.NewProgram(m, tea.WithAltScreen()).Run()
}
```

----------

## Paginator Navigation
Pagination component for long lists
```go
package main

import (
    "fmt"
    "strings"
    "github.com/charmbracelet/bubbles/paginator"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

var (
    subtle = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
    
    itemStyle = lipgloss.NewStyle().
        PaddingLeft(4)
        
    selectedItemStyle = lipgloss.NewStyle().
        PaddingLeft(2).
        Foreground(lipgloss.Color("170"))
)

type model struct {
    paginator paginator.Model
    items     []string
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "q", "esc", "ctrl+c":
            return m, tea.Quit
        }
    }
    m.paginator, cmd = m.paginator.Update(msg)
    return m, cmd
}

func (m model) View() string {
    var b strings.Builder

    b.WriteString("Items:\n\n")

    start, end := m.paginator.GetSliceBounds(len(m.items))
    for i, item := range m.items[start:end] {
        if i == m.paginator.Page {
            b.WriteString(selectedItemStyle.Render(fmt.Sprintf("• %s", item)))
        } else {
            b.WriteString(itemStyle.Render(item))
        }
        b.WriteRune('\n')
    }

    b.WriteString("\n" + m.paginator.View())
    b.WriteString("\n\nPress q to quit")
    return b.String()
}

func main() {
    items := make([]string, 100)
    for i := range items {
        items[i] = fmt.Sprintf("Item #%d", i+1)
    }

    p := paginator.New()
    p.Type = paginator.Arabic
    p.PerPage = 10
    p.SetTotalPages(len(items))

    m := model{
        paginator: p,
        items:     items,
    }

    tea.NewProgram(m).Run()
}
```

----------

## File Picker
Directory navigation and file selection
```go
package main

import (
    "github.com/charmbracelet/bubbles/filepicker"
    tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    filepicker filepicker.Model
    selectedFile string
    quitting     bool
    err          error
}

type clearErrorMsg struct{}

func clearErrorAfter(t time.Duration) tea.Cmd {
    return func() tea.Msg {
        time.Sleep(t)
        return clearErrorMsg{}
    }
}

func (m model) Init() tea.Cmd {
    return m.filepicker.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q":
            m.quitting = true
            return m, tea.Quit
        }
    case clearErrorMsg:
        m.err = nil
    }

    var cmd tea.Cmd
    m.filepicker, cmd = m.filepicker.Update(msg)

    if didSelect, path := m.filepicker.DidSelectFile(msg); didSelect {
        m.selectedFile = path
    }

    if didSelect, path := m.filepicker.DidSelectDisabledFile(msg); didSelect {
        m.err = errors.New(path + " is not valid.")
        m.selectedFile = ""
        return m, tea.Batch(cmd, clearErrorAfter(2*time.Second))
    }

    return m, cmd
}

func (m model) View() string {
    if m.quitting {
        return ""
    }
    var s strings.Builder
    s.WriteString("\n  ")
    if m.err != nil {
        s.WriteString(m.filepicker.Styles.DisabledFile.Render(m.err.Error()))
    } else if m.selectedFile == "" {
        s.WriteString("Pick a file:")
    } else {
        s.WriteString("Selected file: " + m.filepicker.Styles.Selected.Render(m.selectedFile))
    }
    s.WriteString("\n\n" + m.filepicker.View() + "\n")
    return s.String()
}

func main() {
    fp := filepicker.New()
    fp.AllowedTypes = []string{".mod", ".sum", ".go", ".txt", ".md"}
    fp.CurrentDirectory, _ = os.UserHomeDir()

    m := model{
        filepicker: fp,
    }
    tea.NewProgram(m, tea.WithOutput(os.Stderr)).Run()
}
```

----------

## Timer Countdown
Configurable countdown timer
```go
package main

import (
    "fmt"
    "time"
    "github.com/charmbracelet/bubbles/timer"
    tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    timer timer.Model
}

func (m model) Init() tea.Cmd {
    return m.timer.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case timer.TickMsg:
        var cmd tea.Cmd
        m.timer, cmd = m.timer.Update(msg)
        return m, cmd

    case timer.StartStopMsg:
        var cmd tea.Cmd
        m.timer, cmd = m.timer.Update(msg)
        return m, cmd

    case timer.TimeoutMsg:
        return m, tea.Quit

    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q":
            return m, tea.Quit
        case "s":
            return m, m.timer.Toggle()
        case "r":
            m.timer.Timeout = time.Second * 10
            return m, m.timer.Init()
        }
    }
    return m, nil
}

func (m model) View() string {
    s := fmt.Sprintf("Timer: %s\n", m.timer.View())
    s += "Press 's' to start/stop, 'r' to reset, 'q' to quit"
    return s
}

func main() {
    m := model{
        timer: timer.NewWithInterval(time.Second*10, time.Second),
    }

    tea.NewProgram(m).Run()
}
```

----------

## Stopwatch
Time tracking stopwatch
```go
package main

import (
    "fmt"
    "time"
    "github.com/charmbracelet/bubbles/stopwatch"
    tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    stopwatch stopwatch.Model
}

func (m model) Init() tea.Cmd {
    return m.stopwatch.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q":
            return m, tea.Quit
        case "s":
            return m, m.stopwatch.Toggle()
        case "r":
            return m, m.stopwatch.Reset()
        }
    case stopwatch.TickMsg:
        var cmd tea.Cmd
        m.stopwatch, cmd = m.stopwatch.Update(msg)
        return m, cmd
    case stopwatch.StartStopMsg:
        var cmd tea.Cmd
        m.stopwatch, cmd = m.stopwatch.Update(msg)
        return m, cmd
    }
    return m, nil
}

func (m model) View() string {
    s := fmt.Sprintf("Stopwatch: %s\n", m.stopwatch.View())
    s += "Press 's' to start/stop, 'r' to reset, 'q' to quit"
    return s
}

func main() {
    m := model{
        stopwatch: stopwatch.NewWithInterval(time.Millisecond),
    }

    tea.NewProgram(m).Run()
}
```

----------

## Help Menu
Auto-generated help from key bindings
```go
package main

import (
    "github.com/charmbracelet/bubbles/help"
    "github.com/charmbracelet/bubbles/key"
    tea "github.com/charmbracelet/bubbletea"
)

type keyMap struct {
    Up    key.Binding
    Down  key.Binding
    Left  key.Binding
    Right key.Binding
    Help  key.Binding
    Quit  key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
    return []key.Binding{k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
    return [][]key.Binding{
        {k.Up, k.Down, k.Left, k.Right},
        {k.Help, k.Quit},
    }
}

var keys = keyMap{
    Up: key.NewBinding(
        key.WithKeys("up", "k"),
        key.WithHelp("↑/k", "move up"),
    ),
    Down: key.NewBinding(
        key.WithKeys("down", "j"),
        key.WithHelp("↓/j", "move down"),
    ),
    Left: key.NewBinding(
        key.WithKeys("left", "h"),
        key.WithHelp("←/h", "move left"),
    ),
    Right: key.NewBinding(
        key.WithKeys("right", "l"),
        key.WithHelp("→/l", "move right"),
    ),
    Help: key.NewBinding(
        key.WithKeys("?"),
        key.WithHelp("?", "toggle help"),
    ),
    Quit: key.NewBinding(
        key.WithKeys("q", "esc", "ctrl+c"),
        key.WithHelp("q", "quit"),
    ),
}

type model struct {
    keys keyMap
    help help.Model
    showFullHelp bool
}

func initialModel() model {
    return model{
        keys: keys,
        help: help.New(),
    }
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch {
        case key.Matches(msg, m.keys.Help):
            m.showFullHelp = !m.showFullHelp
        case key.Matches(msg, m.keys.Quit):
            return m, tea.Quit
        }
    case tea.WindowSizeMsg:
        m.help.Width = msg.Width
    }

    return m, nil
}

func (m model) View() string {
    var helpView string
    if m.showFullHelp {
        helpView = m.help.View(m.keys)
    } else {
        helpView = m.help.ShortHelpView([]key.Binding{
            m.keys.Help, m.keys.Quit,
        })
    }
    return fmt.Sprintf("Example app\n\n%s", helpView)
}
```