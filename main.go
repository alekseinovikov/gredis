package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func Sum(a int, b int) int {
	return a + b
}

func main() {
	app := app.New()

	w := app.NewWindow("Hello")
	w.SetContent(widget.NewLabel("Hello Fyne!"))

	w.ShowAndRun()
}
