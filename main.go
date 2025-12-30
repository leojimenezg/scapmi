package main

import (
	"gioui.org/app"
	"github.com/leojimenezg/scapmi/gui/pages"
)

var Scapmi pages.Scapmi

func main() {
	// Create all variables and objects needed.
	Scapmi = *pages.NewWindow()

	// Initialize clipboard to know if is usable.
	Scapmi.Manager.Watcher.Init()

	// Configure window and GUI pages.
	Scapmi.Init()

	// Set the clipboard watcher in a different goroutine.
	Scapmi.Manager.SetWatcher()

	// Set the event listener in a different goroutine.
	Scapmi.Manager.SetListener()

	// Initialize GUI loop.
	go Scapmi.Draw()

	app.Main()
}
