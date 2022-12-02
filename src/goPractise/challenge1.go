package main

import "fmt"

// aliasing type
type Celsius float32
type Fahrenheit float32

func main() {
	fmt.Println(toFahrenheit(100))
}

// Function to convert celsius to fahrenheit
func toFahrenheit(t Celsius) Fahrenheit {

	var temp Fahrenheit
	temp = Fahrenheit((t * Celsius(9) / Celsius(5)) + Celsius(32))
	return temp

}
