package common

import "fyne.io/fyne/v2"

type BrowserStatus string

const (
	DefaultBrushSize   float64       = 15.0
	DefaultMaskPreffix string        = "mask_"
	DefaultMaskDir     string        = "masks"
	DefaultResourceDir string        = "../resources"
	ImageChanged       BrowserStatus = "Changed"
	DefaultBrushChange float64       = 1.0
)

var Size fyne.Size = fyne.NewSize(600, 400)

var imageExts = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
	".bmp":  true,
	".webp": true,
	".tiff": true,
}
