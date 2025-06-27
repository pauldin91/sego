package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type WindowBuilder struct {
	widget.BaseWidget
	window     fyne.Window
	ib         *ImageBrowser
	top        *fyne.Container
	bottom     *fyne.Container
	combined   *fyne.Container
	canvasSize fyne.Size
}

func NewWindowBuilder(size fyne.Size, title string, a fyne.App) *WindowBuilder {

	result := &WindowBuilder{
		window:     a.NewWindow(title),
		top:        container.NewHBox(),
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
	res := NewSidebarMenu(wb.ib)
	wb.top.Add(res.getBrushButtons())
	return wb
}

func (wb *WindowBuilder) WithBottomMenu() *WindowBuilder {
	res := NewBottomMenu(wb.ib, wb.window)
	wb.bottom.Add(res.getButtons())
	return wb
}

func (wb *WindowBuilder) WithDefaultCanvas() *WindowBuilder {
	wb.top.Add(wb.ib)
	return wb
}

func (wb *WindowBuilder) Refresh() {
	wb.combined.Add(wb.top)
	wb.combined.Add(wb.bottom)

	wb.window.SetContent(container.NewCenter(wb.combined))
	wb.window.Canvas().Focus(wb.ib)
	wb.window.Resize(wb.canvasSize)
	wb.window.SetTitle(wb.ib.title)
	wb.window.Content().Refresh()
}

func (wb *WindowBuilder) Build() fyne.Window {
	wb.Refresh()
	return wb.window
}
