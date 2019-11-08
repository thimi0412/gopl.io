package conv

import "fmt"

// Kilogram .
type Kilogram float64

// Pound .
type Pound float64

func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }
func (p Pound) String() string    { return fmt.Sprintf("%glb", p) }

// KtoP convert kilogram to pound.
func KtoP(k Kilogram) Pound { return Pound(k * 2.20462) }

// PtoK convert pound to kilogram.
func PtoK(p Pound) Kilogram { return Kilogram(p / 2.20462) }
