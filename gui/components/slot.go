package components

import (
	"log"
	"strconv"

	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/leojimenezg/scapmi/gui/colors"
	"github.com/leojimenezg/scapmi/gui/widgets"
)

type SlotState int

type SlotType string

const (
	EmptyUsable = iota
	EmptyUnusable
	FilledUsable
)

const (
	SlotText  SlotType = "TEXT"
	SlotImage SlotType = "IMAGE"
)

type Slot struct {
	Button *widget.Clickable
	Number int
	State  SlotState
	Type   SlotType
	Text   string
	Extra  string
}

func NewSlot(gtx layout.Context, theme *material.Theme, slot Slot) layout.Dimensions {
	card := widgets.Card{Width: 215, Height: 280, Radius: 10, StrokeWidth: 2}
	if slot.State != EmptyUnusable {
		card.Color = colors.ColorBackgroundLight
		card.StrokeColor = colors.ColorBackgroundHover
		if slot.Button.Clicked(gtx) {
			// TODO: Use actual function to handle slot
			log.Printf("Slot %d clicked", slot.Number)
		}
		if slot.Button.Hovered() {
			card.Color = colors.ColorBackgroundHover
			pointer.CursorPointer.Add(gtx.Ops)
		}
		return slot.Button.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return drawSlot(gtx, theme, slot, card)
		})
	} else {
		card.Color = colors.ColorTextSub
		card.StrokeColor = colors.ColorTextMain
		return drawSlot(gtx, theme, slot, card)
	}
}

func drawSlot(gtx layout.Context, theme *material.Theme, slot Slot, card widgets.Card) layout.Dimensions {
	return widgets.NewCard(gtx, card, func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle, WeightSum: 1}.Layout(gtx,
				// Number and status
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle, WeightSum: 3}.Layout(gtx,
						// Number
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							num := strconv.Itoa(slot.Number + 1)
							return widgets.NewText(gtx, theme, num, text.Start)
						}),

						// Extra
						layout.Flexed(0.9, func(gtx layout.Context) layout.Dimensions {
							return widgets.NewSubText(gtx, theme, slot.Extra, text.Middle)
						}),

						// Status
						layout.Flexed(1.1, func(gtx layout.Context) layout.Dimensions {
							return layout.E.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								circle := widgets.Circle{Diameter: 20, StrokeWidth: 2}
								switch slot.State {
								case EmptyUsable:
									circle.Color = colors.ColorTextSub
									circle.StrokeColor = colors.ColorBackgroundHover
								case EmptyUnusable:
									circle.Color = colors.ColorTextSub
									circle.StrokeColor = colors.ColorTextMain
								case FilledUsable:
									circle.Color = colors.ColorActionMain
									circle.StrokeColor = colors.ColorBackgroundHover
								}
								return widgets.NewCircle(gtx, circle)
							})
						}),
					)
				}),

				// Spacer or "EMPTY" text
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					if slot.State != FilledUsable {
						return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return widgets.NewText(gtx, theme, "EMPTY", text.Middle)
						})
					}
					return layout.Spacer{
						Width: unit.Dp(gtx.Constraints.Max.X), Height: unit.Dp(10),
					}.Layout(gtx)
				}),

				// Type or Nothing
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if slot.State != FilledUsable {
						return layout.Dimensions{}
					}
					return widgets.NewText(gtx, theme, string(slot.Type), text.Start)
				}),

				// Spacer or Nothing
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if slot.State != FilledUsable {
						return layout.Dimensions{}
					}
					return layout.Spacer{
						Width: unit.Dp(gtx.Constraints.Max.X), Height: unit.Dp(5),
					}.Layout(gtx)
				}),

				// Text or Nothing
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if slot.State != FilledUsable {
						return layout.Dimensions{}
					}
					return widgets.NewSubText(gtx, theme, string(slot.Text), text.Start)
				}),
			)
		})
	})
}
