package common

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"golang.org/x/image/draw"
)

func DefaultBlankImage(size fyne.Size) (*canvas.Image, *image.RGBA) {
	rgba := image.NewRGBA(image.Rect(0, 0, int(size.Width), int(size.Height)))
	draw.Draw(rgba, rgba.Bounds(), &image.Uniform{color.Transparent}, image.Point{}, draw.Src)

	currImg := canvas.NewImageFromImage(rgba)
	currImg.FillMode = canvas.ImageFillContain
	currImg.Resize(size)
	return currImg, rgba
}

func ListDir(dir string) []string {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil
	}
	paths := []string{}
	for _, entry := range entries {
		if entry.Type().IsRegular() && isImage(entry.Name()) {
			paths = append(paths, filepath.Join(dir, entry.Name()))
		}
	}
	return paths
}

func isImage(file string) bool {
	ext := strings.ToLower(filepath.Ext(file))
	return imageExts[ext]
}

func ScaleImage(img image.Image, dstSize fyne.Size) *image.RGBA {
	width := int(dstSize.Width)
	height := int(dstSize.Height)

	scaled := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.CatmullRom.Scale(scaled, scaled.Bounds(), img, img.Bounds(), draw.Over, nil)

	return scaled
}
