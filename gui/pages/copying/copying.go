package copying

import (
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/leojimenezg/scapmi/gui/components"
	"github.com/leojimenezg/scapmi/gui/widgets"
	"github.com/leojimenezg/scapmi/internal/manager"
	"github.com/leojimenezg/scapmi/internal/vars"
)

// CopyingItems contains all pointers to its interactive items.
type CopyingItems struct {
	Manager      *manager.Manager
	SlotButtons  [5]*widget.Clickable
	SourceButton *widget.Clickable
	DocsButton   *widget.Clickable
}

func (c *CopyingItems) Draw(gtx layout.Context, theme *material.Theme) {
	layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
		// Header
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return components.NewHeader(gtx, theme, "copying")
		}),

		// Title
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return widgets.NewTitle(gtx, theme, "Available slots")
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
				gtx, theme, "Select a slot to save recently copied content in it", text.Middle)
		}),

		// Spacer
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{
				Width: unit.Dp(gtx.Constraints.Max.X), Height: unit.Dp(60),
			}.Layout(gtx)
		}),

		// Slots
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
					// Slot 1
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						slot := components.Slot{
							Button: c.SlotButtons[0],
							Number: 0,
							Text:   c.Manager.Slots[0].Summary,
							Extra:  "Default",
						}
						slot.Action = func() { c.Manager.SaveToSlot(slot.Number) }
						slot.State = components.EmptyUsable
						if c.Manager.Slots[0].HasContent {
							slot.State = components.FilledUsable
						}
						slot.Type = components.SlotText
						if c.Manager.Slots[0].Type == vars.TypeImage {
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
							Button: c.SlotButtons[1],
							Number: 1,
							Text:   c.Manager.Slots[1].Summary,
						}
						slot.Action = func() { c.Manager.SaveToSlot(slot.Number) }
						slot.State = components.EmptyUsable
						if c.Manager.Slots[1].HasContent {
							slot.State = components.FilledUsable
						}
						slot.Type = components.SlotText
						if c.Manager.Slots[1].Type == vars.TypeImage {
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
							Button: c.SlotButtons[2],
							Number: 2,
							Text:   c.Manager.Slots[2].Summary,
						}
						slot.Action = func() { c.Manager.SaveToSlot(slot.Number) }
						slot.State = components.EmptyUsable
						if c.Manager.Slots[2].HasContent {
							slot.State = components.FilledUsable
						}
						slot.Type = components.SlotText
						if c.Manager.Slots[2].Type == vars.TypeImage {
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
							Button: c.SlotButtons[3],
							Number: 3,
							Text:   c.Manager.Slots[3].Summary,
						}
						slot.Action = func() { c.Manager.SaveToSlot(slot.Number) }
						slot.State = components.EmptyUsable
						if c.Manager.Slots[3].HasContent {
							slot.State = components.FilledUsable
						}
						slot.Type = components.SlotText
						if c.Manager.Slots[3].Type == vars.TypeImage {
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
							Button: c.SlotButtons[4],
							Number: 4,
							Text:   c.Manager.Slots[4].Summary,
						}
						slot.Action = func() { c.Manager.SaveToSlot(slot.Number) }
						slot.State = components.EmptyUsable
						if c.Manager.Slots[4].HasContent {
							slot.State = components.FilledUsable
						}
						slot.Type = components.SlotText
						if c.Manager.Slots[4].Type == vars.TypeImage {
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
							Image: "clock.png",
							Title: "Default slot behaviour:",
							Text:  "If no slot is selected in 5 seconds, the content goes to default slot",
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
							Image: "save.png",
							Title: "Saving to empty slot:",
							Text:  "If the selected slot is empty, copied content saves in it",
						}
						return components.NewMessage(gtx, theme, msg)
					}),

					// Spacer
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Spacer{
							Width: unit.Dp(90), Height: unit.Dp(1),
						}.Layout(gtx)
					}),

					// Message 3
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						msg := components.Message{
							Image: "replace.png",
							Title: "Replacing content:",
							Text:  "If the selected is not empty, its content wil be replaced",
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
					SourceButton: c.SourceButton,
					DocsButton:   c.DocsButton,
				}
				return components.NewFooter(gtx, theme, f)
			})
		}),
	)
}
