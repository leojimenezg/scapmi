package widgets

import (
	"log"

	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
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
		// TODO: Implement logic to open browser and go to the url
		log.Printf("%s link was clicked", obj.Text)
	}
	return obj.Button.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return NewSmallText(gtx, obj.Theme, obj.Text)
	})
}
