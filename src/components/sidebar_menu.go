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
	btnMapping   map[common.SidebarButtonType]func()
	btnIconMap   map[common.SidebarButtonType]fyne.Resource
}

func NewSidebarMenu(ib *ImageBrowser) *SidebarMenu {
	res := &SidebarMenu{
		buttons: container.NewVBox(),
		ib:      ib,
	}
	res.btnMapping = map[common.SidebarButtonType]func(){
		common.IncBtn: res.onIncreaseBrushButton,
		common.DecBtn: res.onDecreaseBrushButton,
		common.Toggle: res.onToggleBrushClicked,
	}
	res.btnIconMap = map[common.SidebarButtonType]fyne.Resource{
		common.IncBtn: theme.ContentAddIcon(),
		common.DecBtn: theme.ContentRemoveIcon(),
	}

	return res
}

func (wb *SidebarMenu) Build() *fyne.Container {
	return container.NewCenter(wb.buttons)
}

func (wb *SidebarMenu) WithButton(btnType common.SidebarButtonType) *SidebarMenu {
	if _, ok := wb.btnMapping[btnType]; !ok {
		return wb
	}
	if btnType == common.Toggle {
		wb.withToggleBrushBtn()
	} else {
		plusButton := widget.NewButton("", wb.btnMapping[btnType])
		plusButton.Icon = wb.btnIconMap[btnType]
		plusButton.Resize(common.DefaultIconSize)
		wb.buttons.Add(plusButton)

	}
	return wb
}

func (wb *SidebarMenu) withToggleBrushBtn() *SidebarMenu {
	wb.toggleButton = widget.NewButton("", wb.onToggleBrushClicked)
	wb.setToggleIcon()
	wb.toggleButton.Resize(common.DefaultIconSize)
	wb.buttons.Add(wb.toggleButton)
	return wb
}

func (wb *SidebarMenu) onIncreaseBrushButton() {
	wb.ib.Inc()
}

func (wb *SidebarMenu) onDecreaseBrushButton() {
	wb.ib.Dec()
}

func (wb *SidebarMenu) onToggleBrushClicked() {
	wb.ib.ToogleBrush()
	wb.setToggleIcon()
	wb.buttons.Refresh()
}

func (wb *SidebarMenu) setToggleIcon() {
	if wb.ib.canvas.toogleBrush {
		wb.toggleButton.Icon = theme.ColorChromaticIcon()
	} else {
		wb.toggleButton.Icon = theme.ColorAchromaticIcon()
	}
	wb.buttons.Refresh()
}
