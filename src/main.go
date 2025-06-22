package main

import (
	"encoding/hex"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var topWindow fyne.Window
var preferenceCurrentTutorial = []byte(time.Now().UTC().String())

func main() {
	app := app.NewWithID(hex.EncodeToString(preferenceCurrentTutorial))
	w := app.NewWindow("Hakuna Mattata")
	w.SetFixedSize(true)
	// message := widget.NewLabel("Welcome")
	// button := widget.NewButton("Update", func() {
	// 	formatted := time.Now().Format("Time: 03:04:05")
	// 	message.SetText(formatted)
	// })

	// w.SetContent(container.NewVBox(message, button))
	w.Resize(fyne.NewSize(640, 460))
	w.ShowAndRun()
	//w.Close()
}
