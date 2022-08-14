// Ex2.2
package convformat

import "fmt"

type Pound float64
type Kilogram float64

const (
	KilogramsInPound Kilogram = 2.204
)

func (f Pound) String() string    { return fmt.Sprintf("%g lb", f) }
func (m Kilogram) String() string { return fmt.Sprintf("%g kg", m) }

func PToK(p Pound) Kilogram { return Kilogram(p) / KilogramsInPound }

func KToP(k Kilogram) Pound { return Pound(k * KilogramsInPound) }
