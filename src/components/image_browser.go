package components

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/google/uuid"
	"github.com/pauldin91/sego/src/common"
)

type ImageBrowser struct {
	widget.BaseWidget
	path    string
	index   int
	files   []string
	currImg *canvas.Image
	size    fyne.Size
	canvas  *DrawableCanvas
	pressed bool
}

func NewImageBrowser() *ImageBrowser {
	initPath, _ := os.Getwd()
	initPath = path.Join(initPath, common.DefaultResourceDir)
	ib := &ImageBrowser{
		path:   initPath,
		index:  0,
		size:   common.DefaultCanvasSize,
		canvas: NewDrawableCanvas(),
	}
	ib.currImg, _ = common.DefaultBlankImage(common.DefaultCanvasSize)
	ib.ExtendBaseWidget(ib)
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
	ib.canvas.clear()
	ib.path = path
	ib.index = 0
	ib.files = common.ListDir(ib.path)
	ib.Refresh()
}

func (ib *ImageBrowser) loadContent(path string) {
	ib.canvas.clear()
	ib.path = filepath.Dir(path)

	ib.index = 0
	ib.files = common.ListDir(ib.path)
	for i, f := range ib.files {
		if f == path {
			ib.index = i
			break
		}
	}
	ib.Refresh()
}

func (ib *ImageBrowser) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(container.NewStack(ib.currImg, ib.canvas.img))
}

func (dc *ImageBrowser) clear() {
	dc.pressed = false
	dc.canvas.clear()
}

func (ib *ImageBrowser) FocusLost()       {}
func (ib *ImageBrowser) FocusGained()     {}
func (ib *ImageBrowser) TypedRune(r rune) {}
func (ib *ImageBrowser) Focused() bool    { return true }
func (ib *ImageBrowser) TypedKey(event *fyne.KeyEvent) {
	switch event.Name {
	case fyne.KeyLeft:
		ib.getPrevious()
	case fyne.KeyRight:
		ib.getNext()
	case fyne.KeyS:
		ib.getNext()
		ib.Save()
	case fyne.KeyEqual:
		ib.canvas.brushSize++
	case fyne.KeyMinus:
		ib.canvas.brushSize--
	case fyne.KeyEscape:
		os.Exit(0)
	}
}
func (ib *ImageBrowser) Save() {
	var dir string = path.Join(ib.path, common.DefaultMaskDir)
	err := os.MkdirAll(dir, 0755)
	var filename string

	if err != nil || (ib.index >= len(ib.files) || ib.index < 0) {
		filename = path.Join(dir, "empty_"+uuid.New().String()+".png")
	} else {
		names := strings.Split(ib.files[ib.index], "/")
		filename = path.Join(dir, common.DefaultMaskPreffix+names[len(names)-1])
	}
	ib.canvas.SaveMask(filename)
	ib.clear()
}

func (d *ImageBrowser) DragEnd()                        { d.pressed = false }
func (d *ImageBrowser) MouseDown(e *desktop.MouseEvent) { d.pressed = true }
func (d *ImageBrowser) MouseUp(e *desktop.MouseEvent)   { d.pressed = false }
func (d *ImageBrowser) Dragged(e *fyne.DragEvent) {
	if !d.pressed {
		return
	}

	d.canvas.update(e.Position)
}
