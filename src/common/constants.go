package common

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type ButtonType string
type SidebarButtonType string

const (
	DefaultBrushSize   float64 = 15.0
	DefaultBrushChange float64 = 1.0
	DefaultMaskPreffix string  = "mask_"
	DefaultMaskDir     string  = "masks"
	DefaultResourceDir string  = "../resources"
)
const (
	OpenBtn  ButtonType = "Open Folder"
	LoadBtn  ButtonType = "Load File"
	ClearBtn ButtonType = "Clear Mask"

	Toggle SidebarButtonType = "toggle"
	IncBtn SidebarButtonType = "+"
	DecBtn SidebarButtonType = "-"
)

var DefaultCanvasSize fyne.Size = fyne.NewSize(600, 400)
var DefaultButtonSize fyne.Size = fyne.NewSize(60, 40)
var DefaultIconSize fyne.Size = fyne.NewSize(40, 40)
var DefaultPaddingSize fyne.Size = fyne.NewSize(3, 3)

var DefaultPaintColor color.RGBA = color.RGBA{R: 182, G: 245, B: 0, A: 127}

var imageExts = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
	".bmp":  true,
	".webp": true,
	".tiff": true,
}
