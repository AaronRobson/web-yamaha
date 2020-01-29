package main

func muteToStr(mute bool) string {
	if mute {
		return "mute"
	} else {
		return "unmute"
	}
}

func boolToUpDown(volumeUp bool) string {
	if volumeUp {
		return "up"
	} else {
		return "down"
	}
}

func findMuteUrl(mute bool) string {
	return setMuteUrl + "?enable=" + boolToStr(mute)
}

func findVolumeUrl(volumeUp bool) string {
	return setVolumeUrl + "?volume=" + boolToUpDown(volumeUp)
}

const (
	hifiUrl      = "http://192.168.0.99/YamahaExtendedControl/v1"
	setMuteUrl   = hifiUrl + "/main/setMute"
	setVolumeUrl = hifiUrl + "/main/setVolume"
)
