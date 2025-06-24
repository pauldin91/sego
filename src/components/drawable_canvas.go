package components

import (
	"image"
	"image/color"
	"image/draw"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type DrawCanvas struct {
	widget.BaseWidget
	img     *canvas.Image
	rgba    *image.RGBA
	pressed bool
}

func NewDrawCanvas(size fyne.Size) *DrawCanvas {
	dc := &DrawCanvas{
		rgba: image.NewRGBA(image.Rect(0, 0, int(size.Width), int(size.Height))),
	}
	draw.Draw(dc.rgba, dc.rgba.Bounds(), &image.Uniform{color.Transparent}, image.Point{}, draw.Src)

	dc.img = canvas.NewImageFromImage(dc.rgba)
	dc.img.FillMode = canvas.ImageFillContain
	dc.img.SetMinSize(size)
	dc.ExtendBaseWidget(dc)
	return dc
}

func (d *DrawCanvas) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(d.img)
}

func (d *DrawCanvas) Dragged(e *fyne.DragEvent) {
	if !d.pressed {
		return
	}
	x, y := int(e.Position.X), int(e.Position.Y)
	d.rgba.Set(x, y, color.White)
	d.img.Image = d.rgba
	d.img.Refresh()
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
