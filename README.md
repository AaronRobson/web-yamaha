# web-yamaha
[![Build Status](https://travis-ci.org/AaronRobson/web-yamaha.svg?branch=master)](https://travis-ci.org/AaronRobson/web-yamaha)

A web based app for controlling a Yamaha amplifier using the Extended Control API.

## Prerequisites
```bash
sudo apt update
sudo apt install golang
go get -v github.com/gorilla/mux
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
