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
	Width      int
	Height     int
	Radius     int
	StokeWidth int
	StokeColor color.NRGBA
	Color      color.NRGBA
}

func NewCard(gtx layout.Context, c Card, widget func(gtx layout.Context) layout.Dimensions) layout.Dimensions {
	gtx.Constraints.Min.X = gtx.Dp(unit.Dp(c.Width))
	gtx.Constraints.Min.Y = gtx.Dp(unit.Dp(c.Height))
	gtx.Constraints.Max.X = gtx.Dp(unit.Dp(c.Width))
	gtx.Constraints.Max.Y = gtx.Dp(unit.Dp(c.Height))
	radius := unit.Dp(c.Radius)
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
