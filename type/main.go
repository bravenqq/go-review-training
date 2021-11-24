// Package main provides ...
package main

import "fmt"

//START OMIT/
type Celsius float64
type Fahrenheit float64
type Kelvins float64

const (
	AbsolutedZeroC Celsius = -273.15
	FreezingC      Celsius = 0
	BoilingC       Celsius = 100
)

const (
	AbsolutedZero Kelvins = 0
	Freezing      Kelvins = 273.15
	Boiling       Kelvins = 373.15
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius(f-32) * 5 / 9 }

func KToC(k Kelvins) Celsius { return Celsius(float64(k) - float64(Freezing)) }
func CToK(c Celsius) Kelvins { return Kelvins(float64(c) + float64(Freezing)) }

func main() {
	freezing := CToF(FreezingC)
	fmt.Println("the temperature is:", freezing)
}

//END OMIT/
