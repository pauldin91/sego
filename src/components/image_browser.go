package components

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"path"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/pauldin91/sego/src/common"
)

type ImageBrowser struct {
	widget.BaseWidget
	path        string
	index       int
	files       []string
	currImg     *canvas.Image
	pressed     bool
	title       string
	brushSize   float64
	img         *canvas.Image
	rgba        *image.RGBA
	toogleBrush bool
	color       color.RGBA
}

func NewImageBrowser() *ImageBrowser {
	initPath, _ := os.Getwd()
	initPath = path.Join(initPath, common.DefaultResourceDir)
	ib := &ImageBrowser{
		path:        initPath,
		index:       0,
		brushSize:   common.DefaultBrushSize,
		toogleBrush: true,
		color:       common.DefaultPaintColor,
	}
	ib.currImg, ib.rgba = common.DefaultBlankImage(common.DefaultCanvasSize)
	ib.currImg.FillMode = canvas.ImageFillContain
	ib.title = "Canvas"
	ib.currImg.SetMinSize(common.DefaultCanvasSize)

	ib.img = canvas.NewImageFromImage(ib.rgba)
	ib.img.FillMode = canvas.ImageFillContain
	ib.img.SetMinSize(common.DefaultCanvasSize)
	ib.ExtendBaseWidget(ib)
	return ib
}

func (ib *ImageBrowser) Refresh() {

	if len(ib.files) == 0 {
		return
	}
	imgPath := ib.files[ib.index]
	ib.currImg.File = imgPath
	ib.title = filepath.Base(imgPath)
	ib.loadMask(ib.files[ib.index])
	ib.currImg.Refresh()
}

func (ib *ImageBrowser) UpdatePath(path string) {
	ib.Clear()
	ib.path = path
	ib.index = 0
	ib.files = common.ListDir(ib.path)
	ib.Refresh()
}

func (ib *ImageBrowser) Resize(size fyne.Size) {
	ib.BaseWidget.Resize(size)
	ib.currImg.Resize(size)
	ib.img.Resize(size)
	ib.rgba = common.ScaleImage(ib.rgba, fyne.NewSize(size.Width, size.Height))
	ib.img.Image = ib.rgba
}

func (ib *ImageBrowser) loadContent(selectedImgFile string) {
	ib.path = filepath.Dir(selectedImgFile)

	ib.index = 0
	ib.files = common.ListDir(ib.path)
	for i, f := range ib.files {
		if f == selectedImgFile {
			ib.index = i
			break
		}
	}

	ib.Refresh()
}

func (ib *ImageBrowser) CreateRenderer() fyne.WidgetRenderer {

	stack := container.NewStack(
		ib.currImg,
		ib.img,
	)
	return widget.NewSimpleRenderer(stack)
}

func (ib *ImageBrowser) loadMask(selectedImgFile string) {
	name := common.DefaultMaskPreffix + filepath.Base(selectedImgFile)
	mask := path.Join(ib.path, common.DefaultMaskDir, name)
	if file, err := os.Open(mask); err == nil {
		defer file.Close()
		if img, err := png.Decode(file); err == nil {

			ib.rgba = common.ScaleImage(img, ib.Size())
			ib.img.Image = ib.rgba
			ib.img.Refresh()
		}

	}
}
