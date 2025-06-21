package main

import (
	"encoding/hex"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.NewWithID(hex.EncodeToString([]byte(time.Now().UTC().String())))
	w := app.NewWindow("Hakuna Mattata")
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(640, 460))
	w.ShowAndRun()
	w.Close()
}
