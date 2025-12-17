package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/leojimenezg/scapmi/gui/assets"
	"github.com/leojimenezg/scapmi/gui/colors"
	"github.com/leojimenezg/scapmi/gui/interfaces"
	"github.com/leojimenezg/scapmi/internal/listener"
	"github.com/leojimenezg/scapmi/internal/utils"
	"github.com/leojimenezg/scapmi/internal/vars"
	"github.com/leojimenezg/scapmi/internal/watcher"
)

var CurrentAppState vars.AppState
var Window *app.Window

func main() {
	err := watcher.InitClipboard()
	if err != nil {
		log.Fatal(err.Error())
	}
	Window = interfaces.NewWindow()
	go watcher.WatchClipboard(&CurrentAppState, Window)
	go listener.SetHooks(&CurrentAppState, Window)
	go ShowGUI()
	app.Main()
}

func ShowGUI() {
	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
	for {
		switch e := Window.Event().(type) {
		case app.DestroyEvent:
			os.Exit(0)
		case app.FrameEvent:
			var ops op.Ops
			gtx := app.NewContext(&ops, e)
			paint.ColorOp{Color: colors.ColorBackground}.Add(gtx.Ops)
			paint.PaintOp{}.Add(gtx.Ops)

			switch CurrentAppState {
			case vars.StateInit:
				layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						inset := layout.Inset{
							Top:    unit.Dp(30),
							Bottom: unit.Dp(10),
						}
						return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return widget.Image{
								Src:      utils.LoadImage(&assets.PngImgs, "logo.png"),
								Fit:      widget.Unscaled,
								Position: layout.Center,
								Scale:    .2,
							}.Layout(gtx)
						})
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						title := material.Label(th, unit.Sp(38), "Welcome to scapmi")
						title.Color = colors.ColorTextMain
						title.Alignment = text.Middle
						return title.Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						sub := material.Label(th, unit.Sp(14), "Simultaneously Copy And Paste Multiple Items")
						sub.Color = colors.ColorTextSub
						sub.Alignment = text.Middle
						return sub.Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						inset := layout.Inset{Top: unit.Dp(50)}
						return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							subtitle := material.Label(th, unit.Sp(24), "App Shortcuts")
							subtitle.Color = colors.ColorTextMain
							subtitle.Alignment = text.Middle
							return subtitle.Layout(gtx)
						})
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Dimensions{}
					}))

			case vars.StateIdle:
				material.H1(th, "Idle State").Layout(gtx)

			case vars.StateCopying:
				material.H1(th, "Copying State").Layout(gtx)

			case vars.StatePasting:
				material.H1(th, "Pasting State").Layout(gtx)
			}

			e.Frame(gtx.Ops)
		}
	}
}
