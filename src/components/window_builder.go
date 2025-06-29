package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/pauldin91/sego/src/common"
)

type WindowBuilder struct {
	widget.BaseWidget
	window   fyne.Window
	ib       *ImageBrowser
	canvas   *fyne.Container
	left     *fyne.Container
	combined *fyne.Container
}

func NewWindowBuilder(title string, a fyne.App) *WindowBuilder {

	result := &WindowBuilder{
		window:   a.NewWindow(title),
		canvas:   container.NewHBox(),
		left:     container.NewHBox(),
		combined: container.NewVBox(),
	}
	result.ib = NewImageBrowser(result.window)
	result.ExtendBaseWidget(result)
	return result
}
func (wb *WindowBuilder) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(wb.combined)
}

func (wb *WindowBuilder) WithBottomMenu() *WindowBuilder {
	res := NewBottomMenu(wb.ib, wb.window).
		WithButton(common.ColorBtn).
		WithButton(common.IncBtn).
		WithButton(common.DecBtn).
		WithButton(common.Toggle).
		Build()

	wb.left.Add(res)
	return wb
}

func (wb *WindowBuilder) WithDefaultCanvas() *WindowBuilder {
	wb.canvas.Add(wb.ib)
	return wb
}

func (wb *WindowBuilder) Refresh() {
	wb.combined = container.NewBorder(
		nil,
		container.NewCenter(wb.left),
		nil,
		nil,
		container.NewStack(wb.ib),
	)

	wb.window.SetContent(wb.combined)
	wb.window.Canvas().Focus(wb.ib)
	wb.window.Resize(wb.calcWindowSize())
	wb.window.SetTitle(wb.ib.title)
	wb.window.Content().Refresh()
}

func (wb *WindowBuilder) calcWindowSize() fyne.Size {

	var width = wb.window.Canvas().Size().Width + common.DefaultIconSize.Width + common.DefaultPaddingSize.Width
	var height = wb.window.Canvas().Size().Height + common.DefaultButtonSize.Height + common.DefaultPaddingSize.Height
	return fyne.NewSize(width, height)
}

func (wb *WindowBuilder) Build() fyne.Window {
	wb.Refresh()
	return wb.window
}

func (wb *WindowBuilder) WithMainMenu() *WindowBuilder {
	open := fyne.NewMenuItem("Open Folder", wb.ib.onOpenFolderButtonClicked)
	load := fyne.NewMenuItem("Load Image", wb.ib.onLoadFileButtonClicked)
	clear := fyne.NewMenuItem("Clear Mask", wb.ib.onClearButtonClicked)
	menu := fyne.NewMainMenu(fyne.NewMenu("Main Menu", open, load, clear))
	wb.window.SetMainMenu(menu)

	return wb
}
