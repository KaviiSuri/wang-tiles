package linalg

import (
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

// Map iterates through a vector and applies a function to it
func Map(v Vec, f func(a float64, i int) float64) Vec {
	values := []float64{}
	for i, val := range v.Values {
		values = append(values, f(val, i))
	}

	return Vec{
		Values: values,
		Size:   v.Size,
	}
}

// Construct A New Vector from Pair of vectors by applying the given argument `f` on each pair
func NewFromEachPair(v Vec, o Vec, f func(a, b float64) float64) Vec {
	checkEqualSize(v, o)
	return Map(v, func(a float64, i int) float64 {
		return f(a, o.Values[i])
	})
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
