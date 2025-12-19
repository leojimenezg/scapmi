package main

import (
	"image"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/leojimenezg/scapmi/gui/colors"
	"github.com/leojimenezg/scapmi/gui/interfaces"
	"github.com/leojimenezg/scapmi/gui/widgets"
	"github.com/leojimenezg/scapmi/internal/listener"
	"github.com/leojimenezg/scapmi/internal/vars"
	"github.com/leojimenezg/scapmi/internal/watcher"
)

var CurrentAppState vars.AppState
var Window *app.Window

func main() {
	err := watcher.InitClipboard()
	if err != nil {
		log.Fatal(err.Error())
	}
	Window = interfaces.NewWindow()
	go watcher.WatchClipboard(&CurrentAppState, Window)
	go listener.SetHooks(&CurrentAppState, Window)
	go ShowGUI()
	app.Main()
}

func ShowGUI() {
	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
	for {
		switch e := Window.Event().(type) {
		case app.DestroyEvent:
			os.Exit(0)
		case app.FrameEvent:
			var ops op.Ops
			gtx := app.NewContext(&ops, e)
			paint.ColorOp{Color: colors.ColorBackground}.Add(gtx.Ops)
			paint.PaintOp{}.Add(gtx.Ops)

			switch CurrentAppState {
			case vars.StateInit:
				layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						inset := layout.Inset{Top: unit.Dp(30), Bottom: unit.Dp(10)}
						return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return widgets.NewImage(gtx, "logo.png", layout.Center, 0.2)
						})
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return widgets.NewMainTitle(gtx, th, "Welcome to scapmi")
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return widgets.NewSubText(gtx, th,
							"Simultaneously Copy And Paste Multiple Items", text.Middle)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						inset := layout.Inset{Top: unit.Dp(50), Bottom: unit.Dp(10)}
						return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return widgets.NewTitle(gtx, th, "App shortcuts")
						})
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						card := widgets.Card{
							Size:         image.Pt(1000, 380),
							CornerRadius: 15,
							StokeWidth:   2,
							StokeColor:   colors.ColorBackgroundHover,
							Color:        colors.ColorBackgroundLight,
						}
						return widgets.NewCard(gtx, card, func(gtx layout.Context) layout.Dimensions {
							return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle, WeightSum: 1}.Layout(gtx,
											layout.Rigid(func(gtx layout.Context) layout.Dimensions {
												gtx.Constraints.Min.X = 100
												gtx.Constraints.Min.Y = 100
												gtx.Constraints.Max.X = 100
												gtx.Constraints.Max.Y = 100
												return widgets.NewImage(gtx, "img1.png", layout.Center, 0.04)
											}),
											layout.Rigid(func(gtx layout.Context) layout.Dimensions {
												return widgets.NewText(gtx, th, "Select slot", text.Middle)
											}),
											layout.Flexed(.93, func(gtx layout.Context) layout.Dimensions {
												return widgets.NewText(gtx, th, "cmd + c", text.End)
											}),
										)
									}),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return widgets.NewDivider(gtx, 2, colors.ColorBackgroundHover)
									}),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle, WeightSum: 1}.Layout(gtx,
											layout.Rigid(func(gtx layout.Context) layout.Dimensions {
												gtx.Constraints.Min.X = 100
												gtx.Constraints.Min.Y = 100
												gtx.Constraints.Max.X = 100
												gtx.Constraints.Max.Y = 100
												return widgets.NewImage(gtx, "img2.png", layout.Center, 0.04)
											}),
											layout.Rigid(func(gtx layout.Context) layout.Dimensions {
												return widgets.NewText(gtx, th, "Select content", text.Middle)
											}),
											layout.Flexed(.93, func(gtx layout.Context) layout.Dimensions {
												return widgets.NewText(gtx, th, "ctrl + alt + v", text.End)
											}),
										)
									}),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										extra := image.Pt(gtx.Constraints.Max.X, gtx.Dp(4))
										return layout.Dimensions{Size: extra}
									}),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return widgets.NewDivider(gtx, 2, colors.ColorBackgroundHover)
									}),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle, WeightSum: 1}.Layout(gtx,
											layout.Rigid(func(gtx layout.Context) layout.Dimensions {
												gtx.Constraints.Min.X = 100
												gtx.Constraints.Min.Y = 100
												gtx.Constraints.Max.X = 100
												gtx.Constraints.Max.Y = 100
												return widgets.NewImage(gtx, "img3.png", layout.Center, 0.04)
											}),
											layout.Rigid(func(gtx layout.Context) layout.Dimensions {
												return widgets.NewText(gtx, th, "Exit application", text.Middle)
											}),
											layout.Flexed(.93, func(gtx layout.Context) layout.Dimensions {
												return widgets.NewText(gtx, th, "ctrl + alt + q", text.End)
											}),
										)
									}),
								)
							})
						})
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Spacer{Width: unit.Dp(10), Height: unit.Dp(210)}.Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return widgets.NewSmallText(gtx, th, "v1.0.0")
								}),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return layout.Dimensions{Size: image.Pt(gtx.Dp(10), gtx.Dp(10))}
								}),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return widgets.NewSmallText(gtx, th, "Source")
								}),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return layout.Dimensions{Size: image.Pt(gtx.Dp(10), gtx.Dp(10))}
								}),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return widgets.NewSmallText(gtx, th, "Docs")
								}),
							)
						})
					}),
				)

			case vars.StateIdle:
				material.H1(th, "Idle State").Layout(gtx)

			case vars.StateCopying:
				material.H1(th, "Copying State").Layout(gtx)

			case vars.StatePasting:
				material.H1(th, "Pasting State").Layout(gtx)
			}

			e.Frame(gtx.Ops)
		}
	}
}
