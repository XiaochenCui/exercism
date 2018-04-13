package space

// Planet is string
type Planet string

// Age return years in that planet
func Age(seconds float64, planet Planet) float64 {
	secondsOfEarthYear := 31557600
	planetYear := map[string]float64{
		"Earth":   1.0,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	earthYears := seconds / float64(secondsOfEarthYear)
	planetYears := earthYears / planetYear[string(planet)]
	return planetYears
}
