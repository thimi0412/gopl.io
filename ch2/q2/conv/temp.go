package conv

import "fmt"

// Celsius .
type Celsius float64

// Fahrenheit .
type Fahrenheit float64

//  C
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g˚C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g˚F", f) }

// CToF convert celsius to fahrenheit .
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC convert fahrenheit to celsius .
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
