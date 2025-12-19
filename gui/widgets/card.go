package widgets

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	w "gioui.org/widget"
)

type Card struct {
	Size         image.Point
	CornerRadius int
	StokeWidth   int
	StokeColor   color.NRGBA
	Color        color.NRGBA
}

func NewCard(gtx layout.Context, c Card, widget func(gtx layout.Context) layout.Dimensions) layout.Dimensions {
	gtx.Constraints.Min.X = c.Size.X
	gtx.Constraints.Min.Y = c.Size.Y
	gtx.Constraints.Max.X = c.Size.X
	gtx.Constraints.Max.Y = c.Size.Y
	radius := unit.Dp(c.CornerRadius)
	return w.Border{
		Color:        c.StokeColor,
		CornerRadius: radius,
		Width:        unit.Dp(c.StokeWidth),
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		inside := clip.UniformRRect(image.Rectangle{
			Max: image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Max.Y),
		}, gtx.Dp(radius)).Op(gtx.Ops)
		paint.FillShape(gtx.Ops, c.Color, inside)
		return widget(gtx)
	})
}
