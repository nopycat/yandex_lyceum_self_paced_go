package trafficCongestion

import (
	basePrice "awesomeProject/taxi/BasePrice"
	timeCoefficient "awesomeProject/taxi/TimeCoefficients"
	weatherCoefficient "awesomeProject/taxi/WeatherCoefficient"
	"math"
	"time"
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

type TrafficClient interface {
	GetTrafficLevel(lat, lng float64) int // 1–5
}

func GetTrafficMultiplier(trafficLevel int) float64 {
	return 1.0 + float64(trafficLevel-1)*0.1
}

type PriceCalculator struct {
	TrafficClient TrafficClient
}

type RealTrafficClient struct{}

func (c *RealTrafficClient) GetTrafficLevel(lat, lng float64) int {
	return 3 // Константное значение в нашем примере, в реальности оно будет вычисляться сервисом Яндекс Карты
}

func (c *PriceCalculator) CalculatePrice(trip basePrice.TripParameters, now time.Time, weather weatherCoefficient.WeatherData, lat, lng float64) float64 {
	base := basePrice.CalculateBasePrice(trip)
	timeMult := timeCoefficient.GetTimeMultiplier(now)
	weatherMult := GetWeatherMultiplier(weather)
	trafficMult := GetTrafficMultiplier(c.TrafficClient.GetTrafficLevel(lat, lng))

	finalPrice := base * timeMult * weatherMult * trafficMult

	return basePrice.ApplyPriceLimits(math.Round(finalPrice))
}
