package main

import (
	"gioui.org/app"
	"github.com/leojimenezg/scapmi/gui/pages"
	"github.com/leojimenezg/scapmi/internal/listener"
	"github.com/leojimenezg/scapmi/internal/watcher"
)

var Scapmi pages.Scapmi

func main() {
	w := watcher.NewWatcher()
	w.Init()

	l := listener.NewListener()

	Scapmi = *pages.NewWindow()

	go w.WatchClipboard(&Scapmi.AppState, Scapmi.Window)
	go l.SetHooks(&Scapmi.AppState, Scapmi.Window)
	go Scapmi.Draw()

	app.Main()
}
