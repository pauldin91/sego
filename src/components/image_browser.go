package components

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/pauldin91/sego/src/common"
)

type ImageBrowser struct {
	widget.BaseWidget
	fb          *FileBrowser
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

	ib := &ImageBrowser{
		brushSize:   common.DefaultBrushSize,
		toogleBrush: true,
		color:       common.DefaultPaintColor,
		fb:          NewFileBrowser(),
	}
	ib.rgba = common.DefaultBlankImage(common.DefaultCanvasSize)
	ib.currImg = canvas.NewImageFromImage(ib.rgba)
	ib.img = canvas.NewImageFromImage(ib.rgba)
	ib.currImg.FillMode = canvas.ImageFillContain
	ib.img.FillMode = canvas.ImageFillContain
	ib.title = "Canvas"

	ib.currImg.SetMinSize(common.DefaultCanvasSize)
	ib.img.SetMinSize(common.DefaultCanvasSize)
	ib.ExtendBaseWidget(ib)
	return ib
}

func (ib *ImageBrowser) Refresh() {

	ib.currImg.File = ib.fb.GetFilename()
	if ib.currImg.File == "" {
		return
	}
	ib.title = filepath.Base(ib.currImg.File)
	ib.loadMask(ib.currImg.File)
	ib.currImg.Refresh()
}

func (ib *ImageBrowser) UpdatePath(path string) {
	ib.Clear()
	ib.fb.UpdatePath(path)
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
	ib.fb.FindFilename(selectedImgFile)
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
	mask := ib.fb.GetMaskOrDefault(selectedImgFile)
	if file, err := os.Open(mask); err == nil {
		defer file.Close()
		if img, err := png.Decode(file); err == nil {

			ib.rgba = common.ScaleImage(img, ib.Size())
			ib.img.Image = ib.rgba
			ib.img.Refresh()
		}

	}
}
