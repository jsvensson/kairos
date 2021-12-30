package weather

import "time"

// Now generates a value in the 0..100 range that can be mapped to the percentages of
// weather in zones.
func Now() int {
	return Forecast(time.Now())
}

// Forecast generates a value in the 0..100 range that can be mapped to the percentages of
// weather in zones. It accepts a timestamp for which to generate the forecast.
func Forecast(t time.Time) int {
	// Algorithm from the SaintCoinach library
	// https://github.com/Rogueadyn/SaintCoinach

	currentSecond := t.Unix()

	// 175 seconds/bell
	bell := currentSecond / 175

	// Hour 0 = 8:00 ET
	// Hour 8 = 16:00 ET
	increment := (bell + 8 - (bell % 8)) % 24

	days := currentSecond / 4200 // seconds in an Eorzea day
	days = days << 32

	// 0x64 = 100
	base := days*100 + increment

	// 0xB = 11
	forecast := (base << 11) ^ base
	forecast = (forecast >> 8) ^ forecast

	return int(forecast % 100)
}
