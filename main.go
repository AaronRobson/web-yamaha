package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Print("Starting")

	shutdown := make(chan bool)

	go func() {
		if err := http.ListenAndServe(":80", MuteRouter()); err != nil {
			log.Print(err)
		}
		shutdown <- true
	}()

	<-shutdown
}

func MuteRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", indexHandler).Methods("GET")
	router.HandleFunc("/mute", muteHandler).Methods("GET")
	router.HandleFunc("/unmute", unmuteHandler).Methods("GET")

	return router
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("index")
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(fmt.Sprintf(
		`<!DOCTYPE html>
<html lang="en-US">
<head>
<title>Hi-Fi Control</title>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
</head>
<body>
<script>
function makeCall(endpointName) {
  console.log(endpointName);
  var request = new XMLHttpRequest()
  request.open('GET', endpointName, true)
  request.onreadystatechange = function () {
    if (request.readyState !== 4) {
      return;
    }
    if (request.status === 200) {
      var jsonResponse = JSON.parse(request.responseText);
      var response_code = jsonResponse['response_code'];
      if (response_code === 0) {
        console.log('done');
      } else {
        console.log('non-zero response code found ' + response_code);
      }
    } else {
      console.log('error: ' + request.status + ' ' + request.statusText +
        '\n' + request.responseText);
    }
  }
  request.send();
}
function mute() {
  makeCall('/mute')
}
function unmute() {
  makeCall('/unmute')
}
</script>
<h1>Hi-Fi Control</h1>
<p title="Go ahead">Control your Hi-Fi with the buttons below.</p>
<button onclick="mute()">Mute</button>
<button onclick="unmute()">Unmute</button>
</body>
</html>`)))
}

const (
	hifiUrl    = "http://192.168.0.99/YamahaExtendedControl/v1"
	setMuteUrl = hifiUrl + "/main/setMute"
	muteUrl    = setMuteUrl + "?enable=true"
	unmuteUrl  = setMuteUrl + "?enable=false"
)

// Timeouts: https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
var netClient = &http.Client{
	Timeout: time.Second * 10,
}

type Error struct {
	Message string
}

func muteHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("mute")
	w.Header().Add("Content-Type", "application/json")
	resp, err := netClient.Get(muteUrl)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		error := Error{Message: "failed to connect"}
		json.NewEncoder(w).Encode(error)
		return
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		error := Error{Message: "failed to accept body"}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(resp.StatusCode)
	w.Write(content)
}

func unmuteHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("unmute")
	w.Header().Add("Content-Type", "application/json")
	resp, err := netClient.Get(unmuteUrl)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		error := Error{Message: "failed to connect"}
		json.NewEncoder(w).Encode(error)
		return
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		error := Error{Message: "failed to accept body"}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(resp.StatusCode)
	w.Write(content)
}
