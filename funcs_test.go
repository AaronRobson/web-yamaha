package main

import (
	"testing"
)

var muteTests = []struct {
	given    bool
	expected string
}{
	{false, "unmute"},
	{true, "mute"},
}

func TestMuteToStr(t *testing.T) {
	for _, tt := range muteTests {
		actual := muteToStr(tt.given)
		if actual != tt.expected {
			t.Errorf(
				"muteToStr(%v) returned: '%v' expected '%v'.",
				tt.given, actual, tt.expected)
		}
	}
}
