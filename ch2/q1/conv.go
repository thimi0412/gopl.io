package tempconv

// CToF .
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC .
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK .
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// FToK .
func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }

// KToC .
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

// KToF .
func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }
