//some simple vector stuff for moving the snake around
package main

type Vector2 struct {
	x, y float64
}

func (v1 Vector2) Add(v2 Vector2) Vector2 {
	return AddVectors(v1, v2)
}

func AddVectors(v1, v2 Vector2) Vector2 {
	v1.x = v1.x + v2.x
	v1.y = v1.y + v2.y
	return v1
}

func (v Vector2) Multiply(scale float64) Vector2 {
	return MultiplyVector(v, scale)
}

func MultiplyVector(v Vector2, scale float64) Vector2 {
	v.x = v.x * scale
	v.y = v.y * scale
	return v
}

func (v Vector2) Reverse() Vector2 {
	return Vector2{v.y, v.x}
}
