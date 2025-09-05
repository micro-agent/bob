## Topics Covered [Huh]
Examples and snippets for huh terminal forms
- Single text input with validation
- Multi-select options
- Confirmation prompts
- Multi-step forms
- Themed forms
- Bubbletea integration
- Dynamic form fields
- Form validation
- Conditional form logic
- Quick single prompts
- Background spinner tasks
- Accessibility features

----------

## Available Components
Form field types and utilities in huh
```go
// Form fields
huh.NewInput()       // Single-line text input
huh.NewText()        // Multi-line text area
huh.NewSelect()      // Single selection from list
huh.NewMultiSelect() // Multiple selections
huh.NewConfirm()     // Yes/No confirmation

// Form structure
huh.NewForm()        // Create form container
huh.NewGroup()       // Group related fields

// Utilities
huh.NewSpinner()     // Loading spinner
huh.NewTheme()       // Custom themes
```

----------

## Simple Text Input
Basic text input with validation
```go
package main

import (
    "fmt"
    "log"
    "github.com/charmbracelet/huh"
)

func main() {
    var name string

    form := huh.NewForm(
        huh.NewGroup(
            huh.NewInput().
                Title("What's your name?").
                Value(&name).
                Validate(func(str string) error {
                    if len(str) < 2 {
                        return fmt.Errorf("name must be at least 2 characters")
                    }
                    return nil
                }),
        ),
    )

    err := form.Run()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Hello, %s!\n", name)
}
```

----------

## Multi-line Text Area
Text area for longer content
```go
package main

import (
    "fmt"
    "log"
    "github.com/charmbracelet/huh"
)

func main() {
    var description string

    form := huh.NewForm(
        huh.NewGroup(
            huh.NewText().
                Title("Project Description").
                Description("Tell us about your project").
                CharLimit(500).
                Value(&description),
        ),
    )

    err := form.Run()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Description: %s\n", description)
}
```

----------

## Single Selection
Choose one option from a list
```go
package main

import (
    "fmt"
    "log"
    "github.com/charmbracelet/huh"
)

func main() {
    var flavor string

    form := huh.NewForm(
        huh.NewGroup(
            huh.NewSelect[string]().
                Title("Choose your ice cream flavor").
                Options(
                    huh.NewOption("Vanilla", "vanilla"),
                    huh.NewOption("Chocolate", "chocolate"),
                    huh.NewOption("Strawberry", "strawberry"),
                    huh.NewOption("Mint Chip", "mint"),
                ).
                Value(&flavor),
        ),
    )

    err := form.Run()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("You chose: %s\n", flavor)
}
```

----------

## Multiple Selection
Choose multiple options from a list
```go
package main

import (
    "fmt"
    "log"
    "strings"
    "github.com/charmbracelet/huh"
)

func main() {
    var toppings []string

    form := huh.NewForm(
        huh.NewGroup(
            huh.NewMultiSelect[string]().
                Title("Select pizza toppings").
                Options(
                    huh.NewOption("Pepperoni", "pepperoni"),
                    huh.NewOption("Mushrooms", "mushrooms"),
                    huh.NewOption("Bell Peppers", "peppers"),
                    huh.NewOption("Olives", "olives"),
                    huh.NewOption("Extra Cheese", "cheese"),
                ).
                Value(&toppings),
        ),
    )

    err := form.Run()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Toppings: %s\n", strings.Join(toppings, ", "))
}
```

----------

## Confirmation Prompt
Yes/No confirmation dialog
```go
package main

import (
    "fmt"
    "log"
    "github.com/charmbracelet/huh"
)

func main() {
    var confirmed bool

    form := huh.NewForm(
        huh.NewGroup(
            huh.NewConfirm().
                Title("Delete all files?").
                Description("This action cannot be undone").
                Affirmative("Yes!").
                Negative("No.").
                Value(&confirmed),
        ),
    )

    err := form.Run()
    if err != nil {
        log.Fatal(err)
    }

    if confirmed {
        fmt.Println("Files deleted!")
    } else {
        fmt.Println("Operation cancelled")
    }
}
```

----------

## Multi-Step Form
Form with multiple groups/pages
```go
package main

import (
    "fmt"
    "log"
    "github.com/charmbracelet/huh"
)

func main() {
    var (
        firstName string
        lastName  string
        email     string
        age       int
        subscribe bool
    )

    form := huh.NewForm(
        huh.NewGroup(
            huh.NewInput().
                Title("First Name").
                Value(&firstName),
            huh.NewInput().
                Title("Last Name").
                Value(&lastName),
        ).Title("Personal Information"),

        huh.NewGroup(
            huh.NewInput().
                Title("Email").
                Value(&email),
            huh.NewInput().
                Title("Age").
                Value(&age),
        ).Title("Contact Details"),

        huh.NewGroup(
            huh.NewConfirm().
                Title("Subscribe to newsletter?").
                Value(&subscribe),
        ).Title("Preferences"),
    )

    err := form.Run()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Name: %s %s\n", firstName, lastName)
    fmt.Printf("Email: %s, Age: %d\n", email, age)
    fmt.Printf("Newsletter: %t\n", subscribe)
}
```

----------

## Themed Form
Apply custom theme to form
```go
package main

import (
    "fmt"
    "log"
    "github.com/charmbracelet/huh"
)

func main() {
    var name string

    theme := huh.ThemeCharm()

    form := huh.NewForm(
        huh.NewGroup(
            huh.NewInput().
                Title("Enter your name").
                Value(&name),
        ),
    ).WithTheme(theme)

    err := form.Run()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Hello, %s!\n", name)
}
```

----------

## Form with Validation
Multiple validation rules
```go
package main

import (
    "fmt"
    "log"
    "regexp"
    "github.com/charmbracelet/huh"
)

func main() {
    var (
        email    string
        password string
        age      int
    )

    emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

    form := huh.NewForm(
        huh.NewGroup(
            huh.NewInput().
                Title("Email").
                Value(&email).
                Validate(func(str string) error {
                    if !emailRegex.MatchString(str) {
                        return fmt.Errorf("invalid email format")
                    }
                    return nil
                }),

            huh.NewInput().
                Title("Password").
                EchoMode(huh.EchoModePassword).
                Value(&password).
                Validate(func(str string) error {
                    if len(str) < 8 {
                        return fmt.Errorf("password must be at least 8 characters")
                    }
                    return nil
                }),

            huh.NewInput().
                Title("Age").
                Value(&age).
                Validate(func(i int) error {
                    if i < 18 {
                        return fmt.Errorf("must be 18 or older")
                    }
                    return nil
                }),
        ),
    )

    err := form.Run()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Account created for %s\n", email)
}
```

----------

## Conditional Form Logic
Dynamic form based on user input
```go
package main

import (
    "fmt"
    "log"
    "github.com/charmbracelet/huh"
)

func main() {
    var (
        userType     string
        companyName  string
        personalName string
    )

    form := huh.NewForm(
        huh.NewGroup(
            huh.NewSelect[string]().
                Title("Account Type").
                Options(
                    huh.NewOption("Personal", "personal"),
                    huh.NewOption("Business", "business"),
                ).
                Value(&userType),
        ),

        huh.NewGroup(
            huh.NewInput().
                Title("Your Name").
                Value(&personalName),
        ).WithHideFunc(func() bool {
            return userType != "personal"
        }),

        huh.NewGroup(
            huh.NewInput().
                Title("Company Name").
                Value(&companyName),
        ).WithHideFunc(func() bool {
            return userType != "business"
        }),
    )

    err := form.Run()
    if err != nil {
        log.Fatal(err)
    }

    if userType == "personal" {
        fmt.Printf("Welcome, %s!\n", personalName)
    } else {
        fmt.Printf("Welcome, %s!\n", companyName)
    }
}
```

----------

## Bubbletea Integration
Integrate huh form into bubbletea app
```go
package main

import (
    "fmt"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/huh"
)

type model struct {
    form   *huh.Form
    result string
}

func initialModel() model {
    var name string

    form := huh.NewForm(
        huh.NewGroup(
            huh.NewInput().
                Title("What's your name?").
                Value(&name),
        ),
    ).WithWidth(45).WithHeight(8)

    return model{form: form}
}

func (m model) Init() tea.Cmd {
    return m.form.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c":
            return m, tea.Quit
        }
    }

    form, cmd := m.form.Update(msg)
    if f, ok := form.(*huh.Form); ok {
        m.form = f
    }

    if m.form.State == huh.StateCompleted {
        return m, tea.Quit
    }

    return m, cmd
}

func (m model) View() string {
    if m.form.State == huh.StateCompleted {
        return fmt.Sprintf("Form completed!\n")
    }
    return m.form.View()
}

func main() {
    tea.NewProgram(initialModel()).Run()
}
```

----------

## Quick Single Prompt
Standalone single field prompt
```go
package main

import (
    "fmt"
    "log"
    "github.com/charmbracelet/huh"
)

func main() {
    var proceed bool

    err := huh.NewConfirm().
        Title("Continue with installation?").
        Value(&proceed).
        Run()

    if err != nil {
        log.Fatal(err)
    }

    if proceed {
        fmt.Println("Installing...")
    } else {
        fmt.Println("Installation cancelled")
    }
}
```

----------

## Loading Spinner
Background task with spinner
```go
package main

import (
    "context"
    "time"
    "github.com/charmbracelet/huh/spinner"
)

func main() {
    action := func() {
        time.Sleep(3 * time.Second)
    }

    err := spinner.New().
        Title("Loading data...").
        Action(action).
        Run()

    if err != nil {
        panic(err)
    }

    println("Data loaded successfully!")
}
```

----------

## Accessible Form
Form with accessibility features
```go
package main

import (
    "fmt"
    "log"
    "github.com/charmbracelet/huh"
)

func main() {
    var name string

    form := huh.NewForm(
        huh.NewGroup(
            huh.NewInput().
                Title("Name").
                Description("Enter your full name").
                Placeholder("John Doe").
                Value(&name),
        ),
    ).WithAccessible(true)

    err := form.Run()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Name: %s\n", name)
}
```