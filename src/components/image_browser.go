package components

import (
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2/canvas"
)

var imageExts = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
	".bmp":  true,
	".webp": true,
	".tiff": true,
}

type ImageBrowser struct {
	path  string
	index int
	files []string
}

func NewImageBrowser(path string) ImageBrowser {
	var res = ImageBrowser{
		path:  path,
		index: 0,
	}
	res.files = res.listDir()
	return res
}

func (ib *ImageBrowser) Next() {
	ib.index++
}

func (ib *ImageBrowser) Previous() {
	ib.index--
}

func (ib *ImageBrowser) DirCount() int {
	return len(ib.files)
}

func (ib *ImageBrowser) UpdatePath(path string) {
	ib.path = path
	ib.index = 0
	ib.files = ib.listDir()
}

func (ib *ImageBrowser) listDir() []string {
	imgSrc, _ := os.ReadDir(ib.path)
	paths := make([]string, 0)
	for _, entry := range imgSrc {
		if entry.Type().IsRegular() && isImageFile(entry.Name()) {
			paths = append(paths, filepath.Join(ib.path, entry.Name()))
		}
	}
	return paths
}

func isImageFile(file string) bool {
	ext := strings.ToLower(filepath.Ext(file))
	return imageExts[ext]
}

func (ib *ImageBrowser) GetCurrent() *canvas.Image {
	if 0 > ib.index {
		ib.index = len(ib.files) - 1
	}
	if len(ib.files) <= ib.index {
		ib.index = 0
	}
	var imgSrc = ib.files[ib.index]
	var img = canvas.NewImageFromFile(imgSrc)
	img.FillMode = canvas.ImageFillContain

	return img
}

func (ib *ImageBrowser) GetNext() *canvas.Image {
	ib.index++
	return ib.GetCurrent()
}
