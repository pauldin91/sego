package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/pauldin91/sego/src/common"
	"github.com/pauldin91/sego/src/components/controls"
	browser "github.com/pauldin91/sego/src/components/viewer"
)

type WindowBuilder struct {
	widget.BaseWidget
	window   fyne.Window
	ib       *browser.ImageBrowser
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
	result.ib = browser.NewImageBrowser(result.window)
	result.ExtendBaseWidget(result)
	return result
}
func (wb *WindowBuilder) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(wb.combined)
}

func (wb *WindowBuilder) WithBottomMenu() *WindowBuilder {
	res := controls.NewBottomMenu(wb.ib, wb.window).
		WithButtons(common.ColorBtn, common.Toggle).
		WithButtons(common.IncBtn, common.DecBtn).
		WithButtons(common.SaveBtn, common.ClearBtn).
		WithFloatWidget(wb.ib.ChangeBrushSize).
		Build()

	wb.left.Add(res)
	return wb
}

func (wb *WindowBuilder) WithDefaultCanvas() *WindowBuilder {
	wb.canvas.Add(wb.ib)
	return wb
}

func (wb *WindowBuilder) WithNextBtn() *fyne.Container {
	next := widget.NewButton("", wb.ib.GetNext)
	next.Icon = theme.NavigateNextIcon()
	next.Resize(common.DefaultVButtonSize)
	return container.NewStack(next)
}

func (wb *WindowBuilder) WithPrevBtn() *fyne.Container {
	prev := widget.NewButton("", wb.ib.GetPrevious)
	prev.Icon = theme.NavigateBackIcon()
	prev.Resize(common.DefaultVButtonSize)
	return container.NewStack(prev)

}

func (wb *WindowBuilder) Refresh() {
	wb.combined = container.NewBorder(
		nil,
		container.NewCenter(wb.left),
		wb.WithPrevBtn(),
		wb.WithNextBtn(),
		container.NewStack(wb.ib),
	)

	wb.window.SetContent(wb.combined)
	wb.window.Canvas().Focus(wb.ib)
	wb.window.Resize(wb.calcWindowSize())
	wb.window.SetTitle(wb.ib.GetTitle())
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
	open := fyne.NewMenuItem("Open Folder", wb.ib.OnOpenFolderButtonClicked)
	load := fyne.NewMenuItem("Load Image", wb.ib.OnLoadFileButtonClicked)
	menu := fyne.NewMainMenu(fyne.NewMenu("Main Menu", open, load))
	wb.window.SetMainMenu(menu)

	return wb
}
