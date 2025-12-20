package idle

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/leojimenezg/scapmi/gui/components"
	"github.com/leojimenezg/scapmi/gui/widgets"
)

// IdleItems contains all pointers to its interactive items
type IdleItems struct {
	SourceButton *widget.Clickable
	DocsButton   *widget.Clickable
}

func (i *IdleItems) Draw(gtx layout.Context, theme *material.Theme) {
	layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
		// Header
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return components.NewHeader(gtx, theme, "idle")
		}),

		// Spacer
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{
				Width: unit.Dp(gtx.Constraints.Max.X), Height: unit.Dp(150),
			}.Layout(gtx)
		}),

		// Title
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return widgets.NewTitle(gtx, theme, "App shortcuts")
			})
		}),

		// Spacer
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{
				Width: unit.Dp(gtx.Constraints.Max.X), Height: unit.Dp(10),
			}.Layout(gtx)
		}),

		// Shortcuts
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return components.NewShortcuts(gtx, theme)
			})
		}),

		// Spacer
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{
				Width: unit.Dp(gtx.Constraints.Max.X), Height: unit.Dp(250),
			}.Layout(gtx)
		}),

		// Footer
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				f := components.FooterItems{
					SourceButton: i.SourceButton,
					DocsButton:   i.DocsButton,
				}
				return components.NewFooter(gtx, theme, f)
			})
		}),
	)
}
