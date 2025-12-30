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
	TextChan  <-chan []byte
	ImageChan <-chan []byte
	MsgChan   <-chan []byte
}

func NewWatcher() *Watcher {
	w := new(Watcher)
	w.Ctx = context.Background()
	w.TextChan = clipboard.Watch(w.Ctx, clipboard.FmtText)
	w.ImageChan = clipboard.Watch(w.Ctx, clipboard.FmtImage)
	w.MsgChan = make(<-chan []byte, 1)
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
			*appState = vars.StateCopying
			window.Invalidate()
		case <-w.ImageChan:
			*appState = vars.StateCopying
			window.Invalidate()
		}
	}
}
