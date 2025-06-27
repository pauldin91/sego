package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/pauldin91/sego/src/common"
)

type SidebarMenu struct {
	buttons      *fyne.Container
	ib           *ImageBrowser
	toggleButton *widget.Button
}

func NewSidebarMenu(ib *ImageBrowser) *SidebarMenu {
	return &SidebarMenu{
		buttons: container.NewVBox(),
		ib:      ib,
	}
}

func (wb *SidebarMenu) Build() *fyne.Container {

	return container.NewCenter(wb.buttons)
}

func (wb *SidebarMenu) WithIncreaseBrushSizeButton() *SidebarMenu {
	plusButton := widget.NewButton("", wb.onIncreaseBrushButton)
	plusButton.Icon = theme.ContentAddIcon()
	plusButton.Resize(common.DefaultIconSize)
	wb.buttons.Add(plusButton)

	return wb
}

func (wb *SidebarMenu) onIncreaseBrushButton() {
	wb.ib.Inc()
}

func (wb *SidebarMenu) WithDecreaseBrushSizeButton() *SidebarMenu {
	minusButton := widget.NewButton("", wb.onDecreaseBrushButton)
	minusButton.Icon = theme.ContentRemoveIcon()
	minusButton.Resize(common.DefaultIconSize)
	wb.buttons.Add(minusButton)

	return wb
}

func (wb *SidebarMenu) onDecreaseBrushButton() {
	wb.ib.Dec()
}

func (wb *SidebarMenu) WithToggleBrushButton() *SidebarMenu {
	wb.toggleButton = widget.NewButton("", wb.onToggleBrushClicked)
	wb.updateToogleIcon()
	wb.toggleButton.Resize(common.DefaultIconSize)
	wb.buttons.Add(wb.toggleButton)

	return wb
}

func (wb *SidebarMenu) onToggleBrushClicked() {
	wb.ib.ToogleBrush()
	wb.updateToogleIcon()
	wb.buttons.Refresh()
}

func (wb *SidebarMenu) updateToogleIcon() {
	if wb.ib.canvas.toogleBrush {
		wb.toggleButton.Icon = theme.ColorChromaticIcon()
	} else {
		wb.toggleButton.Icon = theme.ColorAchromaticIcon()
	}
}
