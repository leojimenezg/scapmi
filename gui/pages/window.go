package pages

import (
	"gioui.org/app"
	"gioui.org/unit"
)

func NewWindow() *app.Window {
	window := new(app.Window)
	window.Option(
		app.Title("scapmi"),
		app.Size(unit.Dp(1280), unit.Dp(720)),
		app.MaxSize(unit.Dp(1280), unit.Dp(720)),
		app.MinSize(unit.Dp(1280), unit.Dp(720)),
	)
	return window
}
