package components

import (
	"image"
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/pauldin91/sego/src/common"
)

type DrawCanvas struct {
	widget.BaseWidget
	img     *canvas.Image
	rgba    *image.RGBA
	pressed bool
}

func NewDrawCanvas(size fyne.Size) *DrawCanvas {
	drawableCanvas := &DrawCanvas{}

	img, rgba := common.DefaultBlankImage(size)
	drawableCanvas.img = img
	drawableCanvas.rgba = rgba

	drawableCanvas.ExtendBaseWidget(drawableCanvas)
	return drawableCanvas
}

func (d *DrawCanvas) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(d.img)
}

func (d *DrawCanvas) Dragged(e *fyne.DragEvent) {
	if !d.pressed {
		return
	}
	circle := canvas.NewCircle(color.White)
	circle.Resize(fyne.NewSize(5, 5))
	d.drawCircle(e.Position)
	d.img.Image = d.rgba
	d.img.Refresh()
}
func (d *DrawCanvas) drawCircle(center fyne.Position) {
	for r := -5.0; r < 5.0; r += 1 {
		for th := -math.Pi; th < math.Pi; th += math.Pi / 16 {
			x := r*math.Cos(th) + float64(center.X)
			y := r*math.Sin(th) + float64(center.Y)
			d.rgba.Set(int(x), int(y), color.White)
		}
	}
}

func (d *DrawCanvas) DragEnd() {
	d.pressed = false
}

func (d *DrawCanvas) MouseDown(e *desktop.MouseEvent) {
	d.pressed = true
}

func (d *DrawCanvas) MouseUp(e *desktop.MouseEvent) {
	d.pressed = false
}
