package main

import (
	"testing"
)

func TestValidateMute(t *testing.T) {
	tables := []struct {
		value         string
		expectedValid bool
	}{
		{"mute", true},
		{"unmute", true},
		{"invalid", false},
	}

	for _, table := range tables {
		obj := MuteRequest{
			Mute: table.value,
		}
		err := validate.Struct(obj)
		actualValid := err == nil
		if actualValid != table.expectedValid {
			t.Errorf("%v valid status incorrect, got: %v, want: %v.", obj, actualValid, table.expectedValid)
		}
	}
}

func TestValidateVolume(t *testing.T) {
	tables := []struct {
		value         string
		expectedValid bool
	}{
		{"up", true},
		{"down", true},
		{"invalid", false},
	}

	for _, table := range tables {
		obj := VolumeRequest{
			Volume: table.value,
		}
		err := validate.Struct(obj)
		actualValid := err == nil
		if actualValid != table.expectedValid {
			t.Errorf("%v valid status incorrect, got: %v, want: %v.", obj, actualValid, table.expectedValid)
		}
	}
}
