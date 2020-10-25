# web-yamaha
[![Build Status](https://travis-ci.org/AaronRobson/web-yamaha.svg?branch=master)](https://travis-ci.org/AaronRobson/web-yamaha)

A web based app for controlling a Yamaha amplifier using the Extended Control API.

## Prerequisites
```bash
sudo apt update
sudo apt install golang
go get -v github.com/gorilla/mux
go get -v gopkg.in/go-playground/validator.v9
```

## Format and Test
```bash
make
```

## Run
```bash
make run
```
Open http://localhost:8080/

You may call it programmatically, for example:
```bash
curl -i http://localhost:8080/ping
```
Which ought to return the JSON `true`.

Or more usefully:
```bash
curl -i -X POST --data '{"mute": "mute"}' --header 'Content-Type: application/json' http://localhost:8080/\$setMute
```
Or:
```bash
curl -i -X POST --data '{"volume": "down"}' --header 'Content-Type: application/json' http://localhost:8080/\$setVolume
```

