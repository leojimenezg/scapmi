package pasting

import (
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/leojimenezg/scapmi/gui/components"
	"github.com/leojimenezg/scapmi/gui/widgets"
	"github.com/leojimenezg/scapmi/internal/vars"
)

// PastingItems contains all pointers to its interactive items.
type PastingItems struct {
	Slots        [5]*vars.Slot
	SlotButtons  [5]*widget.Clickable
	SourceButton *widget.Clickable
	DocsButton   *widget.Clickable
}

func (p *PastingItems) Draw(gtx layout.Context, theme *material.Theme) {
	layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
		// Header
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return components.NewHeader(gtx, theme, "pasting")
		}),

		// Title
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return widgets.NewTitle(gtx, theme, "Usable slots")
		}),

		// Spacer
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{
				Width: unit.Dp(gtx.Constraints.Max.X), Height: unit.Dp(5),
			}.Layout(gtx)
		}),

		// Description
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return widgets.NewSubText(
				gtx, theme, "Select a slot to place its content in your clipboard", text.Middle)
		}),

		// Spacer
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{
				Width: unit.Dp(gtx.Constraints.Max.X), Height: unit.Dp(60),
			}.Layout(gtx)
		}),

		// Slots
		// TODO: Use a function or something to get a useful desription of the slot content, and not see just bytes (numbers).
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
					// Slot 1
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						slot := components.Slot{
							Button: p.SlotButtons[0],
							Number: 0,
							Text:   "description of content...",
							Extra:  "Default",
						}
						slot.State = components.EmptyUnusable
						if p.Slots[0].HasContent {
							slot.State = components.FilledUsable
						}
						slot.Type = components.SlotText
						if p.Slots[0].Type == vars.TypeImage {
							slot.Type = components.SlotImage
						}
						return components.NewSlot(gtx, theme, slot)
					}),

					// Spacer
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Spacer{
							Width: unit.Dp(20), Height: unit.Dp(1),
						}.Layout(gtx)
					}),

					// Slot 2
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						slot := components.Slot{
							Button: p.SlotButtons[1],
							Number: 1,
							Text:   "description of content...",
						}
						slot.State = components.EmptyUnusable
						if p.Slots[1].HasContent {
							slot.State = components.FilledUsable
						}
						slot.Type = components.SlotText
						if p.Slots[1].Type == vars.TypeImage {
							slot.Type = components.SlotImage
						}
						return components.NewSlot(gtx, theme, slot)
					}),

					// Spacer
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Spacer{
							Width: unit.Dp(20), Height: unit.Dp(1),
						}.Layout(gtx)
					}),

					// Slot 3
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						slot := components.Slot{
							Button: p.SlotButtons[2],
							Number: 2,
							Text:   "description of content...",
						}
						slot.State = components.EmptyUnusable
						if p.Slots[2].HasContent {
							slot.State = components.FilledUsable
						}
						slot.Type = components.SlotText
						if p.Slots[2].Type == vars.TypeImage {
							slot.Type = components.SlotImage
						}
						return components.NewSlot(gtx, theme, slot)
					}),

					// Spacer
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Spacer{
							Width: unit.Dp(20), Height: unit.Dp(1),
						}.Layout(gtx)
					}),

					// Slot 4
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						slot := components.Slot{
							Button: p.SlotButtons[3],
							Number: 3,
							Text:   "description of content...",
						}
						slot.State = components.EmptyUnusable
						if p.Slots[3].HasContent {
							slot.State = components.FilledUsable
						}
						slot.Type = components.SlotText
						if p.Slots[3].Type == vars.TypeImage {
							slot.Type = components.SlotImage
						}
						return components.NewSlot(gtx, theme, slot)
					}),

					// Spacer
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Spacer{
							Width: unit.Dp(20), Height: unit.Dp(1),
						}.Layout(gtx)
					}),

					// Slot 5
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						slot := components.Slot{
							Button: p.SlotButtons[4],
							Number: 4,
							Text:   "description of content...",
						}
						slot.State = components.EmptyUnusable
						if p.Slots[4].HasContent {
							slot.State = components.FilledUsable
						}
						slot.Type = components.SlotText
						if p.Slots[4].Type == vars.TypeImage {
							slot.Type = components.SlotImage
						}
						return components.NewSlot(gtx, theme, slot)
					}),
				)
			})
		}),

		// Spacer
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{
				Width: unit.Dp(gtx.Constraints.Max.X), Height: unit.Dp(85),
			}.Layout(gtx)
		}),

		// Messages
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
					// Message 1
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						msg := components.Message{
							Image: "img4.png",
							Title: "Default slot behaviour:",
							Text:  "If no slot is selected in 5 seconds, content from default will be used",
						}
						return components.NewMessage(gtx, theme, msg)
					}),

					// Spacer
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Spacer{
							Width: unit.Dp(90), Height: unit.Dp(1),
						}.Layout(gtx)
					}),

					// Message 2
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						msg := components.Message{
							Image: "img7.png",
							Title: "Slot selection:",
							Text:  "You can only select any of the non empty slots",
						}
						return components.NewMessage(gtx, theme, msg)
					}),
				)
			})
		}),

		// Spacer
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{
				Width: unit.Dp(gtx.Constraints.Max.X), Height: unit.Dp(98),
			}.Layout(gtx)
		}),

		// Footer
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				f := components.FooterItems{
					SourceButton: p.SourceButton,
					DocsButton:   p.DocsButton,
				}
				return components.NewFooter(gtx, theme, f)
			})
		}),
	)
}
