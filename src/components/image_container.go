package components

import (
	"image"
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type ImageContainer struct {
	height  float32
	width   float32
	content *canvas.Image
}

func NewImageContainer(height, width float32) ImageContainer {
	blank := image.NewRGBA(image.Rect(0, 0, int(height), int(width)))
	imgContainer := canvas.NewImageFromImage(blank)
	imgContainer.FillMode = canvas.ImageFillContain
	imgContainer.SetMinSize(fyne.NewSize(float32(height), float32(width)))
	return ImageContainer{
		height:  height,
		width:   width,
		content: imgContainer,
	}
}

func (ic *ImageContainer) UpdateContent(reader fyne.URIReadCloser, err error) {
	if err != nil {
		return
	}
	if reader == nil {
		return
	}

	data, err := io.ReadAll(reader)
	if err != nil {
		return
	}
	res := fyne.NewStaticResource(reader.URI().Name(), data)

	ic.content.Resource = res
	ic.content.Refresh()

}
