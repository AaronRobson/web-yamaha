package main

func muteToStr(mute bool) string {
	if mute {
		return "mute"
	} else {
		return "unmute"
	}
}

func findMuteUrl(mute bool) string {
	return setMuteUrl + "?enable=" + boolToStr(mute)
}

const (
	hifiUrl    = "http://192.168.0.99/YamahaExtendedControl/v1"
	setMuteUrl = hifiUrl + "/main/setMute"
)
