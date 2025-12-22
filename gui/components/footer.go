package components

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/leojimenezg/scapmi/gui/widgets"
)

// FooterItems contains all the pointers to its interactive items.
type FooterItems struct {
	SourceButton *widget.Clickable
	DocsButton   *widget.Clickable
}

func NewFooter(gtx layout.Context, theme *material.Theme, obj FooterItems) layout.Dimensions {
	return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
		// Version
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return widgets.NewSmallText(gtx, theme, "v1.0.0")
		}),

		// Spacer
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			s := image.Pt(gtx.Dp(unit.Dp(10)), 0)
			return layout.Dimensions{Size: s}
		}),

		// Source
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			source := widgets.Link{
				Button: obj.SourceButton,
				Theme:  theme,
				Text:   "Source",
				Url:    "https://github.com/leojimenezg/scapmi",
			}
			return widgets.NewLink(gtx, source)
		}),

		// Spacer
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			s := image.Pt(gtx.Dp(unit.Dp(10)), 0)
			return layout.Dimensions{Size: s}
		}),

		// Docs
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			source := widgets.Link{
				Button: obj.DocsButton,
				Theme:  theme,
				Text:   "Docs",
				Url:    "https://github.com/leojimenezg/scapmi/blob/main/README.md",
			}
			return widgets.NewLink(gtx, source)
		}),
	)
}
