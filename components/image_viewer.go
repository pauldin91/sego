package components

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"path/filepath"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/pauldin91/sego/utils"
)

type ImageViewer struct {
	widget.BaseWidget
	fb                *FileBrowser
	currImg           *canvas.Image
	title             string
	brushSize         float64
	img               *canvas.Image
	rgba              *image.RGBA
	toggleBrush       bool
	color             color.RGBA
	transparrentColor color.RGBA
	parent            fyne.Window
	cfg               utils.Config
}

func NewImageBrowser(cfg utils.Config, parent fyne.Window) *ImageViewer {

	ib := &ImageViewer{
		brushSize:         cfg.DefaultBrushSize,
		toggleBrush:       true,
		color:             utils.DefaultPaintColor,
		fb:                NewFileBrowser(cfg),
		cfg:               cfg,
		transparrentColor: utils.DefaultTransparrentColor,
		parent:            parent,
	}
	ib.rgba = utils.DefaultBlankImage(utils.DefaultCanvasSize)
	ib.currImg = canvas.NewImageFromImage(ib.rgba)
	ib.img = canvas.NewImageFromImage(ib.rgba)
	ib.currImg.FillMode = canvas.ImageFillContain
	ib.img.FillMode = canvas.ImageFillContain
	ib.title = "Canvas"
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

func (ib *ImageViewer) updateImage(path string) {
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

func (ib *ImageViewer) loadContent(selectedImgFile string) {
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

func (ib *ImageViewer) GetToggle() bool           { return ib.toggleBrush }
func (ib *ImageViewer) GetTitle() string          { return ib.title }
func (ib *ImageViewer) DragEnd()                  {}
func (ib *ImageViewer) Tapped(e *fyne.PointEvent) { ib.update(e.Position) }
func (ib *ImageViewer) Dragged(e *fyne.DragEvent) { ib.update(e.Position) }
func (ib *ImageViewer) FocusLost()                {}
func (ib *ImageViewer) FocusGained()              {}
func (ib *ImageViewer) TypedRune(r rune)          {}
func (ib *ImageViewer) Focused() bool             { return true }
func (ib *ImageViewer) TypedKey(event *fyne.KeyEvent) {

	switch event.Name {
	case fyne.KeyLeft:
		ib.GetPrevious()
	case fyne.KeyRight:
		ib.GetNext()
	case fyne.KeyS:
		ib.Save()
	case fyne.KeyEqual:
		ib.IncBrush()
	case fyne.KeyMinus:
		ib.DecBrush()
	case fyne.KeyC:
		ib.Clear()
	case fyne.KeyEscape:
		os.Exit(0)
	}
}

func (ib *ImageViewer) GetNext() {
	ib.fb.Next()
	ib.Clear()
	ib.Refresh()
}

func (ib *ImageViewer) GetPrevious() {
	ib.fb.Previous()
	ib.Clear()
	ib.Refresh()
}
func (ib *ImageViewer) ChooseColor(c color.Color) {
	r, g, b, _ := c.RGBA()
	ib.color = color.RGBA{
		R: uint8(r >> 8),
		G: uint8(g >> 8),
		B: uint8(b >> 8),
		A: 127}
	ib.transparrentColor = utils.DefaultTransparrentColor
	ib.toggleBrush = true

}

func (ib *ImageViewer) Clear() {
	ib.rgba = utils.DefaultBlankImage(ib.BaseWidget.Size())
	ib.img.Image = ib.rgba
	ib.img.Refresh()
}

func (ib *ImageViewer) Save() {
	utils.SaveMask(ib.rgba, ib.fb.GetMaskOrDefault())
	ib.GetNext()
}

func (ib *ImageViewer) update(e fyne.Position) {
	ib.drawCircle(e)
	ib.img.Image = ib.rgba
	ib.img.Refresh()
}

func (ib *ImageViewer) IncBrush() {
	if ib.brushSize < ib.cfg.DefaultMaxBrushSize {
		ib.brushSize += ib.cfg.DefaultBrushChange
	}
}

func (ib *ImageViewer) DecBrush() {
	if ib.brushSize >= 2*ib.cfg.DefaultBrushChange {
		ib.brushSize -= ib.cfg.DefaultBrushChange
	}
}

func (ib *ImageViewer) Toggle() {
	temp := ib.color
	ib.color = ib.transparrentColor
	ib.transparrentColor = temp

	if !ib.toggleBrush {
		ib.toggleBrush = true
	} else {
		ib.toggleBrush = false
	}
}

func (ib *ImageViewer) drawCircle(center fyne.Position) {
	for r := -ib.brushSize; r < ib.brushSize; r += 1.0 {
		bounds := ib.rgba.Bounds()
		for th := -math.Pi; th < math.Pi; th += math.Pi / 16 {
			x := r*math.Cos(th) + float64(center.X)
			y := r*math.Sin(th) + float64(center.Y)
			if int(x) >= bounds.Min.X && int(x) < bounds.Max.X && int(y) >= bounds.Min.Y && int(y) < bounds.Max.Y {
				ib.rgba.Set(int(x), int(y), ib.color)
			}
		}
	}
}

func (ib *ImageViewer) ChangeBrushSize(s string) {
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		ib.brushSize = f
	}
}

func (ib *ImageViewer) OnOpenFolderButtonClicked() {
	fd := dialog.NewFolderOpen(func(lu fyne.ListableURI, err error) {
		if err != nil || lu == nil {
			return
		}
		ib.updateImage(lu.Path())

	}, ib.parent)
	ib.setLocation(fd)
}

func (ib *ImageViewer) OnLoadFileButtonClicked() {
	fd := dialog.NewFileOpen(func(lu fyne.URIReadCloser, err error) {
		if err != nil || lu == nil {
			return
		}
		ib.loadContent(lu.URI().Path())
	}, ib.parent)

	ib.setLocation(fd)
}

func (ib *ImageViewer) OnClearButtonClicked() {
	ib.Clear()
}

func (ib *ImageViewer) setLocation(fd *dialog.FileDialog) {
	uri, err := storage.ListerForURI(storage.NewFileURI(ib.fb.GetPath()))
	if err == nil {
		fd.SetLocation(uri)
	}
	fd.Show()
}
