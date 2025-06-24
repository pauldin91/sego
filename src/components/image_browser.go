package components

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/pauldin91/sego/src/common"
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
	widget.BaseWidget

	path    string
	index   int
	files   []string
	currImg *canvas.Image
	size    fyne.Size
}

func NewImageBrowser(size fyne.Size, path string) *ImageBrowser {
	ib := &ImageBrowser{
		path:  path,
		index: 0,
		size:  size,
	}
	ib.files = ib.listDir()
	ib.currImg, _ = common.DefaultBlankImage(size)
	ib.ExtendBaseWidget(ib)
	ib.loadCurrentImage()
	return ib
}

func (ib *ImageBrowser) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(ib.currImg)
}

func (ib *ImageBrowser) listDir() []string {
	entries, err := os.ReadDir(ib.path)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil
	}
	paths := []string{}
	for _, entry := range entries {
		if entry.Type().IsRegular() && isImage(entry.Name()) {
			paths = append(paths, filepath.Join(ib.path, entry.Name()))
		}
	}
	return paths
}

func isImage(file string) bool {
	ext := strings.ToLower(filepath.Ext(file))
	return imageExts[ext]
}

func (ib *ImageBrowser) loadCurrentImage() {
	if len(ib.files) == 0 {
		return
	}
	imgPath := ib.files[ib.index]
	ib.currImg.File = imgPath
	ib.currImg.Refresh()
}

func (ib *ImageBrowser) Refresh() {
	ib.loadCurrentImage()
}

func (ib *ImageBrowser) getNext() {
	if len(ib.files) == 0 {
		return
	}
	ib.index = (ib.index + 1) % len(ib.files)
	ib.Refresh()
}

func (ib *ImageBrowser) getPrevious() {
	if len(ib.files) == 0 {
		return
	}
	ib.index = (ib.index - 1 + len(ib.files)) % len(ib.files)
	ib.Refresh()
}

func (ib *ImageBrowser) DirCount() int {
	return len(ib.files)
}

func (ib *ImageBrowser) UpdatePath(path string) {
	ib.path = path
	ib.index = 0
	ib.files = ib.listDir()
	ib.Refresh()
}

func (ib *ImageBrowser) FocusLost()       {}
func (ib *ImageBrowser) FocusGained()     {}
func (ib *ImageBrowser) TypedRune(r rune) {}
func (ib *ImageBrowser) Focused() bool    { return true }

func (ib *ImageBrowser) TypedKey(event *fyne.KeyEvent) {
	fmt.Printf("Key pressed: %s (%v)\n", event.Name, event.Physical)

	switch event.Name {
	case fyne.KeyLeft:
		ib.getPrevious()
	case fyne.KeyRight:
		ib.getNext()
	}
}
