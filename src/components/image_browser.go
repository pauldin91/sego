package components

import (
	"image"
	"image/draw"
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
	path    string
	index   int
	files   []string
	currImg *canvas.Image
	size    fyne.Size
	canvas  *DrawableCanvas
	pressed bool
	title   string
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
	ib.title = "Canvas"
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
	return widget.NewSimpleRenderer(container.NewStack(ib.currImg, ib.canvas.img))
}

func (ib *ImageBrowser) loadMask(selectedImgFile string) {
	name := common.DefaultMaskPreffix + filepath.Base(selectedImgFile)
	mask := path.Join(ib.path, common.DefaultMaskDir, name)
	if file, err := os.Open(mask); err == nil {
		defer file.Close()
		if img, err := png.Decode(file); err == nil {
			bounds := img.Bounds()
			rgba := image.NewRGBA(bounds)
			draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)

			ib.canvas.rgba = rgba
			ib.canvas.img.Image = rgba
			ib.canvas.Refresh()
		}

	}
}
