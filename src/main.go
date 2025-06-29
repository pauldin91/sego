package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"github.com/pauldin91/sego/src/components"
)

func main() {
	a := app.NewWithID("image.viewer." + time.Now().String())

	w := components.
		NewWindowBuilder("Hakuna Matata", a).
		WithSidebarMenu().
		WithDefaultCanvas().
		WithBottomMenu().
		Build()

	w.ShowAndRun()
	w.Close()
}
