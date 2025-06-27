package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/pauldin91/sego/src/common"
)

type SidebarMenu struct {
	buttons []fyne.CanvasObject
	ib      *ImageBrowser
}

func NewSidebarMenu(ib *ImageBrowser) *SidebarMenu {
	return &SidebarMenu{
		buttons: make([]fyne.CanvasObject, 0),
		ib:      ib,
	}
}

func (wb *SidebarMenu) getBrushButtons() *fyne.Container {
	iconButtons := container.NewVBox()

	plusButton := widget.NewButton("", wb.onIncreaseBrushButton)
	plusButton.Icon = theme.ContentAddIcon()
	plusButton.Resize(common.DefaultIconSize)
	minusButton := widget.NewButton("", wb.onDecreaseBrushButton)
	minusButton.Icon = theme.ContentRemoveIcon()
	minusButton.Resize(common.DefaultIconSize)

	iconButtons.Add(plusButton)
	iconButtons.Add(minusButton)
	return container.NewCenter(iconButtons)
}
func (wb *SidebarMenu) onIncreaseBrushButton() {
	wb.ib.Inc()
}

func (wb *SidebarMenu) onDecreaseBrushButton() {
	wb.ib.Dec()
}
