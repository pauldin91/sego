package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type WindowBuilder struct {
	w        fyne.Window
	contents []fyne.CanvasObject
}

func NewWindowBuilder(title string, a fyne.App) WindowBuilder {
	return WindowBuilder{
		w:        a.NewWindow(title),
		contents: make([]fyne.CanvasObject, 0),
	}
}

func (wb *WindowBuilder) OfSize(width, height float32) *WindowBuilder {
	wb.w.Resize(fyne.NewSize(width, height))
	return wb
}

func (wb *WindowBuilder) AddContent(content fyne.CanvasObject) *WindowBuilder {
	wb.contents = append(wb.contents, content)
	return wb
}

func (wb *WindowBuilder) WithDefaultLayout(ctxImage ImageContainer) *WindowBuilder {
	openFileDialog := widget.NewButton("Open Image", func() { wb.tapped(ctxImage) })
	wb.AddContent(openFileDialog)
	wb.AddContent(ctxImage.content)
	return wb
}

func (wb *WindowBuilder) tapped(ctxImage ImageContainer) {
	fd := dialog.NewFileOpen(ctxImage.UpdateContent, wb.w)
	fd.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg"}))
	fd.Show()
}

func (wb *WindowBuilder) Build() fyne.Window {
	containers := container.NewHBox()
	for _, obj := range wb.contents {
		containers.Add(obj)
	}
	wb.w.SetContent(containers)
	return wb.w
}
