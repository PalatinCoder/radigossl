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

	playerEventManager, err = streamPlayer.EventManager()
	if err != nil {
		log.Printf("[%s] could not initialize event manager: %d", tag, err)
	}

	for _, event := range playerEvents {
		eventID, err := playerEventManager.Attach(event, handlePlayerEvents, nil)
		if err != nil {
			log.Printf("[%s/eventmanager] %d", tag, err)
		}
		playerEventIDs = append(playerEventIDs, eventID)
	}
}

// Release cleans up the libvlc instance
func Release() {
	log.Printf("[%s] closing", tag)

	for _, eventID := range playerEventIDs {
		playerEventManager.Detach(eventID)
	}

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

var playerEventManager *vlc.EventManager
var playerEventIDs []vlc.EventID

var playerEvents = []vlc.Event{
	vlc.MediaPlayerBuffering,
}

func handlePlayerEvents(event vlc.Event, userData interface{}) {
	switch event {
	case vlc.MediaPlayerBuffering:
		log.Printf("[%s/eventhandler] buffering", tag)
	}
}
