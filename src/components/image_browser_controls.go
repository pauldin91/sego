package components

import (
	"image/color"
	"math"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/pauldin91/sego/src/common"
)

func (mlst *ImageBrowser) DragEnd() {}
func (mlst *ImageBrowser) MouseDown(e *desktop.MouseEvent) {
	mlst.pressed = true
	mlst.update(e.Position)
	if canvas := fyne.CurrentApp().Driver().CanvasForObject(mlst); canvas != nil {
		canvas.Focus(mlst)
	}
}
func (mlst *ImageBrowser) MouseUp(e *desktop.MouseEvent) { mlst.pressed = false }
func (klst *ImageBrowser) FocusLost()                    {}
func (klst *ImageBrowser) FocusGained()                  {}
func (klst *ImageBrowser) TypedRune(r rune)              {}
func (klst *ImageBrowser) Focused() bool                 { return true }

func (mlst *ImageBrowser) Dragged(e *fyne.DragEvent) {
	if !mlst.pressed {
		return
	}
	mlst.update(e.Position)
}

func (klst *ImageBrowser) TypedKey(event *fyne.KeyEvent) {

	switch event.Name {
	case fyne.KeyLeft:
		klst.getPrevious()
	case fyne.KeyRight:
		klst.getNext()
	case fyne.KeyS:
		klst.Save()
	case fyne.KeyEqual:
		klst.IncBrush()
	case fyne.KeyMinus:
		klst.DecBrush()
	case fyne.KeyC:
		klst.Clear()
	case fyne.KeyEscape:
		os.Exit(0)
	}
}

func (ib *ImageBrowser) getNext() {
	ib.fb.Next()
	ib.Clear()
	ib.Refresh()
}

func (ib *ImageBrowser) getPrevious() {
	ib.fb.Previous()
	ib.Clear()
	ib.Refresh()
}
func (d *ImageBrowser) ChooseColor(c color.Color) {
	r, g, b, _ := c.RGBA()
	d.color = color.RGBA{
		R: uint8(r >> 8),
		G: uint8(g >> 8),
		B: uint8(b >> 8),
		A: 127}
}

func (dc *ImageBrowser) Clear() {
	dc.pressed = false
	dc.rgba = common.DefaultBlankImage(dc.BaseWidget.Size())
	dc.img.Image = dc.rgba
	fyne.Do(dc.img.Refresh)
}

func (ib *ImageBrowser) Save() {
	common.SaveMask(ib.rgba, ib.fb.GetMaskOrDefault())
	ib.Clear()
	ib.getNext()
}

func (dc *ImageBrowser) update(e fyne.Position) {
	dc.drawCircle(e)
	dc.img.Image = dc.rgba
	dc.img.Refresh()
}

func (dc *ImageBrowser) IncBrush() {
	dc.brushSize += common.DefaultBrushChange
}

func (dc *ImageBrowser) DecBrush() {
	if dc.brushSize >= 2*common.DefaultBrushChange {
		dc.brushSize -= common.DefaultBrushChange
	}
}

func (dc *ImageBrowser) Toggle() {
	if !dc.toogleBrush {
		dc.color = common.DefaultPaintColor
		dc.toogleBrush = true
	} else {
		dc.color = color.RGBA{R: 0, G: 0, B: 0, A: 0}
		dc.toogleBrush = false
	}
}

func (d *ImageBrowser) drawCircle(center fyne.Position) {
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
