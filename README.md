# Flinch

A collection of terminal-based widgets for richer Golang CLI apps.

Ships with a library to build your own widgets/TUIs too.

**Warning**: This module is experimental right now.

### Input

![](example_input.png)

```golang
name, _ := widgets.Input("Enter your name...")
```

### Password Input

![](example_password.png)

```golang 
password, _ := widgets.PasswordInput("Enter your password...")
```

### List Selection

![](example_list.png)

```golang
_, item, err := widgets.ListSelect(
    "Select an environment...",
    []string{
        "Development",
        "Test",
        "Staging",
        "Production",
    },
)
```

### Multi List Selection

![](example_multi.png)

```golang
_, items, err := widgets.MultiSelect(
    "Select an option...",
    options,
)
```

(scrollbars appear for long lists)

### Confirmation

![](example_confirm.png)

```golang
userConfirmed, _ := widgets.Confirm("Are you sure?")
```

## Build Your Own Widgets

Check out the [widgets](widgets) package for some inspiration.
