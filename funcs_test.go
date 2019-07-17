package main

import (
	"testing"
)

func TestMuteToStr(t *testing.T) {
	var tables = []struct {
		given    bool
		expected string
	}{
		{false, "unmute"},
		{true, "mute"},
	}

	for _, table := range tables {
		actual := muteToStr(table.given)
		if actual != table.expected {
			t.Errorf(
				"muteToStr(%v) returned: '%v' expected '%v'.",
				table.given, actual, table.expected)
		}
	}
}
