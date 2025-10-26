package trafficCongestion

import (
	"awesomeProject/taxi/weatherCoefficient"
)

func GetWeatherMultiplier(weather weatherCoefficient.WeatherData) float64 {
	coef := 1.0
	switch weather.Condition {
	case weatherCoefficient.Rain:
		coef += 0.125
	case weatherCoefficient.HeavyRain:
		coef += 0.2
	case weatherCoefficient.Snow:
		coef += 0.15
	}
	if weather.WindSpeed > 15 {
		coef += 0.1
	}
	return coef
}
