package main

import (
	"net/http"
)

var token = "BQAcNU-YXnL0GjOLfDH10JwPqsT5Xe6Vi1hCKuqgy2tEyWDsBgcw4poPOzHUIxFxhGjJ9aoCWYlyBpsENilTBXrZ91rI-3qitonfP9Qy8XdMgZSbkiQxZHVRiObZCTKZmvI8rDfXOMpge9UpATdfNAbn1_hdDzg"
var startURL = "https://api.spotify.com/v1/me/player/play"
var stopURL = "https://api.spotify.com/v1/me/player/pause"
var skipURL = "https://api.spotify.com/v1/me/player/next"
var skipPrevURL = "https://api.spotify.com/v1/me/player/previous"
var client = &http.Client{}

func addHeaders(req *http.Request) {
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
}

func putRequest(url string) {
	req, _ := http.NewRequest("PUT", url, nil)
	addHeaders(req)
	client.Do(req)
}

func postRequest(url string) {
	req, _ := http.NewRequest("POST", url, nil)
	addHeaders(req)
	client.Do(req)
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

// SkipPrevPlayback skips playback to start of currentr track
func SkipPrevPlayback() {
	postRequest(skipPrevURL)
}
