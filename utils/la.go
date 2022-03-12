package utils

type Vec2 struct {
	x int64
	y int64
}

func (v *Vec2) Add(other *Vec2) Vec2 {
	return Vec2{
		x: v.x + other.x,
		y: v.y + other.y,
	}
}

func (v *Vec2) Sub(other *Vec2) Vec2 {
	return Vec2{
		x: v.x - other.x,
		y: v.y - other.y,
	}
}

func (v *Vec2) Mul(other *Vec2) Vec2 {
	return Vec2{
		x: v.x * other.x,
		y: v.y * other.y,
	}
}

func (v *Vec2) Divide(other *Vec2) Vec2 {
	return Vec2{
		x: v.x / other.x,
		y: v.y / other.y,
	}
}
