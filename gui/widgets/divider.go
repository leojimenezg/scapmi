package widgets

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

func NewDivider(gtx layout.Context, height int, color color.NRGBA) layout.Dimensions {
	h := gtx.Dp(unit.Dp(height))
	maxPoint := image.Pt(gtx.Constraints.Max.X, h)
	line := clip.Rect{Max: maxPoint}.Op()
	paint.FillShape(gtx.Ops, color, line)
	return layout.Dimensions{Size: maxPoint}
}
