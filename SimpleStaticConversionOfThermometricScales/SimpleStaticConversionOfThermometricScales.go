package main

import "fmt"

const boilingPointOfWater = 373.00

func main() {
	temperatureKelvin := boilingPointOfWater
	temperatureCelsius := temperatureKelvin - 273.00

	fmt.Printf("A temperatura de ebulição da água em Kelvin é %.2f, já a temperatura de ebulição da água em Celsius é %.2f", temperatureKelvin, temperatureCelsius)
}
