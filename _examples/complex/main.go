package main

import (
    "github.com/liamg/flinch/components"
    "github.com/liamg/flinch/core"
    window2 "github.com/liamg/flinch/window"
)

func main() {

    window, err := window2.New()
    if err != nil {
        panic(err)
    }

    rows := components.NewRowLayout()
    frame := components.NewFrame(rows)
    limiter := components.NewLimiter(frame)
    limiter.WithPercentageLimitOnWidth(60)

    checkbox1 := components.NewCheckbox("Enable A", false)
    checkbox2 := components.NewCheckbox("Enable B", false)

    rows.Add(checkbox1)
    rows.Add(checkbox2)

    window.WithJustification(core.JustifyCenter)
    window.Add(limiter)

    if err := window.Show(); err != nil {
        panic(err)
    }
}
