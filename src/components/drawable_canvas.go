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
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/pauldin91/sego/src/common"
)

type DrawableCanvas struct {
	widget.BaseWidget
	brushSize     float64
	img           *canvas.Image
	rgba          *image.RGBA
	pressed       bool
	size          fyne.Size
	filenames     chan string
	saveCompleted chan bool
}

func NewDrawableCanvas(fileChan chan string, saveCompleted chan bool) *DrawableCanvas {
	dc := &DrawableCanvas{
		brushSize:     common.BrushSize,
		size:          common.Size,
		filenames:     fileChan,
		saveCompleted: saveCompleted,
	}
	dc.img, dc.rgba = common.DefaultBlankImage(common.Size)
	go func() {
		dc.handleFileSaving()
	}()
	return dc
}

func (ib *DrawableCanvas) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(ib.img)
}

func (dc *DrawableCanvas) clear() {
	dc.pressed = false
	_, dc.rgba = common.DefaultBlankImage(dc.size)
	fn := func() {
		dc.img.Image = dc.rgba
		dc.img.Refresh()
	}
	fyne.Do(fn)
}

func (d *DrawableCanvas) Dragged(e *fyne.DragEvent) {
	if !d.pressed {
		return
	}

	d.drawCircle(e.Position)
	d.img.Image = d.rgba
	d.img.Refresh()
}
func (d *DrawableCanvas) drawCircle(center fyne.Position) {
	for r := -d.brushSize; r < d.brushSize; r += 1.0 {
		for th := -math.Pi; th < math.Pi; th += math.Pi / 16 {
			x := r*math.Cos(th) + float64(center.X)
			y := r*math.Sin(th) + float64(center.Y)
			d.rgba.Set(int(x), int(y), color.RGBA{R: 182, G: 245, B: 0, A: 127})
		}
	}
}

func (d *DrawableCanvas) DragEnd()                        { d.pressed = false }
func (d *DrawableCanvas) MouseDown(e *desktop.MouseEvent) { d.pressed = true }
func (d *DrawableCanvas) MouseUp(e *desktop.MouseEvent)   { d.pressed = false }

func (dc *DrawableCanvas) handleFileSaving() {
	for {
		select {
		case filename := <-dc.filenames:
			if filename != string(common.ImageChanged) {
				dc.SaveMask(filename)
				dc.saveCompleted <- true
			}
			dc.clear()
		}
	}
}

func (ib *DrawableCanvas) SaveMask(filename string) {

	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("error saving the image %s : %v\n", filename, err)
	}
	defer file.Close()

	png.Encode(file, ib.rgba)
}
