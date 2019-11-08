package conv

import "fmt"

// Meter .
type Meter float64

// Feet .
type Feet float64

func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
func (f Feet) String() string  { return fmt.Sprintf("%gft", f) }

// MToFe converts meters to feet.
func MToFe(m Meter) Feet { return Feet(m / 0.3048) }

// FeToM converts feet to meters.
func FeToM(f Feet) Meter { return Meter(f * 0.3048) }
