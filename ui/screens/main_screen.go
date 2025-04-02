package screens

import (
	"image/color"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
)

func Run_GUI(window *app.Window) error {
	theme := material.NewTheme()
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			if e.Err != nil {
				return e.Err
			}
			os.Exit(0)
		case app.FrameEvent:
			// this graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, e)

			// define a large label with an appropriate text:
			title := material.H1(theme, "Hello World")

			// change the color of the label.
			maroon := color.NRGBA{R: 0, G: 0, B: 0, A: 255}
			title.Color = maroon

			// change the position of the label.
			title.Alignment = text.Middle

			// draw the label to the graphics context.
			layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return title.Layout(gtx)
			})

			// pass the drawing operations to the GPU.
			e.Frame(gtx.Ops)
		}
	}
}
