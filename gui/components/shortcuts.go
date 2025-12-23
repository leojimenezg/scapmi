package components

import (
	"image"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/leojimenezg/scapmi/gui/colors"
	"github.com/leojimenezg/scapmi/gui/widgets"
)

var iconSize = unit.Dp(17)

func NewShortcuts(gtx layout.Context, theme *material.Theme) layout.Dimensions {
	card := widgets.Card{
		Width:       500,
		Height:      190,
		Radius:      15,
		StrokeWidth: 2,
		StrokeColor: colors.ColorBackgroundHover,
		Color:       colors.ColorBackgroundLight,
	}
	return widgets.NewCard(gtx, card, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
			// First row
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle, WeightSum: 1}.Layout(gtx,
						// Icon
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							gtx.Constraints.Max.X = gtx.Dp(iconSize)
							gtx.Constraints.Max.Y = gtx.Dp(iconSize)
							return widgets.NewImage(gtx, "img1.png", layout.Center, 0)
						}),

						// Spacer
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							s := image.Pt(gtx.Dp(unit.Dp(10)), 0)
							return layout.Dimensions{Size: s}
						}),

						// Description
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return widgets.NewText(gtx, theme, "Select slot", text.Middle)
						}),

						// Command
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							return widgets.NewText(gtx, theme, "Cmd + c", text.End)
						}))
				})
			}),

			// Divider (line)
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return widgets.NewDivider(gtx, 2, colors.ColorBackgroundHover)
			}),

			// Second row
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle, WeightSum: 1}.Layout(gtx,
						// Icon
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							gtx.Constraints.Max.X = gtx.Dp(iconSize)
							gtx.Constraints.Max.Y = gtx.Dp(iconSize)
							return widgets.NewImage(gtx, "img2.png", layout.Center, 0)
						}),

						// Spacer
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							s := image.Pt(gtx.Dp(unit.Dp(10)), 0)
							return layout.Dimensions{Size: s}
						}),

						// Description
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return widgets.NewText(gtx, theme, "Select content", text.Middle)
						}),

						// Command
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							return widgets.NewText(gtx, theme, "Ctrl + Alt + v", text.End)
						}))
				})
			}),

			// Divider (line)
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return widgets.NewDivider(gtx, 2, colors.ColorBackgroundHover)
			}),

			// Third row
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle, WeightSum: 1}.Layout(gtx,
						// Icon
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							gtx.Constraints.Max.X = gtx.Dp(iconSize)
							gtx.Constraints.Max.Y = gtx.Dp(iconSize)
							return widgets.NewImage(gtx, "img3.png", layout.Center, 0)
						}),

						// Spacer
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							s := image.Pt(gtx.Dp(unit.Dp(10)), 0)
							return layout.Dimensions{Size: s}
						}),

						// Description
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return widgets.NewText(gtx, theme, "Exit application", text.Middle)
						}),

						// Command
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							return widgets.NewText(gtx, theme, "Ctrl + Alt + q", text.End)
						}))
				})
			}),
		)
	})
}
