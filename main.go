package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"jan-sl.de/radigossl/player"
	"jan-sl.de/radigossl/streams"
)

const tag = "main"

// signal when to quit the program
var quitChan = make(chan bool)

func main() {

	// setup logging
	file, err := os.OpenFile("radigossl.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	log.Printf("[%s] Starting application", tag)

	// start SIGTERM handler
	go watchSignal()

	// start the stream
	streams.Get()
	SSLUrl := "http:" + streams.Streams["1"].URLLow
	log.Printf("[%s] playing %s from url %s", tag, streams.Streams["1"].Stream, SSLUrl)

	player.Init()
	player.Play(SSLUrl)

	// setup ui
	app := tview.NewApplication()
	title := tview.NewTextView()

	title.SetText("\n\n[orange]sunshine live\n\n[blue]electronic music radio").SetTextAlign(tview.AlignCenter).SetDynamicColors(true)
	title.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyESC {
			log.Printf("[%s/ui] exit", tag)
			app.Stop()
			close(quitChan)
		}
		return event
	})
	log.Printf("[%s/ui] starting ui", tag)
	if err := app.SetRoot(title, true).Run(); err != nil {
		log.Fatalf("[%s] ui initialization failed: %d", tag, err)
		fmt.Printf("UI failed, but playing anyways.\nPress CTRL-C to quit")
	}

	// now wait until we quit
	<-quitChan
	log.Printf("[%s] Cleaning up", tag)
	player.Stop()
	player.Close()
	log.Printf("[%s] End", tag)
}

// SIGTERM handler to gracefully end
func watchSignal() {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-signalChan
	log.Printf("[%s] received interrupt", tag)
	close(quitChan)
}
