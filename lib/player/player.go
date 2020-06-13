package player

import (
	"log"

	vlc "github.com/adrg/libvlc-go/v3"
)

const tag = "player"

var streamPlayer *vlc.Player

// Init sets up the media player
func Init() {
	log.Printf("[%s] init", tag)

	if err := vlc.Init("--no-video", "--quiet"); err != nil {
		log.Fatalf("[%s] %d", tag, err)
	}

	// Create a new player.
	var err error
	streamPlayer, err = vlc.NewPlayer()
	if err != nil {
		log.Fatalf("[%s] %d", tag, err)
	}
}

// Close releases the libvlc instance
func Close() {
	log.Printf("[%s] closing", tag)
	streamPlayer.Release()
	vlc.Release()
}

// Play plays the stream
func Play(streamURL string) {
	log.Printf("[%s] start playback", tag)

	// Add a media file from path or from URL.
	// Set player media from path:
	// media, err := player.LoadMediaFromPath("localpath/test.mp4")
	// Set player media from URL:
	var err error
	_, err = streamPlayer.LoadMediaFromURL(streamURL)
	if err != nil {
		log.Fatalf("[%s] %d", tag, err)
	}

	// Start playing the media.
	if err = streamPlayer.Play(); err != nil {
		log.Fatalf("[%s] %d", tag, err)
	}
}

// Stop stops the playback
func Stop() {
	streamPlayer.Stop()
	media, _ := streamPlayer.Media()
	media.Release()
}
