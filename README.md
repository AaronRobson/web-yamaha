# web-yamaha
[![Build Status](https://travis-ci.org/AaronRobson/web-yamaha.svg?branch=master)](https://travis-ci.org/AaronRobson/web-yamaha)

A web based app for controlling a Yamaha amplifier using the Extended Control API.

## Prerequisites
```bash
sudo apt update
sudo apt install golang golint
make install
```

## Format code
Recommend this be run before committing:
```bash
make format
```

## Check and Test
```bash
make
```

### Check
Run `golint`.
```bash
make check
```

### Test
```bash
make test
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
