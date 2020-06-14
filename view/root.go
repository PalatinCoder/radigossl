package view

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell"
	"gitlab.com/tslocum/cview"
	"jan-sl.de/radigossl/lib/streams"
)

const tag = "view/root"

var app = cview.NewApplication()

// Run starts the user interface
func Run() {
	log.Printf("[%s] starting ui", tag)

	// stream list
	streamList := cview.NewList()
	streamList.
		Clear().
		SetHighlightFullLine(true).
		ShowSecondaryText(false).
		SetBorder(true).
		SetTitle("📻 streams").
		SetTitleAlign(cview.AlignLeft)

	for _, stationID := range streams.StreamStationIDs {
		streamList.AddItem(streams.Streams[stationID].Stream, "", 0, func() { log.Printf("[%s] selected station id %d", tag, stationID) })
	}

	// now playing
	nowplaying := cview.NewTextView()
	nowplaying.
		SetText("...").
		SetTextAlign(cview.AlignCenter).
		SetTitle("🎵 now playing").
		SetTitleAlign(cview.AlignLeft).
		SetBorder(true)

	// player controls
	controls := cview.NewGrid()
	controls.
		SetSize(1, 4, 0, 0).
		AddItem(cview.NewButton("Play/Pause"), 0, 0, 1, 1, 0, 0, false).
		AddItem(cview.NewButton("Quit"), 0, 3, 1, 1, 0, 0, false).
		SetBorders(true).
		SetTitle("🎛️ controls").
		SetTitleAlign(cview.AlignLeft).
		SetBorder(true)

	// main layout
	flex := cview.NewFlex().
		SetDirection(cview.FlexRow).
		AddItem(streamList, 0, 1, true).
		AddItem(nowplaying, 3, 1, false).
		AddItem(controls, 5, 1, false)

	app.SetInputCapture(handleKeyEvent)
	if err := app.SetRoot(flex, true).Run(); err != nil {
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
