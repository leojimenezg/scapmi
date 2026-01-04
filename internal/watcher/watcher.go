package watcher

import (
	"context"
	"log"

	"gioui.org/app"
	"github.com/leojimenezg/scapmi/internal/vars"
	"golang.design/x/clipboard"
)

type WatcherType int

const (
	NoType = iota
	TextType
	ImageType
)

type Watcher struct {
	Context      context.Context
	IgnoreChange bool
	CurrentType  WatcherType
	TextChan     <-chan []byte
	ImageChan    <-chan []byte
	DetectChan   chan bool
}

func NewWatcher() *Watcher {
	w := new(Watcher)
	w.Context = context.Background()
	w.IgnoreChange = false
	w.TextChan = clipboard.Watch(w.Context, clipboard.FmtText)
	w.ImageChan = clipboard.Watch(w.Context, clipboard.FmtImage)
	w.DetectChan = make(chan bool, 1)
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
				w.CurrentType = NoType
				w.IgnoreChange = false
				continue
			}
			w.CurrentType = TextType
			w.DetectChan <- true
		case <-w.ImageChan:
			if w.IgnoreChange {
				w.CurrentType = NoType
				w.IgnoreChange = false
				continue
			}
			w.CurrentType = ImageType
			w.DetectChan <- true
		}
	}
}
