package main

import (
	"log"

	"gioui.org/app"
	"github.com/leojimenezg/scapmi/gui/pages"
	"github.com/leojimenezg/scapmi/internal/listener"
	"github.com/leojimenezg/scapmi/internal/watcher"
)

var Scapmi pages.Scapmi

func main() {
	err := watcher.InitClipboard()
	if err != nil {
		log.Fatal(err.Error())
	}

	Scapmi = *pages.NewWindow()

	go watcher.WatchClipboard(&Scapmi.AppState, Scapmi.Window)
	go listener.SetHooks(&Scapmi.AppState, Scapmi.Window)
	go Scapmi.Draw()

	app.Main()
}
