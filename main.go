package main

import (
	"log"
	"sentrypass/logic/database"
	/*
		"os"
		"image/color"
		"gioui.org/app"
		"gioui.org/layout"
		"gioui.org/op"
		"gioui.org/text"
		"gioui.org/widget/material"
	*/)

/*
func run_GUI(window *app.Window) error {
	theme := material.NewTheme()
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			// this graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, e)

			// define an large label with an appropriate text:
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

*/

func main() {
	// Inicializa o banco de dados
	db := database.InitDB()
	defer db.Close()

	// Executa as migrações
	database.RunMigrations(db)

	log.Println("Aplicação iniciada com sucesso.")
	/*
		go func() {
			window := new(app.Window)
			err := run_GUI(window)
			if err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
		}()
		app.Main()
	*/
}
