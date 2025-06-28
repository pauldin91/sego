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
	bottom   *fyne.Container
	combined *fyne.Container
}

func NewWindowBuilder(size fyne.Size, title string, a fyne.App) *WindowBuilder {

	result := &WindowBuilder{
		window:   a.NewWindow(title),
		canvas:   container.NewHBox(),
		left:     container.NewVBox(),
		bottom:   container.NewHBox(),
		combined: container.NewVBox(),
		ib:       NewImageBrowser(),
	}
	result.window.Resize(result.calcWindowSize())
	result.ExtendBaseWidget(result)
	return result
}
func (ib *WindowBuilder) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(container.NewCenter(ib.combined))
}

func (wb *WindowBuilder) WithSidebarMenu() *WindowBuilder {
	res := NewSidebarMenu(wb.ib).
		WithButton(common.IncBtn).
		WithButton(common.DecBtn).
		WithButton(common.Toggle).
		Build()

	wb.left.Add(res)
	return wb
}

func (wb *WindowBuilder) WithBottomMenu() *WindowBuilder {
	res := NewBottomMenu(wb.ib, wb.window).
		WithButton(common.OpenBtn).
		WithButton(common.LoadBtn).
		WithButton(common.ClearBtn).
		Build()

	wb.bottom.Add(res)
	return wb
}

func (wb *WindowBuilder) WithDefaultCanvas() *WindowBuilder {
	wb.canvas.Add(wb.ib)
	return wb
}

func (wb *WindowBuilder) Refresh() {
	wb.combined = container.NewBorder(
		nil,
		container.NewCenter(wb.bottom),
		container.NewCenter(wb.left),
		nil,
		container.NewStack(wb.ib),
	)

	wb.window.SetContent(wb.combined)
	wb.window.Canvas().Focus(wb.ib)
	wb.Resize(wb.calcWindowSize())
	wb.window.SetTitle(wb.ib.title)
	wb.window.Content().Refresh()
}

func (wb *WindowBuilder) calcWindowSize() fyne.Size {

	var width = wb.window.Canvas().Size().Width + common.DefaultIconSize.Width + common.DefaultPaddingSize.Width
	var height = wb.window.Canvas().Size().Height + common.DefaultButtonSize.Height + common.DefaultPaddingSize.Height
	wb.ib.Resize(wb.window.Canvas().Size())
	return fyne.NewSize(width, height)
}

func (wb *WindowBuilder) Build() fyne.Window {
	wb.Refresh()
	return wb.window
}
