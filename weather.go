package main

import (
	"fmt"
	"log"
	"time"

	owm "github.com/briandowns/openweathermap"
)

func GetWeatherText() string {
	configP := ReadConfig()
	apiKey := configP.Weather.APIKey
	longitude := configP.Weather.Longitude
	latitude := configP.Weather.Latitude

	w, err := owm.NewForecast("5", "C", "zh_cn", apiKey) // 5 day forecast in Celsius with Chinese output
	//w, err := owm.NewCurrent("C", "zh_cn", apiKey)
	if err != nil {
		log.Fatalf("Error get weather: %v", err)
	}

	err = w.DailyByCoordinates(&owm.Coordinates{
		Longitude: longitude,
		Latitude:  latitude,
	}, 6) // 6 * 3h_forecast

	if err != nil {
		log.Fatalf("Error get weather: %v", err)
	}

	forecast := w.ForecastWeatherJson.(*owm.Forecast5WeatherData)

	text := "当前时间是：" + time.Now().Format("2006-01-02 15:04") + "，"
	text += "以下是未来18小时天气预报：\n"
	for item := range forecast.List {
		// 时间
		text += forecast.List[item].DtTxt.Local().Format("2006-01-02 15:04") + "，"
		// 温度
		text += fmt.Sprintf("温度：%v℃，", forecast.List[item].Main.Temp)
		// 天气
		text += forecast.List[item].Weather[0].Description + "，"
		// 雨量
		text += fmt.Sprintf("雨量(1h)：%vmm，", forecast.List[item].Rain.OneH)
		text += fmt.Sprintf("雨量(3h)：%vmm，", forecast.List[item].Rain.ThreeH)
		// 雪量
		text += fmt.Sprintf("雪量(1h)：%vmm，", forecast.List[item].Snow.OneH)
		text += fmt.Sprintf("雪量(3h)：%vmm。\n", forecast.List[item].Snow.ThreeH)
	}
	return text
}
