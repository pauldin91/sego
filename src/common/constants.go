package common

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type BrowserStatus string

const (
	DefaultBrushSize   float64       = 15.0
	DefaultBrushChange float64       = 1.0
	DefaultMaskPreffix string        = "mask_"
	DefaultMaskDir     string        = "masks"
	DefaultResourceDir string        = "../resources"
	ImageChanged       BrowserStatus = "Changed"
)

var DefaultCanvasSize fyne.Size = fyne.NewSize(600, 400)
var DefaultButtonSize fyne.Size = fyne.NewSize(60, 40)
var DefaultIconSize fyne.Size = fyne.NewSize(40, 40)
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
