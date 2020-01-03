package main

import (
	"testing"
)

func TestBoolToStr(t *testing.T) {
	tables := []struct {
		value    bool
		expected string
	}{
		{false, "false"},
		{true, "true"},
	}

	for _, table := range tables {
		actual := boolToStr(table.value)
		if actual != table.expected {
			t.Errorf("String of (%v) was incorrect, got: %s, want: %s.", table.value, actual, table.expected)
		}
	}
}

func TestBoolToUpDown(t *testing.T) {
	tables := []struct {
		value    bool
		expected string
	}{
		{false, "down"},
		{true, "up"},
	}

	for _, table := range tables {
		actual := boolToUpDown(table.value)
		if actual != table.expected {
			t.Errorf("String of (%v) was incorrect, got: %s, want: %s.", table.value, actual, table.expected)
		}
	}
}

func TestFindMuteUrl(t *testing.T) {
	given := false
	actual := findMuteUrl(given)
	expected := "http://192.168.0.99/YamahaExtendedControl/v1/main/setMute?enable=false"
	if actual != expected {
		t.Errorf("Mute URL of (%v) was incorrect, got: %s, want: %s.", given, actual, expected)
	}
}

func TestFindVolumeUrl(t *testing.T) {
	given := false
	actual := findVolumeUrl(given)
	expected := "http://192.168.0.99/YamahaExtendedControl/v1/main/setVolume?volume=down"
	if actual != expected {
		t.Errorf("Mute URL of (%v) was incorrect, got: %s, want: %s.", given, actual, expected)
	}
}
