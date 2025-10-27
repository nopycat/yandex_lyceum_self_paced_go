package basePrice

const (
	pricePerKm     = 10.0 // Руб за км (в реальности это сложная формула)
	pricePerMinute = 2.0  // Руб за минуту
	minPrice       = 99.0
	maxPrice       = 20000.0
)

type TripParameters struct {
	Distance float64 // Расстояние в км (рассчитывается через Яндекс Карты)
	Duration float64 // Время в минутах (прогнозируется с учётом истории поездок)
}

func CalculateBasePrice(tp TripParameters) float64 {
	// Реализуем простую формулу: цена = расстояние * тариф + время * тариф
	return tp.Distance*pricePerKm + tp.Duration*pricePerMinute
}

func ApplyPriceLimits(price float64) float64 {
	switch {
	case price < minPrice:
		return minPrice
	case price > maxPrice:
		return maxPrice
	default:
		return price
	}
}
