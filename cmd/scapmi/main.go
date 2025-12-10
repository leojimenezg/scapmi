package main

import (
	"context"
	"fmt"
	"gioui.org/app"
	"gioui.org/font/gofont"
	"log"
	"os"
	// "gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
	"github.com/robotn/gohook"
	"golang.design/x/clipboard"
)

// ============================================================
// Constants
// ============================================================

const (
	StateInit AppState = iota
	StateIdle
	StateCopying
	StatePasting
)

const errStr string = "ERROR >> %s"

// ============================================================
// Named Types
// ============================================================

type AppState int

type Slot struct {
	Content    []byte
	HasContent bool
}

// ============================================================
// Variables
// ============================================================

var Window *app.Window

var CurrentAppState AppState

var Slots [5]Slot

func init() {
	err := clipboard.Init()
	if err != nil {
		msg := fmt.Sprintf(errStr, "your clipboard can't be used by scapmi :(")
		log.Fatal(msg)
	}
	Window = new(app.Window)
	if Window == nil {
		msg := fmt.Sprintf(errStr, "could not allocate memory for the window program")
		log.Fatal(msg)
	}
	Window.Option(app.Title("scapmi"))
	CurrentAppState = StateInit
}

func main() {
	go ShowGUI()
	go SetHook()
	go WatchClipboard()
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
			switch CurrentAppState {
			case StateInit:
				material.H1(th, "Welcome").Layout(gtx)
			case StateIdle:
				material.H1(th, "Idle State").Layout(gtx)
			case StateCopying:
				material.H1(th, "Copying State").Layout(gtx)
				// TODO: Show buttons to select a slot and their state
			case StatePasting:
				material.H1(th, "Pasting State").Layout(gtx)
				// TODO: Show slots with content and buttons to select them
			}
			e.Frame(gtx.Ops)
		}
	}
}

func SetHook() {
	hook.Register(hook.KeyDown, []string{"ctrl", "alt", "q"}, func(e hook.Event) {
		hook.End()
		os.Exit(0)
	})
	hook.Register(hook.KeyDown, []string{"ctrl", "alt", "v"}, func(e hook.Event) {
		CurrentAppState = StatePasting
		Window.Invalidate()
	})
	s := hook.Start()
	<-hook.Process(s)
}

func WatchClipboard() {
	ctx := context.Background()
	for {
		txt := clipboard.Watch(ctx, clipboard.FmtText)
		img := clipboard.Watch(ctx, clipboard.FmtImage)
		select {
		case <-txt:
			CurrentAppState = StateCopying
			Window.Invalidate()
		case <-img:
			CurrentAppState = StateCopying
			Window.Invalidate()
		}
	}
}
