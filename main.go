package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

const tag = "main"

func main() {

	// setup logging
	file, err := os.OpenFile("radigossl.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	log.Printf("[%s] Starting application", tag)

	// setup ui
	app := tview.NewApplication()
	title := tview.NewTextView()

	title.SetText("\n\n[orange]sunshine live\n\n[blue]electronic music radio").SetTextAlign(tview.AlignCenter).SetDynamicColors(true)
	title.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyESC {
			log.Printf("[%s/ui] exit", tag)
			app.Stop()
		}
		return event
	})
	log.Printf("[%s/ui] starting ui", tag)
	if err := app.SetRoot(title, true).Run(); err != nil {
		log.Fatalf("[%s] ui initialization failed: %d", tag, err)
	log.Printf("[%s] End", tag)
	}

	log.Print("Ending")
}
