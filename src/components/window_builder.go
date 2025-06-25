package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type WindowBuilder struct {
	window     fyne.Window
	ib         *ImageBrowser
	dc         *DrawableCanvas
	contents   []fyne.CanvasObject
	canvasSize fyne.Size
}

func NewWindowBuilder(size fyne.Size, title string, a fyne.App) *WindowBuilder {

	fileChan := make(chan string)
	saveCompleted := make(chan bool)
	result := &WindowBuilder{
		window:     a.NewWindow(title),
		contents:   make([]fyne.CanvasObject, 0),
		ib:         NewImageBrowser(fileChan, saveCompleted),
		dc:         NewDrawableCanvas(fileChan, saveCompleted),
		canvasSize: size,
	}
	result.window.Resize(result.canvasSize)

	return result
}

func (wb *WindowBuilder) AddContent(content fyne.CanvasObject) *WindowBuilder {
	wb.contents = append(wb.contents, content)
	return wb
}

func (wb *WindowBuilder) WithOpenFolderButton() *WindowBuilder {
	openFolderButton := widget.NewButton("Open Folder", func() { wb.onOpenFolderButtonClicked() })
	openFolderButton.Resize(fyne.NewSize(60, 30))
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
	containers := container.NewVBox()
	containers.Add(container.NewStack(wb.ib, wb.dc))

	for _, obj := range wb.contents {
		containers.Add(obj)
	}

	wb.window.SetContent(containers)
	wb.window.Canvas().Focus(wb.ib)

	wb.window.Resize(wb.canvasSize)
	wb.window.Content().Refresh()
}

func (wb *WindowBuilder) Build() fyne.Window {
	wb.setContent()
	return wb.window
}
