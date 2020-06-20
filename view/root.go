package view

import (
	"log"

	"github.com/gdamore/tcell"
	"gitlab.com/tslocum/cview"
	"jan-sl.de/radigossl/lib/player"
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
		streamList.AddItem(streams.Streams[stationID].Stream, "", 0, nil)
	}

	// now playing
	nowplaying := cview.NewTextView()
	nowplaying.
		SetText("...").
		SetTextAlign(cview.AlignCenter).
		SetTitle("🎵 now playing").
		SetTitleAlign(cview.AlignLeft).
		SetBorder(true)

	streamList.SetSelectedFunc(func(idx int, maintext string, secondarytext string, shortcut rune) {
		log.Printf("[%s] selected stream changed", tag)

		stationID := streams.StreamStationIDs[idx]
		url := "http:" + streams.Streams[stationID].URLHigh
		nowplaying.SetText(streams.Streams[stationID].Stream)

		log.Printf("[%s] playing %s from url %s", tag, streams.Streams[stationID].Stream, url)
		player.Stop()
		player.Play(url)
	})

	// main layout
	flex := cview.NewFlex().
		SetDirection(cview.FlexRow).
		AddItem(streamList, 0, 1, true).
		AddItem(nowplaying, 3, 1, false)

	app.SetInputCapture(handleKeyEvent)
	if err := app.SetRoot(flex, true).Run(); err != nil {
		log.Fatalf("[%s] ui initialization failed: %d", tag, err)
	}
}

func handleKeyEvent(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyESC:
		log.Printf("[%s] exit due to ESC key", tag)
		app.Stop()
	case tcell.KeyRune:
		switch event.Rune() {
		case 'p', 's':
			log.Printf("[%s] pausing playback", tag)
			player.Stop()
		case 'q':
			log.Printf("[%s] exit due to q key", tag)
			app.Stop()
		}
	}
	return event
}
