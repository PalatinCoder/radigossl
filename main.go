package main

import (
	"log"
	"os"

	"jan-sl.de/radigossl/lib/player"
	"jan-sl.de/radigossl/lib/streams"
	"jan-sl.de/radigossl/view"
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

	// start the stream
	streams.RetrieveStreams()

	player.Init()
	defer player.Release()

	// run the view
	view.Run() // this call blocks while the view (i.e. the app) is running

	log.Printf("[%s] Cleaning up", tag)
}
