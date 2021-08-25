package forms

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"log"
)

func NewConnectionForm() fyne.CanvasObject {
	hostEntry := widget.NewFormItem("Host:", widget.NewEntry())
	portEntry := widget.NewFormItem("Port:", widget.NewEntry())

	form := widget.NewForm(hostEntry, portEntry)
	form.SubmitText = "Connect"
	form.OnSubmit = func() {
		log.Println("HELLO!")
	}

	return form
}
