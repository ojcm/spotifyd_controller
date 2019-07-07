package main

import (
	"net/http"
	"os"
)

var token = "TOKEN1234"

var startURL = "https://api.spotify.com/v1/me/player/play"
var stopURL = "https://api.spotify.com/v1/me/player/pause"
var skipURL = "https://api.spotify.com/v1/me/player/next"
var skipPrevURL = "https://api.spotify.com/v1/me/player/previous"
var client = &http.Client{}

func sendRequest(type string, url string) {
	req, _ := http.NewRequest(type, url, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	client.Do(req)
}

func putRequest(url string) {
	sendRequest("PUT", url)
}

func postRequest(url string) {
	sendRequest("POST", url)
}

// StartPlayback resumes playback of current track
func StartPlayback() {
	putRequest(startURL)
}

// StopPlayback pauses playback of current track
func StopPlayback() {
	putRequest(stopURL)
}

// SkipPlayback skips playback to next track
func SkipPlayback() {
	postRequest(skipURL)
}

// SkipPrevPlayback skips playback to start of current track
func SkipPrevPlayback() {
	postRequest(skipPrevURL)
}

func main() {
	switch action := os.Args[1]; action {
	case "stop":
		StopPlayback()
	case "start":
		StartPlayback()
	case "skip":
		SkipPlayback()
	case "prev":
		SkipPrevPlayback()
	}
}
