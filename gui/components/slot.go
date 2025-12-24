package components

import (
	"log"
	"strconv"

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
	SlotEmpty SlotState = iota
	SlotUsed
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
	if slot.Button.Clicked(gtx) {
		// TODO: Use actual function to handle slot
		log.Printf("Slot %d clicked", slot.Number)
	}
	c := colors.ColorBackgroundLight
	if slot.Button.Hovered() {
		c = colors.ColorBackgroundHover
	}
	return slot.Button.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		card := widgets.Card{
			Width:       215,
			Height:      280,
			Radius:      10,
			StrokeWidth: 2,
			StrokeColor: colors.ColorBackgroundHover,
			Color:       c,
		}
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
									color := colors.ColorTextSub
									if slot.State == SlotUsed {
										color = colors.ColorActionMain
									}
									circle := widgets.Circle{
										Diameter:    20,
										Color:       color,
										StrokeColor: colors.ColorBackgroundHover,
										StrokeWidth: 2,
									}
									return widgets.NewCircle(gtx, circle)
								})
							}),
						)
					}),

					// Spacer or "EMPTY" text
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						if slot.State == SlotEmpty {
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
						if slot.State == SlotEmpty {
							return layout.Dimensions{}
						}
						return widgets.NewText(gtx, theme, string(slot.Type), text.Start)
					}),

					// Spacer or Nothing
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						if slot.State == SlotEmpty {
							return layout.Dimensions{}
						}
						return layout.Spacer{
							Width: unit.Dp(gtx.Constraints.Max.X), Height: unit.Dp(5),
						}.Layout(gtx)
					}),

					// Text or Nothing
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						if slot.State == SlotEmpty {
							return layout.Dimensions{}
						}
						return widgets.NewSubText(gtx, theme, string(slot.Text), text.Start)
					}),
				)
			})
		})
	})
}
