package viewer

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
	"github.com/pauldin91/sego/src/components/browser"
	"github.com/pauldin91/sego/src/components/utils"
)

type ImageViewer struct {
	widget.BaseWidget
	fb                *browser.FileBrowser
	currImg           *canvas.Image
	pressed           bool
	title             string
	brushSize         float64
	img               *canvas.Image
	rgba              *image.RGBA
	toggleBrush       bool
	color             color.RGBA
	transparrentColor color.RGBA
	parent            fyne.Window
}

func NewImageBrowser(parent fyne.Window) *ImageViewer {

	ib := &ImageViewer{
		brushSize:         common.DefaultBrushSize,
		toggleBrush:       true,
		color:             common.DefaultPaintColor,
		fb:                browser.NewFileBrowser(),
		transparrentColor: common.DefaultTransparrentColor,
		parent:            parent,
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

func (ib *ImageViewer) Refresh() {

	ib.currImg.File = ib.fb.GetFilename()
	if ib.currImg.File == "" {
		return
	}
	ib.title = filepath.Base(ib.currImg.File)
	ib.loadMask(ib.currImg.File)
	ib.currImg.Refresh()
}

func (ib *ImageViewer) UpdateImage(path string) {
	ib.Clear()
	ib.fb.UpdatePath(path)
	ib.Refresh()
}

func (ib *ImageViewer) Resize(size fyne.Size) {
	ib.BaseWidget.Resize(size)
	ib.currImg.Resize(size)
	ib.img.Resize(size)
	ib.rgba = utils.ScaleImage(ib.rgba, fyne.NewSize(size.Width, size.Height))
	ib.img.Image = ib.rgba
}

func (ib *ImageViewer) LoadContent(selectedImgFile string) {
	ib.fb.SetIndexForFilename(selectedImgFile)
	ib.Refresh()
}

func (ib *ImageViewer) CreateRenderer() fyne.WidgetRenderer {

	stack := container.NewStack(
		ib.currImg,
		ib.img,
	)
	return widget.NewSimpleRenderer(stack)
}

func (ib *ImageViewer) loadMask(selectedImgFile string) {
	mask := ib.fb.GetMask(selectedImgFile)
	if file, err := os.Open(mask); err == nil {
		defer file.Close()
		if img, err := png.Decode(file); err == nil {

			ib.rgba = utils.ScaleImage(img, ib.Size())
			ib.img.Image = ib.rgba
			ib.img.Refresh()
		}

	}
}

func (ib *ImageViewer) GetToggle() bool  { return ib.toggleBrush }
func (ib *ImageViewer) GetTitle() string { return ib.title }
