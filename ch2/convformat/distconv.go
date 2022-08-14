// Ex2.2
package convformat

import "fmt"

type Foot float64
type Metre float64

const (
	FootsInMetre Foot = 3.2808
)

func (f Foot) String() string  { return fmt.Sprintf("%gm", f) }
func (m Metre) String() string { return fmt.Sprintf("%gft", m) }

func FToM(f Foot) Metre { return Metre(f * FootsInMetre) }

func MToF(m Metre) Foot { return Foot(m) / FootsInMetre }
