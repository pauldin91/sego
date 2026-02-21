package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"github.com/pauldin91/sego/components"
)

func main() {
	a := app.NewWithID("image.viewer." + time.Now().String())

	w := components.
		NewWindowBuilder("Hakuna Matata", a).
		WithBottomMenu().
		WithDefaultCanvas().
		WithMainMenu().
		Build()

	w.ShowAndRun()
	w.Close()
}
