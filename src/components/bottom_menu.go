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
	buttons    *fyne.Container
	ib         *ImageBrowser
	parent     fyne.Window
	btnMapping map[common.ButtonType]func()
}

func NewBottomMenu(ib *ImageBrowser, parent fyne.Window) *BottomMenu {
	res := &BottomMenu{
		buttons: container.NewHBox(),
		ib:      ib,
		parent:  parent,
	}
	res.btnMapping = map[common.ButtonType]func(){
		common.OpenBtn:  res.onOpenFolderButtonClicked,
		common.LoadBtn:  res.onLoadFileButtonClicked,
		common.ClearBtn: res.onClearButtonClicked,
	}
	return res
}

func (wb *BottomMenu) Build() *fyne.Container {
	return container.NewCenter(wb.buttons)
}

func (wb *BottomMenu) WithButton(btnType common.ButtonType) *BottomMenu {
	if _, ok := wb.btnMapping[btnType]; !ok {
		return wb
	}
	button := widget.NewButton(string(btnType), wb.btnMapping[btnType])
	button.Resize(fyne.NewSize(60, 30))
	button.Show()
	wb.buttons.Add(button)
	return wb
}

func (wb *BottomMenu) onOpenFolderButtonClicked() {
	fd := dialog.NewFolderOpen(func(lu fyne.ListableURI, err error) {
		if err != nil || lu == nil {
			return
		}
		wb.ib.UpdateImage(lu.Path())

	}, wb.parent)
	wb.setLocation(fd)
}

func (wb *BottomMenu) onLoadFileButtonClicked() {
	fd := dialog.NewFileOpen(func(lu fyne.URIReadCloser, err error) {
		if err != nil || lu == nil {
			return
		}
		wb.ib.LoadContent(lu.URI().Path())
	}, wb.parent)

	wb.setLocation(fd)
}

func (wb *BottomMenu) onClearButtonClicked() {
	wb.ib.Clear()
}

func (wb *BottomMenu) setLocation(fd *dialog.FileDialog) {
	uri, err := storage.ListerForURI(storage.NewFileURI(wb.ib.fb.GetPath()))
	if err == nil {
		fd.SetLocation(uri)
	}
	fd.Show()
}
