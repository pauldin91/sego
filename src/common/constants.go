package common

import "fyne.io/fyne/v2"

const (
	BrushSize float64 = 15.0
)

type BrowserStatus string

const (
	ImageChanged BrowserStatus = "Changed"
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
