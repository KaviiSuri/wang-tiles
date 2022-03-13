package linalg

import "log"

func Add(v, o Vec) Vec {
	return NewFromEachPair(v, o, func(a, b float64) float64 {
		return a + b
	})
}

func Sub(v, o Vec) Vec {
	return NewFromEachPair(v, o, func(a, b float64) float64 {
		return a - b
	})
}
func Mul(v, o Vec) Vec {
	return NewFromEachPair(v, o, func(a, b float64) float64 {
		return a * b
	})
}

func Divide(v, o Vec) Vec {
	return NewFromEachPair(v, o, func(a, b float64) float64 {
		return a / b
	})
}

func Max(v, o Vec) Vec {
	return NewFromEachPair(v, o, func(a, b float64) float64 {
		if a > b {
			return a
		} else {
			return b
		}
	})
}

func Min(v, o Vec) Vec {
	return NewFromEachPair(v, o, func(a, b float64) float64 {
		if a < b {
			return a
		} else {
			return b
		}
	})
}

func Lerp(v, o, t Vec) Vec {
	checkEqualSize(v, o)
	return Map(v, func(a float64, i int) float64 {
		return a + (o.Values[i]-a)*t.Values[i]
	})
}

func checkEqualSize(v Vec, vectors ...Vec) {
	for _, o := range vectors {
		if v.Size != o.Size {
			log.Fatalf("Vectors Should be of the same size: %v != %v", v.Size, o.Size)
		}
	}
}
