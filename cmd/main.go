package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/mfbarnable/wvatha_docker/gui"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	uiElements := make([]ui.Drawable, 0)
	uiElements = append(uiElements, gui.GetHomePage())
	ui.Render(uiElements...)
	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q":
			return
		case "<Up>":
			list := uiElements[0]
			var s = *widgets.List(list)
		}

		ui.Render(uiElements...)

	}

}
