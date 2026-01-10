package manager

import (
	"strings"
	"time"

	"gioui.org/app"
	"gioui.org/io/system"
	"github.com/leojimenezg/scapmi/internal/listener"
	"github.com/leojimenezg/scapmi/internal/vars"
	"github.com/leojimenezg/scapmi/internal/watcher"
	"golang.design/x/clipboard"
)

type Manager struct {
	Slots    [5]*vars.Slot
	Window   *app.Window
	Watcher  *watcher.Watcher
	Listener *listener.Listener
	CTimer   *time.Timer
	PTimer   *time.Timer
	AppState vars.AppState
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
	if number < 0 || number > 4 {
		number = 0
	}
	var slotType vars.SlotType
	switch m.Watcher.CurrentType {
	case watcher.TextType:
		slotType = vars.TypeText
		m.Slots[number].Content = clipboard.Read(clipboard.FmtText)
	case watcher.ImageType:
		slotType = vars.TypeImage
		m.Slots[number].Content = clipboard.Read(clipboard.FmtImage)
	default:
		return
	}
	m.Slots[number].Type = slotType
	m.Slots[number].HasContent = true
	m.Slots[number].Summary = m.createSlotSummary(number)
	m.AppState = vars.StateIdle
	m.Window.Invalidate()
}

func (m *Manager) LoadFromSlot(number int) {
	if number < 0 || number > 4 {
		number = 0
	}
	m.Watcher.IgnoreChange = true
	switch m.Slots[number].Type {
	case vars.TypeText:
		clipboard.Write(clipboard.FmtText, m.Slots[number].Content)
	case vars.TypeImage:
		clipboard.Write(clipboard.FmtImage, m.Slots[number].Content)
	}
}

func (m *Manager) createSlotSummary(number int) string {
	if !m.Slots[number].HasContent {
		return ""
	}
	maxChars := 25
	n := min(len(m.Slots[number].Content), maxChars)
	summary := string(m.Slots[number].Content[:n])
	summary = strings.ReplaceAll(summary, "\n", " ")
	if len(summary) >= maxChars {
		summary += " ..."
	}
	return summary
}

func (m *Manager) ListenForEvents() {
	for {
		select {
		case <-m.Watcher.DetectChan:
			m.AppState = vars.StateCopying
			m.Window.Invalidate()
			m.Window.Perform(system.ActionRaise)
			m.startCopyingTimer()
		case <-m.Listener.DetectChan:
			m.AppState = vars.StatePasting
			m.Window.Invalidate()
			m.Window.Perform(system.ActionRaise)
			m.startPastingTimer()
		}
	}
}

func (m *Manager) startCopyingTimer() {
	if m.CTimer != nil {
		m.CTimer.Stop()
	}
	m.CTimer = time.AfterFunc(5*time.Second, func() {
		if m.AppState == vars.StateCopying {
			m.SaveToSlot(0)
		}
		m.CTimer = nil
	})
}

func (m *Manager) startPastingTimer() {
	if m.PTimer != nil {
		m.PTimer.Stop()
	}
	m.PTimer = time.AfterFunc(10*time.Second, func() {
		if m.AppState == vars.StatePasting {
			m.AppState = vars.StateIdle
			m.Window.Invalidate()
		}
		m.PTimer = nil
	})
}
