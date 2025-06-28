package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/pauldin91/sego/src/common"
)

type WindowBuilder struct {
	widget.BaseWidget
	window     fyne.Window
	ib         *ImageBrowser
	canvas     *fyne.Container
	left       *fyne.Container
	bottom     *fyne.Container
	combined   *fyne.Container
	canvasSize fyne.Size
}

func NewWindowBuilder(size fyne.Size, title string, a fyne.App) *WindowBuilder {

	result := &WindowBuilder{
		window:     a.NewWindow(title),
		canvas:     container.NewHBox(),
		left:       container.NewVBox(),
		bottom:     container.NewHBox(),
		combined:   container.NewVBox(),
		ib:         NewImageBrowser(),
		canvasSize: size,
	}
	result.window.Resize(result.canvasSize)
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
	wb.combined.Add(container.NewBorder(nil, container.NewCenter(wb.bottom), container.NewCenter(wb.left), nil, container.NewCenter(wb.canvas)))

	wb.window.SetContent(wb.combined)
	wb.window.Canvas().Focus(wb.ib)
	wb.window.Resize(fyne.NewSize(800, 600))
	wb.window.SetTitle(wb.ib.title)
	wb.window.Content().Refresh()
}

func (wb *WindowBuilder) Build() fyne.Window {
	wb.Refresh()
	return wb.window
}
