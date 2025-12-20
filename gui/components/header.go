package components

import (
	"fmt"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/leojimenezg/scapmi/gui/colors"
	"github.com/leojimenezg/scapmi/gui/widgets"
)

func NewHeader(gtx layout.Context, theme *material.Theme, state string) layout.Dimensions {
	inset := layout.UniformInset(unit.Dp(10))
	return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle, WeightSum: 1}.Layout(gtx,
			// Spacer
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Spacer{Width: unit.Dp(20), Height: unit.Dp(1)}.Layout(gtx)
			}),

			// Logo
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints.Max.X = gtx.Dp(unit.Dp(35))
				gtx.Constraints.Max.Y = gtx.Dp(unit.Dp(35))
				return widgets.NewImage(gtx, "logo.png", layout.Center, 0)
			}),

			// Spacer
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Spacer{Width: unit.Dp(10), Height: unit.Dp(1)}.Layout(gtx)
			}),

			// Title
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return widgets.NewLabel(
					gtx, theme, "scapmi", 22, colors.ColorTextMain, text.Start, font.Bold,
				)
			}),

			// Status (dependent)
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				return widgets.NewSubText(gtx, theme, fmt.Sprintf("Status: %s", state), text.End)
			}),

			// Spacer
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Spacer{Width: unit.Dp(20), Height: unit.Dp(1)}.Layout(gtx)
			}),
		)
	},
	)
}
