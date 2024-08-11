package main

import "log"

func main() {
	weather := GetWeatherText()
	log.Println(weather)
	chat_reply := Chat(weather)
	log.Println(chat_reply)
	Speak(chat_reply, "晓晓", "output.mp3")
	Play()
}
