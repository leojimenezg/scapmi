package manager

import (
	"log"

	"gioui.org/app"
	"github.com/leojimenezg/scapmi/internal/listener"
	"github.com/leojimenezg/scapmi/internal/vars"
	"github.com/leojimenezg/scapmi/internal/watcher"
)

type Manager struct {
	Window   *app.Window
	AppState vars.AppState
	Slots    [5]*vars.Slot
	Watcher  *watcher.Watcher
	Listener *listener.Listener
}

func NewManager() *Manager {
	m := new(Manager)
	m.Window = new(app.Window)
	m.AppState = vars.StateInit
	m.Slots = [5]*vars.Slot{
		new(vars.Slot),
		new(vars.Slot),
		new(vars.Slot),
		new(vars.Slot),
		new(vars.Slot),
	}
	m.Watcher = watcher.NewWatcher()
	m.Listener = listener.NewListener()
	return m
}

func (m *Manager) SetWatcher() {
	go m.Watcher.WatchClipboard(&m.AppState, m.Window)
}

func (m *Manager) SetListener() {
	go m.Listener.SetHooks(&m.AppState, m.Window)
}

func (m *Manager) SaveToSlot(number int) {
	// TODO: Implement the actual logic to save the content in the selected slot.
	log.Printf("SaveToSlot function used from Slot: %d", number)
}

func (m *Manager) LoadFromSlot(number int) {
	// TODO: Implement the actual logic to load the content from the selected slot.
	log.Printf("LoadFromSlot function used from Slot: %d", number)
}
