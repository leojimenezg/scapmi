package listener

import (
	"os"

	"gioui.org/app"
	"github.com/leojimenezg/scapmi/internal/vars"
	"github.com/robotn/gohook"
)

type Listener struct {
	DetectChan chan bool
}

func NewListener() *Listener {
	l := new(Listener)
	l.DetectChan = make(chan bool, 1)
	return l
}

func (l *Listener) SetHooks(appState *vars.AppState, window *app.Window) {
	hook.Register(hook.KeyDown, []string{"ctrl", "alt", "q"}, func(e hook.Event) {
		hook.End()
		os.Exit(0)
	})

	hook.Register(hook.KeyDown, []string{"ctrl", "alt", "v"}, func(e hook.Event) {
		l.DetectChan <- true
	})

	s := hook.Start()
	<-hook.Process(s)
}
