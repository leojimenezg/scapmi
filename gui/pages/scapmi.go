package pages

import (
	"os"

	"gioui.org/app"
	"gioui.org/font"
	"gioui.org/font/gofont"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/leojimenezg/scapmi/gui/colors"
	"github.com/leojimenezg/scapmi/gui/pages/copying"
	"github.com/leojimenezg/scapmi/gui/pages/idle"
	"github.com/leojimenezg/scapmi/gui/pages/pasting"
	"github.com/leojimenezg/scapmi/gui/pages/welcome"
	"github.com/leojimenezg/scapmi/internal/vars"
)

var welcomeGUI welcome.WelcomeItems
var idleGUI idle.IdleItems
var copyingGUI copying.CopyingItems
var pastingGUI pasting.PastingItems

type Scapmi struct {
	Window   *app.Window
	Theme    *material.Theme
	Slots    [5]*vars.Slot
	AppState vars.AppState
}

func NewWindow() *Scapmi {
	w := Scapmi{Window: new(app.Window)}
	w.Window.Option(
		app.Title("scapmi"),
		app.Size(unit.Dp(1280), unit.Dp(720)),
		app.MaxSize(unit.Dp(1280), unit.Dp(720)),
		app.MinSize(unit.Dp(1280), unit.Dp(720)),
	)
	w.Theme = material.NewTheme()
	w.Theme.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
	w.Theme.Face = font.Typeface("Go")
	w.Slots = [5]*vars.Slot{
		new(vars.Slot),
		new(vars.Slot),
		new(vars.Slot),
		new(vars.Slot),
		new(vars.Slot),
	}
	w.AppState = vars.StateInit
	return &w
}

func (s *Scapmi) initPages() {
	sourceBtn := &widget.Clickable{}
	docsBtn := &widget.Clickable{}
	slotBtns := [5]*widget.Clickable{
		new(widget.Clickable),
		new(widget.Clickable),
		new(widget.Clickable),
		new(widget.Clickable),
		new(widget.Clickable),
	}

	welcomeGUI.SourceButton = sourceBtn
	welcomeGUI.DocsButton = docsBtn

	idleGUI.SourceButton = sourceBtn
	idleGUI.DocsButton = docsBtn

	copyingGUI.Slots = s.Slots
	copyingGUI.SlotButtons = slotBtns
	copyingGUI.SourceButton = sourceBtn
	copyingGUI.DocsButton = docsBtn

	pastingGUI.Slots = s.Slots
	pastingGUI.SlotButtons = slotBtns
	pastingGUI.SourceButton = sourceBtn
	pastingGUI.DocsButton = docsBtn
}

func (s *Scapmi) Draw() {
	s.initPages()
	for {
		switch e := s.Window.Event().(type) {
		case app.DestroyEvent:
			os.Exit(0)
		case app.FrameEvent:
			var ops op.Ops
			gtx := app.NewContext(&ops, e)
			paint.ColorOp{Color: colors.ColorBackground}.Add(gtx.Ops)
			paint.PaintOp{}.Add(gtx.Ops)
			switch s.AppState {
			case vars.StateInit:
				welcomeGUI.Draw(gtx, s.Theme)
			case vars.StateIdle:
				idleGUI.Draw(gtx, s.Theme)
			case vars.StateCopying:
				copyingGUI.Draw(gtx, s.Theme)
			case vars.StatePasting:
				pastingGUI.Draw(gtx, s.Theme)
			}
			e.Frame(gtx.Ops)
		}
	}
}
