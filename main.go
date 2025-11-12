package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()
	w := a.NewWindow("Bellypepper MX!!")

	input := widget.NewEntry()
	input.SetPlaceHolder("Ingrese su Nombre ...")

	hello := widget.NewLabel("")

	send := widget.NewButton("Enviar", func() {
		hello.SetText("Hola " + input.Text)
	})

	content := container.NewVBox(input, send, hello)

	w.SetContent(content)
	w.ShowAndRun()

}
