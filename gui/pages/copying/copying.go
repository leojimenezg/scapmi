package copying

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// CopyingItems contains all pointers to its interactive items.
type CopyingItems struct {
	CardButtons  [5]*widget.Clickable
	SourceButton *widget.Clickable
	DocsButton   *widget.Clickable
}

func (c *CopyingItems) Draw(gtx layout.Context, theme *material.Theme) {}
