package view

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell"
	"gitlab.com/tslocum/cview"
)

const tag = "view/root"

var app = cview.NewApplication()

// Run starts the user interface
func Run() {
	log.Printf("[%s] starting ui", tag)

	// main layout
	view := cview.NewTextView().
		SetDynamicColors(true).
		SetText("\n[orange]sunshine live\n\n[blue]electronic music radio").
		SetTextAlign(cview.AlignCenter)
	view.
		SetBorder(true).
		SetTitle("radigossl")

	app.SetInputCapture(handleKeyEvent)
	if err := app.SetRoot(view, true).Run(); err != nil {
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
