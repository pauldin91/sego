package common

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type ButtonType string
type BottomButtonType string

const (
	OpenBtn ButtonType = "Open Folder"
	LoadBtn ButtonType = "Load File"

	ClearBtn BottomButtonType = "X"
	Toggle   BottomButtonType = "toggle"
	SaveBtn  BottomButtonType = "S"
	IncBtn   BottomButtonType = "+"
	DecBtn   BottomButtonType = "-"
	ColorBtn BottomButtonType = "O"
)

var DefaultCanvasSize fyne.Size = fyne.NewSize(600, 400)
var DefaultButtonSize fyne.Size = fyne.NewSize(60, 40)
var DefaultVButtonSize fyne.Size = fyne.NewSize(20, 60)
var DefaultIconSize fyne.Size = fyne.NewSize(20, 20)
var DefaultIncBrushSize fyne.Size = fyne.NewSize(10, 10)
var DefaultPaddingSize fyne.Size = fyne.NewSize(3, 3)

var DefaultPaintColor color.RGBA = color.RGBA{R: 41, G: 111, B: 246, A: 185}
var DefaultTransparrentColor color.RGBA = color.RGBA{R: 0, G: 0, B: 0, A: 0}

var imageExts = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
	".bmp":  true,
	".webp": true,
	".tiff": true,
}
