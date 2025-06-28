package components

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/pauldin91/sego/src/common"
)

type DrawableCanvas struct {
	widget.BaseWidget
	brushSize   float64
	size        fyne.Size
	img         *canvas.Image
	rgba        *image.RGBA
	toogleBrush bool
	color       color.RGBA
}

func NewDrawableCanvas() *DrawableCanvas {
	dc := &DrawableCanvas{
		brushSize:   common.DefaultBrushSize,
		size:        common.DefaultCanvasSize,
		toogleBrush: true,
		color:       common.DefaultPaintColor,
	}
	dc.img, dc.rgba = common.DefaultBlankImage(common.DefaultCanvasSize)
	dc.ExtendBaseWidget(dc)
	return dc
}

func (ib *DrawableCanvas) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(ib.img)
}

func (ib *DrawableCanvas) SaveMask(filename string) {

	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("error saving the image %s : %v\n", filename, err)
		return
	}
	defer file.Close()

	png.Encode(file, ib.rgba)
}

func (d *DrawableCanvas) drawCircle(center fyne.Position) {
	for r := -d.brushSize; r < d.brushSize; r += 1.0 {
		bounds := d.rgba.Bounds()
		for th := -math.Pi; th < math.Pi; th += math.Pi / 16 {
			x := r*math.Cos(th) + float64(center.X)
			y := r*math.Sin(th) + float64(center.Y)
			if int(x) >= bounds.Min.X && int(x) < bounds.Max.X && int(y) >= bounds.Min.Y && int(y) < bounds.Max.Y {
				d.rgba.Set(int(x), int(y), d.color)
			}
		}
	}
}

func (dc *DrawableCanvas) reset() {
	dc.img.File = ""
	dc.img.Image = dc.rgba
	dc.Refresh()
}

func (dc *DrawableCanvas) clear() {
	_, dc.rgba = common.DefaultBlankImage(dc.size)
	fyne.Do(dc.reset)
}

func (dc *DrawableCanvas) update(e fyne.Position) {
	dc.drawCircle(e)
	dc.img.Image = dc.rgba
	dc.img.Refresh()
}

func (dc *DrawableCanvas) IncBrush() {
	dc.brushSize += common.DefaultBrushChange
}

func (dc *DrawableCanvas) DecBrush() {
	if dc.brushSize >= 2*common.DefaultBrushChange {
		dc.brushSize -= common.DefaultBrushChange
	}
}

func (dc *DrawableCanvas) Toggle() {
	if !dc.toogleBrush {
		dc.color = common.DefaultPaintColor
		dc.toogleBrush = true
	} else {
		dc.color = color.RGBA{R: 0, G: 0, B: 0, A: 0}
		dc.toogleBrush = false
	}
}
