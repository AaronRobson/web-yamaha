package main

func muteToStr(mute bool) string {
	if mute {
		return "mute"
	}
	return "unmute"
}

func boolToUpDown(volumeUp bool) string {
	if volumeUp {
		return "up"
	}
	return "down"
}

func findMuteURL(mute bool) string {
	return setMuteURL + "?enable=" + boolToStr(mute)
}

func findVolumeURL(volumeUp bool) string {
	return setVolumeURL + "?volume=" + boolToUpDown(volumeUp)
}

const (
	hifiURL      = "http://192.168.0.99/YamahaExtendedControl/v1"
	setMuteURL   = hifiURL + "/main/setMute"
	setVolumeURL = hifiURL + "/main/setVolume"
)
