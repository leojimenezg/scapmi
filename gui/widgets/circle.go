package widgets

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

type Circle struct {
	Diameter    int
	Color       color.NRGBA
	StrokeColor color.NRGBA
	StrokeWidth int
}

func NewCircle(gtx layout.Context, circle Circle) layout.Dimensions {
	outerDiameter := unit.Dp(circle.Diameter + circle.StrokeWidth*2)
	outerSize := image.Pt(gtx.Dp(outerDiameter), gtx.Dp(outerDiameter))
	outerCircle := clip.Ellipse{Max: outerSize}.Op(gtx.Ops)
	paint.FillShape(gtx.Ops, circle.StrokeColor, outerCircle)

	offset := gtx.Dp(unit.Dp(circle.StrokeWidth))
	stack := op.Offset(image.Pt(offset, offset)).Push(gtx.Ops)

	innerDiameter := unit.Dp(circle.Diameter)
	innerSize := image.Pt(gtx.Dp(innerDiameter), gtx.Dp(innerDiameter))
	innerCircle := clip.Ellipse{Max: innerSize}.Op(gtx.Ops)
	paint.FillShape(gtx.Ops, circle.Color, innerCircle)

	stack.Pop()
	return layout.Dimensions{Size: outerSize}
}
