package main

import "gopkg.in/go-playground/validator.v10"

type errorResponse struct {
	Message string `json:"message"`
}

type muteRequest struct {
	Mute string `json:"mute" validate:"mute"`
}

type volumeRequest struct {
	Volume string `json:"volume" validate:"volume"`
}

var validate *validator.Validate

func validateMute(fl validator.FieldLevel) bool {
	switch fl.Field().String() {
	case
		"mute",
		"unmute":
		return true
	}
	return false
}

func validateVolume(fl validator.FieldLevel) bool {
	switch fl.Field().String() {
	case
		"up",
		"down":
		return true
	}
	return false
}

func init() {
	validate = validator.New()
	validate.RegisterValidation("mute", validateMute)
	validate.RegisterValidation("volume", validateVolume)
}
