package utils

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"fyne.io/fyne/v2"
	"golang.org/x/image/draw"
)

func ScaleImage(img image.Image, dstSize fyne.Size) *image.RGBA {
	width := int(dstSize.Width)
	height := int(dstSize.Height)

	scaled := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.CatmullRom.Scale(scaled, scaled.Bounds(), img, img.Bounds(), draw.Over, nil)

	return scaled
}

func SaveMask(rgba *image.RGBA, filename string) {

	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("error saving the image %s : %v\n", filename, err)
		return
	}
	defer file.Close()

	png.Encode(file, rgba)
}
