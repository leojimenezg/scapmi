package pasting

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// PastingItems contains all pointers to its interactive items.
type PastingItems struct {
	CardButtons  [5]*widget.Clickable
	SourceButton *widget.Clickable
	DocsButton   *widget.Clickable
}

func (p *PastingItems) Draw(gtx layout.Context, theme *material.Theme) {}
