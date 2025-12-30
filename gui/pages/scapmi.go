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
	"github.com/leojimenezg/scapmi/internal/manager"
	"github.com/leojimenezg/scapmi/internal/vars"
)

var welcomeGUI welcome.WelcomeItems
var idleGUI idle.IdleItems
var copyingGUI copying.CopyingItems
var pastingGUI pasting.PastingItems

type Scapmi struct {
	Manager *manager.Manager
	Theme   *material.Theme
}

func NewWindow() *Scapmi {
	w := new(Scapmi)
	w.Theme = material.NewTheme()
	w.Manager = manager.NewManager()
	return w
}

func (s *Scapmi) Init() {
	s.Theme = material.NewTheme()
	s.Theme.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
	s.Theme.Face = font.Typeface("Go")

	s.Manager.Window.Option(
		app.Title("scapmi"),
		app.Size(unit.Dp(1280), unit.Dp(720)),
		app.MaxSize(unit.Dp(1280), unit.Dp(720)),
		app.MinSize(unit.Dp(1280), unit.Dp(720)),
	)
	s.initPages()
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

	copyingGUI.Manager = s.Manager
	copyingGUI.SlotButtons = slotBtns
	copyingGUI.SourceButton = sourceBtn
	copyingGUI.DocsButton = docsBtn

	pastingGUI.Manager = s.Manager
	pastingGUI.SlotButtons = slotBtns
	pastingGUI.SourceButton = sourceBtn
	pastingGUI.DocsButton = docsBtn
}

func (s *Scapmi) Draw() {
	for {
		switch e := s.Manager.Window.Event().(type) {
		case app.DestroyEvent:
			os.Exit(0)
		case app.FrameEvent:
			var ops op.Ops
			gtx := app.NewContext(&ops, e)
			paint.ColorOp{Color: colors.ColorBackground}.Add(gtx.Ops)
			paint.PaintOp{}.Add(gtx.Ops)
			switch s.Manager.AppState {
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
