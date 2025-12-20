package widgets

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"github.com/leojimenezg/scapmi/internal/utils"
)

func NewImage(gtx layout.Context, name string, position layout.Direction, scale float32) layout.Dimensions {
	image := widget.Image{
		Src:      utils.LoadPNG(name),
		Fit:      widget.ScaleDown,
		Position: position,
		Scale:    scale,
	}
	return image.Layout(gtx)
}
