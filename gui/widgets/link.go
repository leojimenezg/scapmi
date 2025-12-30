package widgets

import (
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/leojimenezg/scapmi/internal/utils"
)

type Link struct {
	Button *widget.Clickable
	Theme  *material.Theme
	Text   string
	Url    string
}

func NewLink(gtx layout.Context, obj Link) layout.Dimensions {
	if obj.Button.Hovered() {
		pointer.Cursor(pointer.CursorPointer).Add(gtx.Ops)
	}
	if obj.Button.Clicked(gtx) {
		utils.OpenURL(obj.Url)
	}
	return obj.Button.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return NewSmallText(gtx, obj.Theme, obj.Text)
	})
}
