package widgets

import (
	"image/color"

	"gioui.org/layout"
	txt "gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/leojimenezg/scapmi/gui/colors"
	"github.com/leojimenezg/scapmi/gui/fonts"
)

func NewLabel(gtx layout.Context, theme *material.Theme, text string,
	size int, color color.NRGBA, alignment txt.Alignment) layout.Dimensions {
	label := material.Label(theme, unit.Sp(size), text)
	label.Color = color
	label.Alignment = alignment
	return label.Layout(gtx)
}

func NewMainTitle(gtx layout.Context, theme *material.Theme, text string) layout.Dimensions {
	label := material.Label(theme, fonts.SizeMainTitle, text)
	label.Color = colors.ColorTextMain
	label.Alignment = txt.Middle
	return label.Layout(gtx)
}

func NewTitle(gtx layout.Context, theme *material.Theme, text string) layout.Dimensions {
	label := material.Label(theme, fonts.SizeTitle, text)
	label.Color = colors.ColorTextMain
	label.Alignment = txt.Middle
	return label.Layout(gtx)
}

func NewText(gtx layout.Context, theme *material.Theme, text string, alignment txt.Alignment) layout.Dimensions {
	label := material.Label(theme, fonts.SizeText, text)
	label.Color = colors.ColorTextNormal
	label.Alignment = alignment
	return label.Layout(gtx)
}

func NewSubText(gtx layout.Context, theme *material.Theme, text string, alignment txt.Alignment) layout.Dimensions {
	label := material.Label(theme, fonts.SizeSubtext, text)
	label.Color = colors.ColorTextSub
	label.Alignment = alignment
	return label.Layout(gtx)
}

func NewSmallText(gtx layout.Context, theme *material.Theme, text string) layout.Dimensions {
	label := material.Label(theme, fonts.SizeSmallText, text)
	label.Color = colors.ColorTextSub
	label.Alignment = txt.Middle
	return label.Layout(gtx)
}
