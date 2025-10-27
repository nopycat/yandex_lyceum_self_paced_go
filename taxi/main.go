package taxi

import (
	basePrice "awesomeProject/taxi/BasePrice"
	trafficCongestion "awesomeProject/taxi/TrafficCongestion"
	weatherCoefficient "awesomeProject/taxi/WeatherCoefficient"
	"fmt"
	"time"
)

func main() {
	calculator := trafficCongestion.PriceCalculator{
		TrafficClient: &trafficCongestion.RealTrafficClient{}, // В продакшене используется настоящий клиент, а мы подключим структуру-заглушку для имитации его работы
	}

	price := calculator.CalculatePrice(
		basePrice.TripParameters{Distance: 8.5, Duration: 20},
		time.Now(),
		weatherCoefficient.WeatherData{Condition: weatherCoefficient.HeavyRain, WindSpeed: 10},
		55.751244, 37.618423,
	)

	fmt.Printf("Ваша цена: %.0f руб.\n", price)
}
