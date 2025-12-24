package components

import (
	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/leojimenezg/scapmi/gui/colors"
	"github.com/leojimenezg/scapmi/gui/widgets"
)

type Message struct {
	Image string
	Title string
	Text  string
}

func NewMessage(gtx layout.Context, theme *material.Theme, message Message) layout.Dimensions {
	return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
		// Image
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints.Max.X = gtx.Dp(unit.Dp(60))
			gtx.Constraints.Max.Y = gtx.Dp(unit.Dp(60))
			return widgets.NewImage(gtx, message.Image, layout.Center, 0)
		}),

		// Spacer
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{
				Width: unit.Dp(10), Height: unit.Dp(1),
			}.Layout(gtx)
		}),

		// Message
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints.Max.X = gtx.Dp(unit.Dp(220))
			return layout.Flex{Axis: layout.Vertical, Alignment: layout.Start}.Layout(gtx,
				// Title
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return widgets.NewLabel(gtx, theme, message.Title, 18, colors.ColorTextSub, text.Start, font.Bold)
				}),

				// Spacer
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Spacer{
						Width: unit.Dp(1), Height: unit.Dp(10),
					}.Layout(gtx)
				}),

				// Text
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return widgets.NewSubText(gtx, theme, message.Text, text.Start)
				}),
			)
		}),
	)
}
