package components

import (
	"os"
	"path"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/google/uuid"
	"github.com/pauldin91/sego/src/common"
)

func (mlst *ImageBrowser) DragEnd()                        { mlst.pressed = false }
func (mlst *ImageBrowser) MouseDown(e *desktop.MouseEvent) { mlst.pressed = true }
func (mlst *ImageBrowser) MouseUp(e *desktop.MouseEvent)   { mlst.pressed = false }
func (klst *ImageBrowser) FocusLost()                      {}
func (klst *ImageBrowser) FocusGained()                    {}
func (klst *ImageBrowser) TypedRune(r rune)                {}
func (klst *ImageBrowser) Focused() bool                   { return true }

func (mlst *ImageBrowser) Dragged(e *fyne.DragEvent) {
	if !mlst.pressed {
		return
	}
	mlst.canvas.update(e.Position)
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
		klst.Inc()
	case fyne.KeyMinus:
		klst.Dec()
	case fyne.KeyC:
		klst.Clear()
	case fyne.KeyEscape:
		os.Exit(0)
	}
}

func (ib *ImageBrowser) getNext() {
	if len(ib.files) == 0 {
		return
	}
	ib.Clear()
	ib.index = (ib.index + 1) % len(ib.files)
	ib.Refresh()
}

func (ib *ImageBrowser) getPrevious() {
	if len(ib.files) == 0 {
		return
	}
	ib.Clear()
	ib.index = (ib.index - 1 + len(ib.files)) % len(ib.files)
	ib.Refresh()
}

func (d *ImageBrowser) Inc()         { d.canvas.IncBrush() }
func (d *ImageBrowser) Dec()         { d.canvas.DecBrush() }
func (d *ImageBrowser) ToogleBrush() { d.canvas.Toggle() }
func (dc *ImageBrowser) Clear() {
	dc.pressed = false
	dc.canvas.clear()
}

func (ib *ImageBrowser) Save() {
	var dir string = path.Join(ib.path, common.DefaultMaskDir)
	err := os.MkdirAll(dir, 0755)
	var filename string

	if err != nil || (ib.index >= len(ib.files) || ib.index < 0) {
		filename = path.Join(dir, "empty_"+uuid.New().String()+".png")
	} else {
		filename = path.Join(dir, common.DefaultMaskPreffix+filepath.Base(ib.files[ib.index]))
	}
	ib.canvas.SaveMask(filename)
	ib.Clear()
	ib.getNext()
}
