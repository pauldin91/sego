package browser

import (
	"os"
	"path"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/pauldin91/sego/src/common"
)

type FileBrowser struct {
	path  string
	index int
	files []string
	cfg   common.Config
}

func NewFileBrowser(cfg common.Config) *FileBrowser {
	initPath, _ := os.Getwd()
	initPath = path.Join(initPath, cfg.DefaultResourceDir)
	res := &FileBrowser{
		path:  initPath,
		index: 0,
		files: make([]string, 0),
		cfg:   cfg,
	}
	return res
}

func (fb *FileBrowser) GetMask(selectedImgFile string) string {
	name := fb.cfg.DefaultMaskPreffix + filepath.Base(selectedImgFile)
	return path.Join(fb.path, fb.cfg.DefaultMaskDir, name)
}

func (fb *FileBrowser) UpdatePath(path string) {
	fb.path = path
	fb.index = 0
	fb.files = common.ListDir(fb.path)
}

func (fb *FileBrowser) GetPath() string { return fb.path }

func (fb *FileBrowser) GetFilename() string {
	if len(fb.files) == 0 {
		return ""
	}
	return fb.files[fb.index]
}

func (fb *FileBrowser) SetIndexForFilename(selectedImgFile string) {
	fb.path = filepath.Dir(selectedImgFile)
	fb.index = 0
	fb.files = common.ListDir(fb.path)
	for i, f := range fb.files {
		if f == selectedImgFile {
			fb.index = i
			break
		}
	}
}

func (ib *FileBrowser) Next() {
	if len(ib.files) == 0 {
		return
	}
	ib.index = (ib.index + 1) % len(ib.files)
}

func (ib *FileBrowser) Previous() {
	if len(ib.files) == 0 {
		return
	}
	ib.index = (ib.index - 1 + len(ib.files)) % len(ib.files)
}

func (ib *FileBrowser) GetMaskOrDefault() string {
	var dir string = path.Join(ib.path, ib.cfg.DefaultMaskDir)
	err := os.MkdirAll(dir, 0755)
	var filename string

	if err != nil || (ib.index >= len(ib.files) || ib.index < 0) {
		filename = path.Join(dir, "empty_"+uuid.New().String()+".png")
	} else {
		filename = path.Join(dir, ib.cfg.DefaultMaskPreffix+filepath.Base(ib.files[ib.index]))
	}

	return filename
}
