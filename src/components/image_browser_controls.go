package components

import (
	"image/color"
	"math"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/storage"
	"github.com/pauldin91/sego/src/common"
)

func (ib *ImageBrowser) DragEnd() {}
func (ib *ImageBrowser) MouseDown(e *desktop.MouseEvent) {
	ib.pressed = true
	ib.update(e.Position)
	if canvas := fyne.CurrentApp().Driver().CanvasForObject(ib); canvas != nil {
		canvas.Focus(ib)
	}
}
func (ib *ImageBrowser) MouseUp(e *desktop.MouseEvent) { ib.pressed = false }
func (ib *ImageBrowser) FocusLost()                    {}
func (ib *ImageBrowser) FocusGained()                  {}
func (ib *ImageBrowser) TypedRune(r rune)              {}
func (ib *ImageBrowser) Focused() bool                 { return true }

func (ib *ImageBrowser) Dragged(e *fyne.DragEvent) {
	if !ib.pressed {
		return
	}
	ib.update(e.Position)
}

func (ib *ImageBrowser) TypedKey(event *fyne.KeyEvent) {

	switch event.Name {
	case fyne.KeyLeft:
		ib.getPrevious()
	case fyne.KeyRight:
		ib.getNext()
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
func (ib *ImageBrowser) ChooseColor(c color.Color) {
	r, g, b, _ := c.RGBA()
	ib.color = color.RGBA{
		R: uint8(r >> 8),
		G: uint8(g >> 8),
		B: uint8(b >> 8),
		A: 127}
	ib.transparrentColor = common.DefaultTransparrentColor
	ib.toogleBrush = true

}

func (ib *ImageBrowser) Clear() {
	ib.pressed = false
	ib.rgba = common.DefaultBlankImage(ib.BaseWidget.Size())
	ib.img.Image = ib.rgba
	fyne.Do(ib.img.Refresh)
}

func (ib *ImageBrowser) Save() {
	common.SaveMask(ib.rgba, ib.fb.GetMaskOrDefault())
	ib.Clear()
	ib.getNext()
}

func (ib *ImageBrowser) update(e fyne.Position) {
	ib.drawCircle(e)
	ib.img.Image = ib.rgba
	ib.img.Refresh()
}

func (ib *ImageBrowser) IncBrush() {
	ib.brushSize += common.DefaultBrushChange
}

func (ib *ImageBrowser) DecBrush() {
	if ib.brushSize >= 2*common.DefaultBrushChange {
		ib.brushSize -= common.DefaultBrushChange
	}
}

func (ib *ImageBrowser) Toggle() {
	temp := ib.color
	ib.color = ib.transparrentColor
	ib.transparrentColor = temp

	if !ib.toogleBrush {
		ib.toogleBrush = true
	} else {
		ib.toogleBrush = false
	}
}

func (ib *ImageBrowser) drawCircle(center fyne.Position) {
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

func (ib *ImageBrowser) onOpenFolderButtonClicked() {
	fd := dialog.NewFolderOpen(func(lu fyne.ListableURI, err error) {
		if err != nil || lu == nil {
			return
		}
		ib.UpdateImage(lu.Path())

	}, ib.parent)
	ib.setLocation(fd)
}

func (ib *ImageBrowser) onLoadFileButtonClicked() {
	fd := dialog.NewFileOpen(func(lu fyne.URIReadCloser, err error) {
		if err != nil || lu == nil {
			return
		}
		ib.LoadContent(lu.URI().Path())
	}, ib.parent)

	ib.setLocation(fd)
}

func (ib *ImageBrowser) onClearButtonClicked() {
	ib.Clear()
}

func (ib *ImageBrowser) setLocation(fd *dialog.FileDialog) {
	uri, err := storage.ListerForURI(storage.NewFileURI(ib.fb.GetPath()))
	if err == nil {
		fd.SetLocation(uri)
	}
	fd.Show()
}
