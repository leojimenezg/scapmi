package main

import (
	"fmt"
	"github.com/robotn/gohook"
	"golang.design/x/clipboard"
	"log"
	"os"
)

var errStr string = "ERROR >> %s"
var msgStr string = "MESSAGE >> %s"

func init() {
	err := clipboard.Init()
	if err != nil {
		msg := fmt.Sprintf(errStr, "your clipboard can't be used by scapmi :(")
		log.Panic(msg)
	}
}

func main() {
	setHook()
}

func setHook() {
	msg := fmt.Sprintf(msgStr, `press "Ctrl+Alt+q" any time to stop scapmi`)
	log.Println(msg)

	hook.Register(hook.KeyDown, []string{"ctrl", "alt", "q"}, func(e hook.Event) {
		msg := fmt.Sprintf(msgStr, "goodbye for now!")
		log.Println(msg)
		hook.End()
		os.Exit(0)
	})

	hook.Register(hook.KeyDown, []string{"ctrl", "alt", "c"}, func(e hook.Event) {
		msg := fmt.Sprintf(msgStr, "copy shortcut pressed")
		log.Println(msg)
		// TODO: show interactive GUI to show and select slots in order to save content
		// TODO: handle copied content in slots
	})

	hook.Register(hook.KeyDown, []string{"ctrl", "alt", "v"}, func(e hook.Event) {
		msg := fmt.Sprintf(msgStr, "paste shortcut pressed")
		log.Println(msg)
		// TODO: show interactive GUI to show and select slots in order to paste content
		// TODO: handle slots content and paste it without more user intervention
	})

	s := hook.Start()
	<-hook.Process(s)
}

func showGUI() {

}
