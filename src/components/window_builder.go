package components

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type WindowBuilder struct {
	window   fyne.Window
	ib       *ImageBrowser
	contents []fyne.CanvasObject
	size     fyne.Size
	canvas   *DrawCanvas
}

func NewWindowBuilder(title string, a fyne.App) *WindowBuilder {
	initPath, _ := os.Getwd()

	result := &WindowBuilder{
		window:   a.NewWindow(title),
		contents: make([]fyne.CanvasObject, 0),
		ib:       NewImageBrowser(initPath),
	}

	return result
}

func (wb *WindowBuilder) WithSize(width, height float32) *WindowBuilder {
	wb.size = fyne.NewSize(width, height)
	wb.window.Resize(wb.size)
	return wb
}

func (wb *WindowBuilder) AddContent(content fyne.CanvasObject) *WindowBuilder {
	wb.contents = append(wb.contents, content)
	return wb
}

func (wb *WindowBuilder) WithOpenFolderButton() *WindowBuilder {
	openFolderButton := widget.NewButton("Open Folder", func() { wb.onOpenFolderButtonClicked() })
	wb.AddContent(openFolderButton)
	return wb
}

func (wb *WindowBuilder) onOpenFolderButtonClicked() {
	fd := dialog.NewFolderOpen(func(lu fyne.ListableURI, err error) {
		if err != nil || lu == nil {
			return
		}
		wb.ib.UpdatePath(lu.Path())
		wb.setContent()

	}, wb.window)

	uri, err := storage.ListerForURI(storage.NewFileURI(wb.ib.path))
	if err == nil {
		fd.SetLocation(uri)
	}
	fd.Show()
}

func (wb *WindowBuilder) setContent() {
	containers := container.NewHBox()
	for _, obj := range wb.contents {
		containers.Add(obj)
	}

	if wb.ib.DirCount() > 0 {
		wb.canvas = NewDrawCanvas(wb.size)
		var img *canvas.Image = wb.ib.GetCurrent()
		img.SetMinSize(wb.size)
		containers.Add(container.NewStack(img, wb.canvas))
		wb.window.Canvas().SetOnTypedKey(wb.KeyPressedEvent)
		wb.window.Resize(wb.size)

	}

	wb.window.SetContent(containers)
	wb.window.Content().Refresh()
}

func (wb *WindowBuilder) Build() fyne.Window {
	wb.setContent()
	return wb.window
}

func (wb *WindowBuilder) KeyPressedEvent(event *fyne.KeyEvent) {

	fmt.Printf("Event name : %s, Event type : %v\n", event.Name, event.Physical)

	switch event.Name {
	case fyne.KeyLeft:
		wb.ib.Previous()
		wb.setContent()
		return
	case fyne.KeyRight:
		wb.ib.Next()
		wb.setContent()

	default:
		return
	}
}
