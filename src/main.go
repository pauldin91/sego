package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"github.com/pauldin91/sego/src/components"
)

var (
	wWidth  float32 = 640
	wHeight float32 = 480
)

func main() {
	a := app.NewWithID("image.viewer." + time.Now().String())

	w := components.NewWindowBuilder("Hakuna Matata", a).
		OfSize(wWidth, wHeight).
		WithOpenFileButton().
		Build()

	w.ShowAndRun()
	w.Close()
}
