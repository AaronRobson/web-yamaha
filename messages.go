package main

import "gopkg.in/go-playground/validator.v10"

type ErrorResponse struct {
	Message string `json:"message"`
}

type MuteRequest struct {
	Mute string `json:"mute" validate:"mute"`
}

type VolumeRequest struct {
	Volume string `json:"volume" validate:"volume"`
}

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

func init() {
	validate = validator.New()
	validate.RegisterValidation("mute", ValidateMute)
	validate.RegisterValidation("volume", ValidateVolume)
}
