package manager

import (
	"gioui.org/app"
	"github.com/leojimenezg/scapmi/internal/listener"
	"github.com/leojimenezg/scapmi/internal/vars"
	"github.com/leojimenezg/scapmi/internal/watcher"
)

type Manager struct {
	Window   *app.Window
	AppState *vars.AppState
	Slots    [5]*vars.Slot
	Watcher  *watcher.Watcher
	Listener *listener.Listener
}
