package main

import (
	"github.com/baohuiming/edge-tts-go/edgeTTS"
)

var displayShortMap = map[string]string{
	"晓晓": "zh-CN-XiaoxiaoNeural",
	"晓伊": "zh-CN-XiaoyiNeural",
	"云健": "zh-CN-YunjianNeural",
	"云希": "zh-CN-YunxiNeural",
	"云夏": "zh-CN-YunxiaNeural",
	"云扬": "zh-CN-YunyangNeural",
}

// 获取发音人的代码
func getSpeakerCode(display string) string {
	display, ok := displayShortMap[display]
	if ok {
		return display
	}
	return ""
}

// text: text to speak
// speaker: speaker name [晓晓, 晓伊, 云健, 云希, 云夏, 云扬]
func Speak(text string, speaker string, path string) {
	// edgeTTS.NewTTS(args).AddText(args.Text, args.Voice, args.Rate, args.Volume).Speak()
	args := edgeTTS.Args{
		WriteMedia: path,
	}
	tts := edgeTTS.NewTTS(args)

	tts.AddTextWithVoice(text, getSpeakerCode(speaker))

	tts.Speak()
}
