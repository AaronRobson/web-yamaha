package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strconv"
	"time"
)

func main() {
	var port = 8080
	log.Print("Starting - http://localhost:" + strconv.Itoa(port) + "/")

	shutdown := make(chan bool)

	go func() {
		if err := http.ListenAndServe(":"+strconv.Itoa(port), MuteRouter()); err != nil {
			log.Print(err)
		}
		shutdown <- true
	}()

	<-shutdown
}

func MuteRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", indexHandler).Methods("GET")
	router.HandleFunc("/ping", pingHandler).Methods("GET")
	router.HandleFunc("/mute", muteHandler).Methods("GET")
	router.HandleFunc("/unmute", unmuteHandler).Methods("GET")

	return router
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("index")
	p := path.Dir("./static/index.html")
	http.ServeFile(w, r, p)
}

// Timeouts: https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
var netClient = &http.Client{
	Timeout: time.Second * 10,
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ping")
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(`true`))
}

type ErrorResponse struct {
	Message string
}

func muteHandler(w http.ResponseWriter, r *http.Request) {
	generalMuteHandler(true, w, r)
}

func unmuteHandler(w http.ResponseWriter, r *http.Request) {
	generalMuteHandler(false, w, r)
}

func generalMuteHandler(mute bool, w http.ResponseWriter, r *http.Request) {
	log.Print(muteToStr(mute))
	w.Header().Add("Content-Type", "application/json")
	resp, err := netClient.Get(findMuteUrl(mute))
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		errorResponse := ErrorResponse{Message: "failed to connect"}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		errorResponse := ErrorResponse{Message: "failed to accept body"}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}
	w.WriteHeader(resp.StatusCode)
	w.Write(content)
}
