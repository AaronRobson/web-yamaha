package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	var port = 8080
	log.Print("Starting - http://localhost:" + strconv.Itoa(port) + "/")

	shutdown := make(chan bool)

	go func() {
		if err := http.ListenAndServe(":"+strconv.Itoa(port), muteRouter()); err != nil {
			log.Print(err)
		}
		shutdown <- true
	}()

	<-shutdown
}

func muteRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", indexHandler).Methods("GET")
	router.HandleFunc("/favicon.png", faviconHandler).Methods("GET")

	router.HandleFunc("/ping", pingHandler).Methods("GET")

	router.HandleFunc("/$setMute", setMuteHandler).Methods("POST")
	router.HandleFunc("/$setVolume", setVolumeHandler).Methods("POST")

	return router
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("index")
	p := path.Dir("./static/index.html")
	http.ServeFile(w, r, p)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("favicon")
	http.ServeFile(w, r, "./static/favicon.png")
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

func setMuteHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("/$setMute")
	w.Header().Add("Content-Type", "application/json")
	var request muteRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		errorResponse := errorResponse{Message: err.Error()}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	err = validate.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := errorResponse{Message: err.Error()}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	muteBool := request.Mute == "mute"

	resp, err := netClient.Get(findMuteURL(muteBool))
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		errorResponse := errorResponse{Message: "failed to connect"}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		errorResponse := errorResponse{Message: "failed to accept body"}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}
	w.WriteHeader(resp.StatusCode)
	w.Write(content)
}

func setVolumeHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("/$setVolume")
	w.Header().Add("Content-Type", "application/json")
	var request volumeRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		errorResponse := errorResponse{Message: err.Error()}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	err = validate.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := errorResponse{Message: err.Error()}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	volumeBool := request.Volume == "up"

	resp, err := netClient.Get(findVolumeURL(volumeBool))
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		errorResponse := errorResponse{Message: "failed to connect"}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		errorResponse := errorResponse{Message: "failed to accept body"}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}
	w.WriteHeader(resp.StatusCode)
	w.Write(content)
}
