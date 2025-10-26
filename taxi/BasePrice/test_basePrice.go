package basePrice

import "testing"

func TestCalculateBasePrice(t *testing.T) {

	trip1 := TripParameters{Distance: 8.5, Duration: 20}

	got := CalculateBasePrice(trip1)
	want := 125
	if got != float64(want) {
		t.Fail()
	}

}
