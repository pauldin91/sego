package utils

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"golang.org/x/image/draw"
)

func DefaultBlankImage(size fyne.Size) *image.RGBA {
	rgba := image.NewRGBA(image.Rect(0, 0, int(size.Width), int(size.Height)))
	draw.Draw(rgba, rgba.Bounds(), &image.Uniform{color.Transparent}, image.Point{}, draw.Src)
	return rgba
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

func SaveMask(rgba *image.RGBA, filename string) {

	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("error saving the image %s : %v\n", filename, err)
		return
	}
	defer file.Close()

	png.Encode(file, rgba)
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
