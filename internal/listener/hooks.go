package listener

import (
	"os"

	"gioui.org/app"
	"github.com/leojimenezg/scapmi/internal/vars"
	"github.com/robotn/gohook"
)

func SetHooks(appState *vars.AppState, window *app.Window) {
	hook.Register(hook.KeyDown, []string{"ctrl", "alt", "q"}, func(e hook.Event) {
		hook.End()
		os.Exit(0)
	})

	hook.Register(hook.KeyDown, []string{"ctrl", "alt", "v"}, func(e hook.Event) {
		*appState = vars.StatePasting
		window.Invalidate()
	})

	s := hook.Start()
	<-hook.Process(s)
}
