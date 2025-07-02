package controls

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/pauldin91/sego/src/common"
	"github.com/pauldin91/sego/src/components/viewer"
)

type BottomMenu struct {
	buttons      *fyne.Container
	ib           *viewer.ImageBrowser
	parent       fyne.Window
	toggleButton *widget.Button
	btnMapping   map[common.BottomButtonType]func()
	btnIconMap   map[common.BottomButtonType]fyne.Resource
	cfg          common.Config
}

func NewBottomMenu(cfg common.Config, ib *viewer.ImageBrowser, parent fyne.Window) *BottomMenu {
	res := &BottomMenu{
		buttons: container.NewHBox(),
		ib:      ib,
		parent:  parent,
		cfg:     cfg,
	}
	res.btnMapping = map[common.BottomButtonType]func(){
		common.IncBtn:   res.onIncreaseBrushButton,
		common.DecBtn:   res.onDecreaseBrushButton,
		common.Toggle:   res.onToggleBrushClicked,
		common.ColorBtn: res.onColorPickerClicked,
		common.SaveBtn:  res.ib.Save,
		common.ClearBtn: res.ib.Clear,
	}
	res.btnIconMap = map[common.BottomButtonType]fyne.Resource{
		common.IncBtn:   theme.ContentAddIcon(),
		common.DecBtn:   theme.ContentRemoveIcon(),
		common.ColorBtn: theme.ColorPaletteIcon(),
		common.SaveBtn:  theme.DocumentSaveIcon(),
		common.ClearBtn: theme.ContentClearIcon(),
	}

	return res
}

func (wb *BottomMenu) Build() *fyne.Container {
	return container.NewCenter(wb.buttons)
}

func (wb *BottomMenu) WithButtons(btnTypes ...common.BottomButtonType) *BottomMenu {
	var vBox = container.NewVBox()
	for _, btnType := range btnTypes {
		if _, ok := wb.btnMapping[btnType]; !ok {
			return wb
		}
		if btnType == common.Toggle {
			wb.toggleButton = widget.NewButton("", wb.onToggleBrushClicked)
			wb.setToggleIcon()
			vBox.Add(wb.toggleButton)
		} else {
			plusButton := widget.NewButton("", wb.btnMapping[btnType])
			plusButton.Icon = wb.btnIconMap[btnType]
			vBox.Add(plusButton)
		}
	}
	vBox.Resize(common.DefaultIconSize)
	wb.buttons.Add(vBox)
	return wb
}

func (wb *BottomMenu) WithButton(btnType common.BottomButtonType) *BottomMenu {
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

func (wb *BottomMenu) WithFloatWidget(onChanged func(s string)) *BottomMenu {

	entry := widget.NewEntry()
	entry.SetText(fmt.Sprintf("%.2f", wb.cfg.DefaultBrushSize))
	entry.OnChanged = onChanged
	wb.buttons.Add(entry)
	return wb
}

func (wb *BottomMenu) withToggleBrushBtn() *BottomMenu {
	wb.toggleButton = widget.NewButton("", wb.onToggleBrushClicked)
	wb.setToggleIcon()
	wb.buttons.Add(wb.toggleButton)
	return wb
}

func (wb *BottomMenu) onColorPickerClicked() {
	cd := dialog.NewColorPicker("Annotation Color", "Choose an annotation color", func(c color.Color) {
		wb.ib.ChooseColor(c)
		wb.setToggleIcon()
	}, wb.parent)
	cd.Advanced = true
	cd.Show()
}

func (wb *BottomMenu) onIncreaseBrushButton() {
	wb.ib.IncBrush()
}

func (wb *BottomMenu) onDecreaseBrushButton() {
	wb.ib.DecBrush()
}

func (wb *BottomMenu) onToggleBrushClicked() {
	wb.ib.Toggle()
	wb.setToggleIcon()
	wb.buttons.Refresh()
}

func (wb *BottomMenu) setToggleIcon() {
	if wb.ib.GetToggle() {
		wb.toggleButton.Icon = theme.ColorChromaticIcon()
	} else {
		wb.toggleButton.Icon = theme.ColorAchromaticIcon()
	}
	wb.buttons.Refresh()
}
