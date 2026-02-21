package components

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/pauldin91/sego/utils"
)

type BottomMenu struct {
	buttons      *fyne.Container
	ib           *ImageViewer
	parent       fyne.Window
	toggleButton *widget.Button
	btnMapping   map[utils.BottomButtonType]func()
	btnIconMap   map[utils.BottomButtonType]fyne.Resource
	cfg          utils.Config
}

func NewBottomMenu(cfg utils.Config, ib *ImageViewer, parent fyne.Window) *BottomMenu {
	res := &BottomMenu{
		buttons: container.NewHBox(),
		ib:      ib,
		parent:  parent,
		cfg:     cfg,
	}
	res.btnMapping = map[utils.BottomButtonType]func(){
		utils.IncBtn:   res.onIncreaseBrushButton,
		utils.DecBtn:   res.onDecreaseBrushButton,
		utils.Toggle:   res.onToggleBrushClicked,
		utils.ColorBtn: res.onColorPickerClicked,
		utils.SaveBtn:  res.ib.Save,
		utils.ClearBtn: res.ib.Clear,
	}
	res.btnIconMap = map[utils.BottomButtonType]fyne.Resource{
		utils.IncBtn:   theme.ContentAddIcon(),
		utils.DecBtn:   theme.ContentRemoveIcon(),
		utils.ColorBtn: theme.ColorPaletteIcon(),
		utils.SaveBtn:  theme.DocumentSaveIcon(),
		utils.ClearBtn: theme.ContentClearIcon(),
	}

	return res
}

func (wb *BottomMenu) Build() *fyne.Container {
	return container.NewCenter(wb.buttons)
}

func (wb *BottomMenu) WithButtons(btnTypes ...utils.BottomButtonType) *BottomMenu {
	var vBox = container.NewVBox()
	for _, btnType := range btnTypes {
		if _, ok := wb.btnMapping[btnType]; !ok {
			return wb
		}
		if btnType == utils.Toggle {
			wb.toggleButton = widget.NewButton("", wb.onToggleBrushClicked)
			wb.setToggleIcon()
			vBox.Add(wb.toggleButton)
		} else {
			plusButton := widget.NewButton("", wb.btnMapping[btnType])
			plusButton.Icon = wb.btnIconMap[btnType]
			vBox.Add(plusButton)
		}
	}
	vBox.Resize(utils.DefaultIconSize)
	wb.buttons.Add(vBox)
	return wb
}

func (wb *BottomMenu) WithButton(btnType utils.BottomButtonType) *BottomMenu {
	if _, ok := wb.btnMapping[btnType]; !ok {
		return wb
	}
	if btnType == utils.Toggle {
		wb.withToggleBrushBtn()
	} else {
		plusButton := widget.NewButton("", wb.btnMapping[btnType])
		plusButton.Icon = wb.btnIconMap[btnType]
		plusButton.Resize(utils.DefaultIconSize)
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
