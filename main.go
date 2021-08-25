package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/alekseinovikov/gredis/forms"
)

func Sum(a int, b int) int {
	return a + b
}

func main() {
	myApp := app.New()
	w := myApp.NewWindow("GRedis")

	connectionForm := forms.NewConnectionForm()

	w.Resize(fyne.Size{
		Width:  400,
		Height: 400,
	})

	w.SetContent(connectionForm)
	w.ShowAndRun()
}
