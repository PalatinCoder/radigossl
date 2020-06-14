package streams

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
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

// Streams is a map of available streams, where the key is the station ID of the stream
var Streams map[int]stream

// StreamStationIDs holds all available stations IDs in ascending order
var StreamStationIDs []int

// RetrieveStreams collects the stream infos from the station's api
func RetrieveStreams() {
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

	var rawList map[string]stream
	err = json.Unmarshal(body, &rawList)
	if err != nil {
		log.Fatalf("[%s] could not decode response: %d", tag, err)
	}

	// bring the list in final order, to make access easier later on
	Streams = make(map[int]stream, len(rawList))
	for _, stream := range rawList {
		Streams[stream.StationID] = stream
		StreamStationIDs = append(StreamStationIDs, stream.StationID)
	}
	sort.Ints(StreamStationIDs)
}
