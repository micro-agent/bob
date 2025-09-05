## Topics Covered [Lipgloss]
Examples and snippets for lipgloss terminal styling
- Basic text styling and colors
- Padding and margins
- Border styles and customization
- Text alignment and positioning
- Adaptive colors for themes
- Layout composition
- Measuring and placing text
- Complex styling combinations
- Table and list rendering
- Color profiles and degradation
- Style inheritance
- Conditional styling

----------

## Available Properties
CSS-like styling properties for terminal text
```go
// Text formatting
.Bold(true)                    // Bold text
.Italic(true)                  // Italic text
.Underline(true)               // Underlined text
.Strikethrough(true)           // Strikethrough text
.Blink(true)                   // Blinking text

// Colors
.Foreground(lipgloss.Color("#FF0000"))  // Text color
.Background(lipgloss.Color("#000000"))  // Background color
.ColorWhitespace(true)                  // Color whitespace too

// Layout
.Width(20)                     // Fixed width
.Height(10)                    // Fixed height
.Padding(1, 2, 1, 2)          // Padding (top, right, bottom, left)
.Margin(1, 2)                  // Margin (vertical, horizontal)
.Align(lipgloss.Center)        // Text alignment

// Borders
.Border(lipgloss.RoundedBorder())      // Border style
.BorderForeground(lipgloss.Color("#FF0000"))  // Border color
```

----------

## Basic Text Styling
Simple text formatting and colors
```go
package main

import (
    "fmt"
    "github.com/charmbracelet/lipgloss"
)

func main() {
    // Basic styling
    bold := lipgloss.NewStyle().Bold(true)
    italic := lipgloss.NewStyle().Italic(true)
    underline := lipgloss.NewStyle().Underline(true)
    
    fmt.Println(bold.Render("Bold text"))
    fmt.Println(italic.Render("Italic text"))
    fmt.Println(underline.Render("Underlined text"))
    
    // Colored text
    red := lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))
    green := lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00"))
    blue := lipgloss.NewStyle().Foreground(lipgloss.Color("#0000FF"))
    
    fmt.Println(red.Render("Red text"))
    fmt.Println(green.Render("Green text"))
    fmt.Println(blue.Render("Blue text"))
}
```

----------

## Background Colors
Text with background styling
```go
package main

import (
    "fmt"
    "github.com/charmbracelet/lipgloss"
)

func main() {
    // Background colors
    highlight := lipgloss.NewStyle().
        Foreground(lipgloss.Color("#FFFFFF")).
        Background(lipgloss.Color("#7D56F4")).
        Padding(0, 1)
    
    warning := lipgloss.NewStyle().
        Foreground(lipgloss.Color("#000000")).
        Background(lipgloss.Color("#FFFF00")).
        Padding(0, 1)
    
    error := lipgloss.NewStyle().
        Foreground(lipgloss.Color("#FFFFFF")).
        Background(lipgloss.Color("#FF0000")).
        Padding(0, 1)
    
    fmt.Println(highlight.Render("Highlighted text"))
    fmt.Println(warning.Render("Warning message"))
    fmt.Println(error.Render("Error message"))
}
```

----------

## Padding and Margins
Control spacing around text
```go
package main

import (
    "fmt"
    "github.com/charmbracelet/lipgloss"
)

func main() {
    // Different padding styles
    padded := lipgloss.NewStyle().
        Background(lipgloss.Color("#7D56F4")).
        Foreground(lipgloss.Color("#FFFFFF")).
        Padding(2, 4)  // vertical, horizontal
    
    asymmetric := lipgloss.NewStyle().
        Background(lipgloss.Color("#FF6B6B")).
        Foreground(lipgloss.Color("#FFFFFF")).
        Padding(1, 2, 3, 4)  // top, right, bottom, left
    
    margined := lipgloss.NewStyle().
        Background(lipgloss.Color("#4ECDC4")).
        Foreground(lipgloss.Color("#000000")).
        Margin(2, 0)  // vertical, horizontal
        Padding(1, 2)
    
    fmt.Println(padded.Render("Padded text"))
    fmt.Println(asymmetric.Render("Asymmetric padding"))
    fmt.Println(margined.Render("Margined text"))
}
```

----------

## Border Styles
Various border types and customization
```go
package main

import (
    "fmt"
    "github.com/charmbracelet/lipgloss"
)

func main() {
    // Different border styles
    rounded := lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(lipgloss.Color("#7D56F4")).
        Padding(1, 2)
    
    thick := lipgloss.NewStyle().
        Border(lipgloss.ThickBorder()).
        BorderForeground(lipgloss.Color("#FF6B6B")).
        Padding(1, 2)
    
    double := lipgloss.NewStyle().
        Border(lipgloss.DoubleBorder()).
        BorderForeground(lipgloss.Color("#4ECDC4")).
        Padding(1, 2)
    
    // Custom border sides
    partial := lipgloss.NewStyle().
        BorderTop(true).
        BorderBottom(true).
        BorderStyle(lipgloss.NormalBorder()).
        BorderForeground(lipgloss.Color("#FFD93D")).
        Padding(0, 2)
    
    fmt.Println(rounded.Render("Rounded border"))
    fmt.Println(thick.Render("Thick border"))
    fmt.Println(double.Render("Double border"))
    fmt.Println(partial.Render("Top/bottom only"))
}
```

----------

## Text Alignment
Align text within fixed widths
```go
package main

import (
    "fmt"
    "github.com/charmbracelet/lipgloss"
)

func main() {
    width := 30
    
    left := lipgloss.NewStyle().
        Width(width).
        Align(lipgloss.Left).
        Background(lipgloss.Color("#E8E8E8")).
        Foreground(lipgloss.Color("#000000"))
    
    center := lipgloss.NewStyle().
        Width(width).
        Align(lipgloss.Center).
        Background(lipgloss.Color("#7D56F4")).
        Foreground(lipgloss.Color("#FFFFFF"))
    
    right := lipgloss.NewStyle().
        Width(width).
        Align(lipgloss.Right).
        Background(lipgloss.Color("#FF6B6B")).
        Foreground(lipgloss.Color("#FFFFFF"))
    
    fmt.Println(left.Render("Left aligned"))
    fmt.Println(center.Render("Centered"))
    fmt.Println(right.Render("Right aligned"))
}
```

----------

## Adaptive Colors
Colors that adapt to terminal background
```go
package main

import (
    "fmt"
    "github.com/charmbracelet/lipgloss"
)

func main() {
    // Adaptive colors for light/dark themes
    adaptiveStyle := lipgloss.NewStyle().
        Foreground(lipgloss.AdaptiveColor{
            Light: "#000000",  // Black on light backgrounds
            Dark:  "#FFFFFF",  // White on dark backgrounds
        }).
        Background(lipgloss.AdaptiveColor{
            Light: "#F0F0F0",  // Light gray on light themes
            Dark:  "#2D2D2D",  // Dark gray on dark themes
        }).
        Padding(1, 2)
    
    highlight := lipgloss.NewStyle().
        Foreground(lipgloss.AdaptiveColor{
            Light: "#874BFD",
            Dark:  "#7D56F4",
        }).
        Bold(true)
    
    fmt.Println(adaptiveStyle.Render("This adapts to your terminal theme"))
    fmt.Println(highlight.Render("Highlighted adaptive text"))
}
```

----------

## Layout Composition
Combining elements horizontally and vertically
```go
package main

import (
    "fmt"
    "github.com/charmbracelet/lipgloss"
)

func main() {
    // Create styled boxes
    box1 := lipgloss.NewStyle().
        Background(lipgloss.Color("#FF6B6B")).
        Foreground(lipgloss.Color("#FFFFFF")).
        Padding(1, 2).
        Render("Box 1")
    
    box2 := lipgloss.NewStyle().
        Background(lipgloss.Color("#4ECDC4")).
        Foreground(lipgloss.Color("#000000")).
        Padding(1, 2).
        Render("Box 2")
    
    box3 := lipgloss.NewStyle().
        Background(lipgloss.Color("#FFD93D")).
        Foreground(lipgloss.Color("#000000")).
        Padding(1, 2).
        Render("Box 3")
    
    // Horizontal layout
    horizontal := lipgloss.JoinHorizontal(lipgloss.Top, box1, box2, box3)
    fmt.Println("Horizontal layout:")
    fmt.Println(horizontal)
    fmt.Println()
    
    // Vertical layout
    vertical := lipgloss.JoinVertical(lipgloss.Left, box1, box2, box3)
    fmt.Println("Vertical layout:")
    fmt.Println(vertical)
}
```

----------

## Text Measurement
Measure and position text precisely
```go
package main

import (
    "fmt"
    "github.com/charmbracelet/lipgloss"
)

func main() {
    text := "Hello, World!"
    
    // Measure text dimensions
    width := lipgloss.Width(text)
    height := lipgloss.Height(text)
    
    fmt.Printf("Text dimensions: %dx%d\n", width, height)
    
    // Create a container and place text
    container := lipgloss.NewStyle().
        Width(40).
        Height(10).
        Border(lipgloss.RoundedBorder()).
        BorderForeground(lipgloss.Color("#7D56F4"))
    
    // Place text in different positions
    topLeft := lipgloss.Place(40, 10, lipgloss.Left, lipgloss.Top, text)
    center := lipgloss.Place(40, 10, lipgloss.Center, lipgloss.Center, text)
    bottomRight := lipgloss.Place(40, 10, lipgloss.Right, lipgloss.Bottom, text)
    
    fmt.Println("Top Left:")
    fmt.Println(container.Render(topLeft))
    
    fmt.Println("\nCenter:")
    fmt.Println(container.Render(center))
    
    fmt.Println("\nBottom Right:")
    fmt.Println(container.Render(bottomRight))
}
```

----------

## Complex Styling
Combining multiple style properties
```go
package main

import (
    "fmt"
    "github.com/charmbracelet/lipgloss"
)

func main() {
    // Complex card-like component
    card := lipgloss.NewStyle().
        Bold(true).
        Foreground(lipgloss.Color("#FFFFFF")).
        Background(lipgloss.Color("#7D56F4")).
        Border(lipgloss.RoundedBorder()).
        BorderForeground(lipgloss.Color("#9B7FFF")).
        Padding(2, 4).
        Margin(1, 0).
        Width(50).
        Align(lipgloss.Center)
    
    // Nested styling
    title := lipgloss.NewStyle().
        Bold(true).
        Underline(true).
        Foreground(lipgloss.Color("#FFD93D"))
    
    subtitle := lipgloss.NewStyle().
        Italic(true).
        Foreground(lipgloss.Color("#E8E8E8"))
    
    content := fmt.Sprintf("%s\n\n%s\n\nThis is the main content of the card with multiple styling elements combined together.",
        title.Render("Card Title"),
        subtitle.Render("Subtitle goes here"))
    
    fmt.Println(card.Render(content))
}
```

----------

## Color Profiles
Different color capabilities
```go
package main

import (
    "fmt"
    "github.com/charmbracelet/lipgloss"
)

func main() {
    // True color (24-bit)
    trueColor := lipgloss.NewStyle().
        Foreground(lipgloss.Color("#FF69B4")).
        Background(lipgloss.Color("#1E1E2E")).
        Padding(0, 1)
    
    // ANSI 256 colors
    ansi256 := lipgloss.NewStyle().
        Foreground(lipgloss.Color("205")).  // Pink
        Background(lipgloss.Color("235")).  // Dark gray
        Padding(0, 1)
    
    // Basic ANSI colors
    basicANSI := lipgloss.NewStyle().
        Foreground(lipgloss.Color("magenta")).
        Background(lipgloss.Color("black")).
        Padding(0, 1)
    
    fmt.Println("Color profile examples:")
    fmt.Println(trueColor.Render("True Color (24-bit)"))
    fmt.Println(ansi256.Render("ANSI 256 Color"))
    fmt.Println(basicANSI.Render("Basic ANSI Color"))
}
```

----------

## Style Inheritance
Creating style variations from base styles
```go
package main

import (
    "fmt"
    "github.com/charmbracelet/lipgloss"
)

func main() {
    // Base button style
    baseButton := lipgloss.NewStyle().
        Padding(0, 2).
        Border(lipgloss.RoundedBorder()).
        Bold(true).
        Align(lipgloss.Center).
        Width(15)
    
    // Inherit and customize for different button types
    primaryButton := baseButton.Copy().
        Foreground(lipgloss.Color("#FFFFFF")).
        Background(lipgloss.Color("#7D56F4")).
        BorderForeground(lipgloss.Color("#9B7FFF"))
    
    secondaryButton := baseButton.Copy().
        Foreground(lipgloss.Color("#7D56F4")).
        Background(lipgloss.Color("#FFFFFF")).
        BorderForeground(lipgloss.Color("#7D56F4"))
    
    dangerButton := baseButton.Copy().
        Foreground(lipgloss.Color("#FFFFFF")).
        Background(lipgloss.Color("#FF4757")).
        BorderForeground(lipgloss.Color("#FF3838"))
    
    fmt.Println("Button styles:")
    fmt.Println(primaryButton.Render("Primary"))
    fmt.Println(secondaryButton.Render("Secondary"))
    fmt.Println(dangerButton.Render("Danger"))
}
```

----------

## Table Rendering
Create styled tables with lipgloss
```go
package main

import (
    "fmt"
    "strings"
    "github.com/charmbracelet/lipgloss"
    "github.com/charmbracelet/lipgloss/table"
)

func main() {
    // Table data
    headers := []string{"Name", "Age", "City"}
    rows := [][]string{
        {"Alice", "25", "New York"},
        {"Bob", "30", "London"},
        {"Charlie", "35", "Tokyo"},
    }
    
    // Create table with styling
    t := table.New().
        Border(lipgloss.NormalBorder()).
        BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("240"))).
        StyleFunc(func(row, col int) lipgloss.Style {
            switch {
            case row == 0:
                return lipgloss.NewStyle().
                    Foreground(lipgloss.Color("#FFFFFF")).
                    Background(lipgloss.Color("#7D56F4")).
                    Bold(true).
                    Align(lipgloss.Center)
            case row%2 == 0:
                return lipgloss.NewStyle().
                    Foreground(lipgloss.Color("#000000")).
                    Background(lipgloss.Color("#F0F0F0"))
            default:
                return lipgloss.NewStyle().
                    Foreground(lipgloss.Color("#000000")).
                    Background(lipgloss.Color("#FFFFFF"))
            }
        }).
        Headers(headers...).
        Rows(rows...)
    
    fmt.Println(t)
}
```

----------

## Conditional Styling
Apply styles based on conditions
```go
package main

import (
    "fmt"
    "github.com/charmbracelet/lipgloss"
)

func main() {
    // Status indicators with conditional styling
    statuses := []struct {
        name   string
        active bool
        level  string
    }{
        {"Server 1", true, "success"},
        {"Server 2", false, "error"},
        {"Server 3", true, "warning"},
        {"Server 4", true, "success"},
    }
    
    for _, status := range statuses {
        // Base style
        style := lipgloss.NewStyle().
            Padding(0, 2).
            Margin(0, 1, 0, 0)
        
        // Conditional styling based on status
        switch {
        case !status.active:
            style = style.
                Foreground(lipgloss.Color("#666666")).
                Background(lipgloss.Color("#2D2D2D")).
                Strikethrough(true)
        case status.level == "success":
            style = style.
                Foreground(lipgloss.Color("#FFFFFF")).
                Background(lipgloss.Color("#28A745"))
        case status.level == "warning":
            style = style.
                Foreground(lipgloss.Color("#000000")).
                Background(lipgloss.Color("#FFC107"))
        case status.level == "error":
            style = style.
                Foreground(lipgloss.Color("#FFFFFF")).
                Background(lipgloss.Color("#DC3545"))
        }
        
        fmt.Print(style.Render(status.name))
    }
    fmt.Println()
}
```

----------

## Interactive Components
Styled components with state
```go
package main

import (
    "fmt"
    "github.com/charmbracelet/lipgloss"
)

func main() {
    // Menu items with active/inactive states
    menuItems := []struct {
        text   string
        active bool
    }{
        {"Home", true},
        {"About", false},
        {"Services", false},
        {"Contact", false},
    }
    
    // Style definitions
    activeStyle := lipgloss.NewStyle().
        Foreground(lipgloss.Color("#FFFFFF")).
        Background(lipgloss.Color("#7D56F4")).
        Padding(0, 2).
        Bold(true)
    
    inactiveStyle := lipgloss.NewStyle().
        Foreground(lipgloss.Color("#666666")).
        Background(lipgloss.Color("")).
        Padding(0, 2)
    
    // Render menu
    var menu []string
    for _, item := range menuItems {
        if item.active {
            menu = append(menu, activeStyle.Render(item.text))
        } else {
            menu = append(menu, inactiveStyle.Render(item.text))
        }
    }
    
    menuBar := lipgloss.JoinHorizontal(lipgloss.Top, menu...)
    
    // Container with border
    container := lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(lipgloss.Color("#7D56F4")).
        Padding(1, 0)
    
    fmt.Println(container.Render(menuBar))
}
```