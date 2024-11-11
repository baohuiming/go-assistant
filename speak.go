package main

import (
	"log"
	"os"

	edgeTTS "github.com/lib-x/edgetts"
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
	opts := make([]edgeTTS.Option, 0)
	opts = append(opts, edgeTTS.WithVoice(getSpeakerCode(speaker)))
	speech, err := edgeTTS.NewSpeech(opts...)
	if err != nil {
		log.Fatal(err)
	}
	audio, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	err = speech.AddSingleTask(text, audio)
	if err != nil {
		log.Fatal(err)
	}

	err = speech.StartTasks()
	if err != nil {
		log.Fatal(err)
	}
}
