package viewer

import (
	"image/color"
	"math"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"github.com/pauldin91/sego/src/common"
	"github.com/pauldin91/sego/src/components/utils"
)

func (ib *ImageViewer) DragEnd() {}
func (ib *ImageViewer) Tapped(e *fyne.PointEvent) {
	ib.update(e.Position)
}
func (ib *ImageViewer) Dragged(e *fyne.DragEvent) {
	ib.update(e.Position)
}

func (ib *ImageViewer) FocusLost()       {}
func (ib *ImageViewer) FocusGained()     {}
func (ib *ImageViewer) TypedRune(r rune) {}
func (ib *ImageViewer) Focused() bool    { return true }
func (ib *ImageViewer) TypedKey(event *fyne.KeyEvent) {

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

func (ib *ImageViewer) getNext() {
	ib.fb.Next()
	ib.Clear()
	ib.Refresh()
}

func (ib *ImageViewer) getPrevious() {
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
	ib.transparrentColor = common.DefaultTransparrentColor
	ib.toggleBrush = true

}

func (ib *ImageViewer) Clear() {
	ib.pressed = false
	ib.rgba = common.DefaultBlankImage(ib.BaseWidget.Size())
	ib.img.Image = ib.rgba
	ib.img.Refresh()
}

func (ib *ImageViewer) Save() {
	utils.SaveMask(ib.rgba, ib.fb.GetMaskOrDefault())
	ib.Clear()
	ib.getNext()
}

func (ib *ImageViewer) update(e fyne.Position) {
	ib.drawCircle(e)
	ib.img.Image = ib.rgba
	ib.img.Refresh()
}

func (ib *ImageViewer) IncBrush() {
	ib.brushSize += common.DefaultBrushChange
}

func (ib *ImageViewer) DecBrush() {
	if ib.brushSize >= 2*common.DefaultBrushChange {
		ib.brushSize -= common.DefaultBrushChange
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

func (ib *ImageViewer) OnOpenFolderButtonClicked() {
	fd := dialog.NewFolderOpen(func(lu fyne.ListableURI, err error) {
		if err != nil || lu == nil {
			return
		}
		ib.UpdateImage(lu.Path())

	}, ib.parent)
	ib.setLocation(fd)
}

func (ib *ImageViewer) OnLoadFileButtonClicked() {
	fd := dialog.NewFileOpen(func(lu fyne.URIReadCloser, err error) {
		if err != nil || lu == nil {
			return
		}
		ib.LoadContent(lu.URI().Path())
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
