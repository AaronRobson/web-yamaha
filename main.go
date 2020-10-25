package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strconv"
	"time"
)

var validate *validator.Validate

func ValidateMute(fl validator.FieldLevel) bool {
	switch fl.Field().String() {
	case
		"mute",
		"unmute":
		return true
	}
	return false
}

func ValidateVolume(fl validator.FieldLevel) bool {
	switch fl.Field().String() {
	case
		"up",
		"down":
		return true
	}
	return false
}

func main() {
	validate = validator.New()
	validate.RegisterValidation("mute", ValidateMute)
	validate.RegisterValidation("volume", ValidateVolume)

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

type ErrorResponse struct {
	Message string `json:"message"`
}

type MuteRequest struct {
	Mute string `json:"mute" validate:"mute"`
}

type VolumeRequest struct {
	Volume string `json:"volume" validate:"volume"`
}

func setMuteHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("/$setMute")
	w.Header().Add("Content-Type", "application/json")
	var request MuteRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		errorResponse := ErrorResponse{Message: err.Error()}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	err = validate.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := ErrorResponse{Message: err.Error()}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	muteBool := request.Mute == "mute"

	resp, err := netClient.Get(findMuteUrl(muteBool))
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

func setVolumeHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("/$setVolume")
	w.Header().Add("Content-Type", "application/json")
	var request VolumeRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		errorResponse := ErrorResponse{Message: err.Error()}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	err = validate.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := ErrorResponse{Message: err.Error()}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	volumeBool := request.Volume == "up"

	resp, err := netClient.Get(findVolumeUrl(volumeBool))
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
