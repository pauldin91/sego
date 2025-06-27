package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/pauldin91/sego/src/common"
)

type WindowBuilder struct {
	window     fyne.Window
	ib         *ImageBrowser
	contents   []fyne.CanvasObject
	buttons    []fyne.CanvasObject
	canvasSize fyne.Size
}

func NewWindowBuilder(size fyne.Size, title string, a fyne.App) *WindowBuilder {

	result := &WindowBuilder{
		window:     a.NewWindow(title),
		contents:   make([]fyne.CanvasObject, 0),
		buttons:    make([]fyne.CanvasObject, 0),
		ib:         NewImageBrowser(),
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
	openFolderButton := widget.NewButton("Open Folder", wb.onOpenFolderButtonClicked)
	openFolderButton.Resize(fyne.NewSize(60, 30))
	wb.buttons = append(wb.buttons, openFolderButton)
	return wb
}

func (wb *WindowBuilder) WithLoadButton() *WindowBuilder {
	loadFileButton := widget.NewButton("Load File", wb.onLoadFileButtonClicked)
	loadFileButton.Resize(common.DefaultButtonSize)
	wb.buttons = append(wb.buttons, loadFileButton)

	return wb
}

func (wb *WindowBuilder) onLoadFileButtonClicked() {
	fd := dialog.NewFileOpen(func(lu fyne.URIReadCloser, err error) {
		if err != nil || lu == nil {
			return
		}
		wb.ib.loadContent(lu.URI().Path())
		wb.setContent()

	}, wb.window)

	uri, err := storage.ListerForURI(storage.NewFileURI(wb.ib.path))
	if err == nil {
		fd.SetLocation(uri)
	}
	fd.Show()
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

	buttonContainer := container.NewHBox()
	for _, b := range wb.buttons {
		buttonContainer.Add(b)
	}

	containers := container.NewVBox()
	containers.Add(wb.ib)
	containers.Add(container.NewCenter(buttonContainer))

	wb.window.SetContent(containers)
	wb.window.Canvas().Focus(wb.ib)

	wb.window.Resize(wb.canvasSize)
	wb.window.SetTitle(wb.ib.title)
	wb.window.Content().Refresh()
}

func (wb *WindowBuilder) Build() fyne.Window {
	wb.setContent()
	return wb.window
}
