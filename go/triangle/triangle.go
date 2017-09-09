package triangle

import "sort"
import "math"

const testVersion = 3

func KindFromSides(a, b, c float64) Kind {
	if (math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c)) ||
		(math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0)) ||
		(a <= 0 || b <= 0 || c <= 0) {
		return NaT
	}

	s := []float64{a, b, c}
	sort.Float64s(s)

	if s[0]+s[1] < s[2] {
		return NaT
	}

	eq1 := s[0] == s[1]
	eq2 := s[1] == s[2]

	if eq1 && eq2 {
		return Equ
	} else if eq1 || eq2 {
		return Iso
	} else {
		return Sca
	}
}

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

// Pick values for the following identifiers used by the test program.
const (
	NaT = 0 // not a triangle
	Equ = 3 // equilateral
	Iso = 2 // isosceles
	Sca = 1 // scalene
)

// Organize your code for readability.
