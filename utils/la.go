package utils

import "math"

type Vec2 struct {
	X float64
	Y float64
}

func (v *Vec2) Add(other *Vec2) Vec2 {
	return Vec2{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v *Vec2) Sub(other *Vec2) Vec2 {
	return Vec2{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

func (v *Vec2) Mul(other *Vec2) Vec2 {
	return Vec2{
		X: v.X * other.X,
		Y: v.Y * other.Y,
	}
}

func (v *Vec2) Divide(other *Vec2) Vec2 {
	return Vec2{
		X: v.X / other.X,
		Y: v.Y / other.Y,
	}
}
func (v *Vec2) Len() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}
