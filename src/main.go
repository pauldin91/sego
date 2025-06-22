package main

import (
	"image"
	"io"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

var (
	w            fyne.Window
	imgContainer *canvas.Image
	imgLabel     *widget.Label
)

func main() {
	a := app.NewWithID("image.viewer." + time.Now().String())
	w = a.NewWindow("Hakuna Matata")
	w.Resize(fyne.NewSize(640, 480))

	// Initial blank image
	blank := image.NewRGBA(image.Rect(0, 0, 640, 480))

	imgContainer = canvas.NewImageFromImage(blank)
	imgContainer.FillMode = canvas.ImageFillContain
	imgContainer.SetMinSize(fyne.NewSize(640, 400))

	imgLabel = widget.NewLabel("No image loaded")
	openBtn := getDialogButton(w)

	content := container.NewVBox(
		openBtn,
		imgLabel,
		imgContainer,
	)
	w.SetContent(content)
	w.ShowAndRun()
}

func getDialogButton(win fyne.Window) *widget.Button {
	return widget.NewButton("Open Image", func() {
		fd := dialog.NewFileOpen(selectImage, win)

		fd.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg"}))
		fd.Show()
	})
}

func selectImage(reader fyne.URIReadCloser, err error) {
	if err != nil {
		dialog.ShowError(err, w)
		return
	}
	if reader == nil {
		return // user cancelled
	}

	data, err := io.ReadAll(reader)
	if err != nil {
		dialog.ShowError(err, w)
		return
	}
	res := fyne.NewStaticResource(reader.URI().Name(), data)

	imgContainer.Resource = res
	imgContainer.Refresh()
	imgLabel.SetText("Loaded: " + reader.URI().Name())

}
