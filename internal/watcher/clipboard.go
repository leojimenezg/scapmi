package watcher

import (
	"context"
	"fmt"

	"gioui.org/app"
	"github.com/leojimenezg/scapmi/internal/vars"
	"golang.design/x/clipboard"
)

type InitErr struct {
	Err error
}

func (e InitErr) Error() string {
	return fmt.Sprintf("your clipboard can't be used: %v", e.Err)
}

func InitClipboard() error {
	err := clipboard.Init()
	if err != nil {
		return InitErr{Err: err}
	}
	return nil
}

func WatchClipboard(appState *vars.AppState, window *app.Window) {
	ctx := context.Background()
	for {
		txt := clipboard.Watch(ctx, clipboard.FmtText)
		img := clipboard.Watch(ctx, clipboard.FmtImage)
		select {
		case <-txt:
			*appState = vars.StateCopying
			window.Invalidate()

		case <-img:
			*appState = vars.StateCopying
			window.Invalidate()
		}
	}
}
