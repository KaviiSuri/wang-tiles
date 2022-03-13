package linalg

import (
	"log"
	"math"
)

// Vector Type that supports basic arithmatic operations
type Vec struct {
	Size   int
	Values []float64
}

// Methods

func (v Vec) Len() float64 {
	result := 0.0
	for _, val := range v.Values {
		result += val * val
	}
	return math.Sqrt(result)
}

// Contructors

// Constructs A New Vector From Arguments
func NewVec(values ...float64) Vec {
	return Vec{
		Size:   len(values),
		Values: values,
	}
}

// Constructs a New Vector with given size
func NewSizedVec(size int, value float64) Vec {
	values := []float64{}
	for i := 0; i < size; i++ {
		values = append(values, value)
	}
	return Vec{
		Values: values,
		Size:   size,
	}
}

// Construct A New Vector from Pair of vectors by applying the given argument `f` on each pair
func NewFromEachPair(v Vec, o Vec, f func(a, b float64) float64) Vec {
	checkEqualSize(v, o)
	values := []float64{}
	for i, val := range v.Values {
		values = append(values, f(val, o.Values[i]))
	}

	return Vec{
		Values: values,
		Size:   v.Size,
	}
}

// Functions

func Add(v, o Vec) Vec {
	return NewFromEachPair(v, o, func(a, b float64) float64 {
		return a + b
	})
}

func Sub(v, o Vec) Vec {
	return NewFromEachPair(v, o, func(a, b float64) float64 {
		return a + b
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

func checkEqualSize(v, o Vec) {
	if v.Size != o.Size {
		log.Fatalf("Vectors Should be of the same size: %v != %v", v.Size, o.Size)
	}
}

// Accessors Aliases

func (v Vec) X() float64 {
	return v.Values[0]
}
func (v Vec) Y() float64 {
	return v.Values[1]
}
func (v Vec) Z() float64 {
	return v.Values[2]
}

func (v Vec) U() float64 {
	return v.Values[0]
}
func (v Vec) V() float64 {
	return v.Values[1]
}

func (v Vec) R() float64 {
	return v.Values[0]
}
func (v Vec) G() float64 {
	return v.Values[1]
}
func (v Vec) B() float64 {
	return v.Values[2]
}
func (v Vec) A() float64 {
	return v.Values[3]
}
