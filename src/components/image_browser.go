package components

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"path"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/pauldin91/sego/src/common"
)

type ImageBrowser struct {
	widget.BaseWidget
	brushSize float64
	img       *canvas.Image
	rgba      *image.RGBA
	pressed   bool
	path      string
	index     int
	files     []string
	currImg   *canvas.Image
	size      fyne.Size
}

func NewImageBrowser(size fyne.Size, path string) *ImageBrowser {
	ib := &ImageBrowser{
		path:      path,
		index:     0,
		size:      size,
		brushSize: common.BrushSize,
	}
	ib.files = common.ListDir(ib.path)
	ib.img, ib.rgba = common.DefaultBlankImage(size)
	ib.currImg, _ = common.DefaultBlankImage(size)
	ib.ExtendBaseWidget(ib)
	ib.Refresh()
	return ib
}

func (ib *ImageBrowser) Refresh() {
	if len(ib.files) == 0 {
		return
	}
	imgPath := ib.files[ib.index]
	ib.currImg.File = imgPath
	ib.currImg.Refresh()
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

func (ib *ImageBrowser) UpdatePath(path string) {
	ib.path = path
	ib.index = 0
	ib.files = common.ListDir(ib.path)
	ib.Refresh()
}

func (ib *ImageBrowser) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(container.NewStack(ib.currImg, ib.img))
}

func (d *ImageBrowser) Dragged(e *fyne.DragEvent) {
	if !d.pressed {
		return
	}

	d.drawCircle(e.Position)
	d.img.Image = d.rgba
	d.img.Refresh()
}
func (d *ImageBrowser) drawCircle(center fyne.Position) {
	for r := -d.brushSize; r < d.brushSize; r += 1.0 {
		for th := -math.Pi; th < math.Pi; th += math.Pi / 16 {
			x := r*math.Cos(th) + float64(center.X)
			y := r*math.Sin(th) + float64(center.Y)
			d.rgba.Set(int(x), int(y), color.RGBA{R: 182, G: 245, B: 0, A: 127})
		}
	}
}

func (d *ImageBrowser) DragEnd() {
	d.pressed = false
}

func (d *ImageBrowser) MouseDown(e *desktop.MouseEvent) {
	d.pressed = true
}

func (d *ImageBrowser) MouseUp(e *desktop.MouseEvent) {
	d.pressed = false
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
	case fyne.KeyS:
		ib.SaveMask()
	}
}

func (ib *ImageBrowser) SaveMask() {
	var dir string = path.Join(ib.path, "masks")
	err := os.Mkdir(dir, 0755)
	if err == nil || (ib.index >= len(ib.files) || ib.index < 0) {
		return
	}
	names := strings.Split(ib.files[ib.index], "/")
	filename := path.Join(dir, "mask_"+names[len(names)-1])

	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("error saving the image %s : %v\n", filename, err)
	}
	defer file.Close()

	png.Encode(file, ib.rgba)
}
