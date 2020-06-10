package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func main() {
	file, err := os.OpenFile("radiossl.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	log.Print("Starting application")

	app := tview.NewApplication()
	title := tview.NewTextView()

	title.SetText("\n\n[orange]sunshine live\n\n[blue]electronic music radio").SetTextAlign(tview.AlignCenter).SetDynamicColors(true)
	title.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyESC {
			app.Stop()
		}
		return event
	})
	if err := app.SetRoot(title, true).Run(); err != nil {
		panic(err)
	}

	log.Print("Ending")
}
