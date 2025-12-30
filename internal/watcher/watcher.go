package watcher

import (
	"context"
	"log"

	"gioui.org/app"
	"github.com/leojimenezg/scapmi/internal/vars"
	"golang.design/x/clipboard"
)

type Watcher struct {
	Ctx       context.Context
	IgnoreChange bool
	TextChan  <-chan []byte
	ImageChan <-chan []byte
}

func NewWatcher() *Watcher {
	w := new(Watcher)
	w.Ctx = context.Background()
	w.IgnoreChange = false
	w.TextChan = clipboard.Watch(w.Ctx, clipboard.FmtText)
	w.ImageChan = clipboard.Watch(w.Ctx, clipboard.FmtImage)
	return w
}

func (w *Watcher) Init() {
	err := clipboard.Init()
	if err != nil {
		log.Fatalf("your clipboard can't be used: %s", err.Error())
	}
}

func (w *Watcher) WatchClipboard(appState *vars.AppState, window *app.Window) {
	for {
		select {
		case <-w.TextChan:
			if w.IgnoreChange {
				w.IgnoreChange = false
				continue
			}
			*appState = vars.StateCopying
			window.Invalidate()
		case <-w.ImageChan:
			if w.IgnoreChange {
				w.IgnoreChange = false
				continue
			}
			*appState = vars.StateCopying
			window.Invalidate()
		}
	}
}
