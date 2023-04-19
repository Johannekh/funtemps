package conv

func FahrenheitToCelsius(value float64) float64 {
	return ((value - 32) * (5.0 / 9.0))
}
func KelvinToCelsius(value float64) float64 {
	return (value - 273.15)
}
func CelsiusToKelvin(value float64) float64 {
	return (value + 273.15)
}
func CelsiusToFahrenheit(value float64) float64 {
	return (value * (9.0/5.0) + 32)
}
func KelvinToFahrenheit(value float64) float64 {
	return ((value - 273.15) * (9.0/5.0) + 32)
}
func FahrenheitToKelvin(value float64) float64 {
	return ((value - 32) * (5.0/9.0) + 273.15)
}
