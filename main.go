package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/leojimenezg/scapmi/gui/colors"
	"github.com/leojimenezg/scapmi/gui/pages"
	"github.com/leojimenezg/scapmi/gui/pages/idle"
	"github.com/leojimenezg/scapmi/gui/pages/welcome"
	"github.com/leojimenezg/scapmi/internal/listener"
	"github.com/leojimenezg/scapmi/internal/vars"
	"github.com/leojimenezg/scapmi/internal/watcher"
)

var CurrentAppState vars.AppState
var Window *app.Window

var WelcomeGUI welcome.WelcomeItems
var IdleGUI idle.IdleItems

func main() {
	err := watcher.InitClipboard()
	if err != nil {
		log.Fatal(err.Error())
	}
	Window = pages.NewWindow()
	CurrentAppState = vars.StateIdle

	sourceButton := &widget.Clickable{}
	docsButton := &widget.Clickable{}

	WelcomeGUI = welcome.WelcomeItems{
		SourceButton: sourceButton,
		DocsButton:   docsButton,
	}
	IdleGUI = idle.IdleItems{
		SourceButton: sourceButton,
		DocsButton:   docsButton,
	}

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
				WelcomeGUI.Draw(gtx, th)
			case vars.StateIdle:
				IdleGUI.Draw(gtx, th)
			case vars.StateCopying:
				material.H1(th, "Copying State").Layout(gtx)
			case vars.StatePasting:
				material.H1(th, "Pasting State").Layout(gtx)
			}

			e.Frame(gtx.Ops)
		}
	}
}
