package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/pauldin91/sego/src/common"
)

type BottomMenu struct {
	buttons *fyne.Container
	ib      *ImageBrowser
	parent  fyne.Window
}

func NewBottomMenu(ib *ImageBrowser, parent fyne.Window) *BottomMenu {
	return &BottomMenu{
		buttons: container.NewHBox(),
		ib:      ib,
		parent:  parent,
	}
}

func (wb *BottomMenu) Build() *fyne.Container {
	return container.NewCenter(wb.buttons)
}

func (wb *BottomMenu) WithOpenFolderButton() *BottomMenu {
	openFolderButton := widget.NewButton("Open Folder", wb.onOpenFolderButtonClicked)
	openFolderButton.Resize(fyne.NewSize(60, 30))
	openFolderButton.Show()
	wb.buttons.Add(openFolderButton)
	return wb
}

func (wb *BottomMenu) onOpenFolderButtonClicked() {
	fd := dialog.NewFolderOpen(func(lu fyne.ListableURI, err error) {
		if err != nil || lu == nil {
			return
		}
		wb.ib.UpdatePath(lu.Path())

	}, wb.parent)

	uri, err := storage.ListerForURI(storage.NewFileURI(wb.ib.path))
	if err == nil {
		fd.SetLocation(uri)
	}
	fd.Show()
}

func (wb *BottomMenu) WithLoadButton() *BottomMenu {
	loadFileButton := widget.NewButton("Load File", wb.onLoadFileButtonClicked)
	loadFileButton.Resize(common.DefaultButtonSize)
	loadFileButton.Show()
	wb.buttons.Add(loadFileButton)

	return wb
}

func (wb *BottomMenu) onLoadFileButtonClicked() {
	fd := dialog.NewFileOpen(func(lu fyne.URIReadCloser, err error) {
		if err != nil || lu == nil {
			return
		}
		wb.ib.loadContent(lu.URI().Path())

	}, wb.parent)

	uri, err := storage.ListerForURI(storage.NewFileURI(wb.ib.path))
	if err == nil {
		fd.SetLocation(uri)
	}
	fd.Show()
}

func (wb *BottomMenu) WithClearButton() *BottomMenu {
	loadFileButton := widget.NewButton("Clear", wb.onClearButtonClicked)
	loadFileButton.Resize(common.DefaultButtonSize)
	loadFileButton.Show()
	wb.buttons.Add(loadFileButton)

	return wb
}

func (wb *BottomMenu) onClearButtonClicked() {
	wb.ib.Clear()
}
