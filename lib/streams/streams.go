package streams

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const tag = "streams"

type stream struct {
	StationID  int    `json:"station_id"`
	Stream     string `json:"stream"`
	StreamLogo string `json:"stream_logo"`
	URLLow     string `json:"url_low"`
	URLHigh    string `json:"url_high"`
}

var Streams map[string]stream

// Get collects the stream infos from the station's api
func Get() {
	log.Printf("[%s] get streams", tag)

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := httpClient.Get("https://sunshinelive-stream-service.loverad.io/v1/live")
	switch {
	case err != nil:
		log.Fatalf("[%s] %d", tag, err)
	case resp.StatusCode != 200:
		log.Fatalf("[%s] stream api returned %s", tag, resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("[%s] Error reading response: %d", tag, err)
	}

	err = json.Unmarshal(body, &Streams)
	if err != nil {
		log.Fatalf("[%s] could not decode response: %d", tag, err)
	}
}
