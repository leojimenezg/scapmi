package listener

import (
	"os"

	"gioui.org/app"
	"gioui.org/io/system"
	"github.com/leojimenezg/scapmi/internal/vars"
	"github.com/robotn/gohook"
)

type Listener struct{}

func NewListener() *Listener {
	return new(Listener)
}

func (l *Listener) SetHooks(appState *vars.AppState, window *app.Window) {
	hook.Register(hook.KeyDown, []string{"ctrl", "alt", "q"}, func(e hook.Event) {
		hook.End()
		os.Exit(0)
	})

	hook.Register(hook.KeyDown, []string{"ctrl", "alt", "v"}, func(e hook.Event) {
		*appState = vars.StatePasting
		window.Invalidate()
		window.Perform(system.ActionRaise)
	})

	s := hook.Start()
	<-hook.Process(s)
}
