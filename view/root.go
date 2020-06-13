package view

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

const tag = "view/root"

var app *tview.Application

// Run starts the user interface
func Run() {
	app = tview.NewApplication()
	title := tview.NewTextView()

	title.SetText("\n\n[orange]sunshine live\n\n[blue]electronic music radio").SetTextAlign(tview.AlignCenter).SetDynamicColors(true)
	title.SetInputCapture(handleKeyEvent)

	log.Printf("[%s] starting ui", tag)
	if err := app.SetRoot(title, true).Run(); err != nil {
		log.Fatalf("[%s] ui initialization failed: %d", tag, err)
		fmt.Printf("UI failed, but playing anyways.\nPress CTRL-C to quit")
	}
}

func handleKeyEvent(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyESC:
		log.Printf("[%s] exit due to ESC key", tag)
		app.Stop()
	}
	return event
}

// Stop ends the ui
func Stop() {
	app.Stop()
}
