package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"github.com/pauldin91/sego/src/common"
	"github.com/pauldin91/sego/src/components"
)

func main() {
	a := app.NewWithID("image.viewer." + time.Now().String())

	w := components.
		NewWindowBuilder(common.Size, "Hakuna Matata", a).
		WithOpenFolderButton().
		Build()

	w.ShowAndRun()
	w.Close()
}
