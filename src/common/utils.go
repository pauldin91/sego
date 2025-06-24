package common

import (
	"image"
	"image/color"
	"image/draw"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func DefaultBlankImage(size fyne.Size) (*canvas.Image, *image.RGBA) {
	rgba := image.NewRGBA(image.Rect(0, 0, int(size.Width), int(size.Height)))
	draw.Draw(rgba, rgba.Bounds(), &image.Uniform{color.Transparent}, image.Point{}, draw.Src)

	currImg := canvas.NewImageFromImage(rgba)
	currImg.FillMode = canvas.ImageFillContain
	currImg.SetMinSize(size)
	return currImg, rgba
}
