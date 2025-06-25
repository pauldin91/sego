package components

import (
	"os"
	"path"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/pauldin91/sego/src/common"
)

type ImageBrowser struct {
	widget.BaseWidget
	path          string
	index         int
	files         []string
	currImg       *canvas.Image
	size          fyne.Size
	filenames     chan string
	saveCompleted chan bool
}

func NewImageBrowser(fileChan chan string, saveCompleted chan bool) *ImageBrowser {
	initPath, _ := os.Getwd()
	initPath = path.Join(initPath, common.DefaultResourceDir)
	ib := &ImageBrowser{
		path:          initPath,
		index:         0,
		size:          common.Size,
		filenames:     fileChan,
		saveCompleted: saveCompleted,
	}
	ib.currImg, _ = common.DefaultBlankImage(common.Size)
	ib.ExtendBaseWidget(ib)
	return ib
}

func (ib *ImageBrowser) Refresh() {

	if len(ib.files) == 0 {
		return
	}
	imgPath := ib.files[ib.index]
	ib.currImg.File = imgPath
	ib.currImg.Refresh()
}

func (ib *ImageBrowser) getNext() {

	ib.filenames <- string(common.ImageChanged)

	if len(ib.files) == 0 {
		return
	}
	ib.index = (ib.index + 1) % len(ib.files)
	ib.Refresh()
}

func (ib *ImageBrowser) getPrevious() {

	ib.filenames <- string(common.ImageChanged)

	if len(ib.files) == 0 {
		return
	}
	ib.index = (ib.index - 1 + len(ib.files)) % len(ib.files)
	ib.Refresh()
}

func (ib *ImageBrowser) UpdatePath(path string) {
	ib.path = path
	ib.index = 0
	ib.files = common.ListDir(ib.path)
	ib.Refresh()
}

func (ib *ImageBrowser) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(ib.currImg)
}

func (ib *ImageBrowser) FocusLost()       {}
func (ib *ImageBrowser) FocusGained()     {}
func (ib *ImageBrowser) TypedRune(r rune) {}
func (ib *ImageBrowser) Focused() bool    { return true }

func (ib *ImageBrowser) TypedKey(event *fyne.KeyEvent) {

	switch event.Name {
	case fyne.KeyLeft:
		ib.getPrevious()
	case fyne.KeyRight:
		ib.getNext()
	case fyne.KeyS:
		ib.saveTrigger()
		<-ib.saveCompleted
		ib.getNext()
	}
}

func (ib *ImageBrowser) saveTrigger() {
	var dir string = path.Join(ib.path, common.DefaultMaskDir)
	err := os.MkdirAll(dir, 0755)
	if err != nil || (ib.index >= len(ib.files) || ib.index < 0) {
		return
	}
	names := strings.Split(ib.files[ib.index], "/")
	filename := path.Join(dir, common.DefaultMaskPreffix+names[len(names)-1])
	ib.filenames <- filename
}
