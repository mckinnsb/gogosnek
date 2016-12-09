//some simple vector stuff for moving the snake around
package main

import "math"

type Vector2 struct {
	x, y float64
}

//add a vector to this one and return the new vector
func (v1 Vector2) Add(v2 Vector2) Vector2 {
	return AddVectors(v1, v2)
}

//add two vectors and return a new one
func AddVectors(v1, v2 Vector2) Vector2 {
	v1.x = v1.x + v2.x
	v1.y = v1.y + v2.y
	return v1
}

//uses constants in main.go
func (v Vector2) ClampToWindow() Vector2 {

	switch {

	//since the first argument always determines
	//the sign, sometimes it is negative and
	//we have to take the absolute value

	case v.x < 0 || v.x > GameWidth:
		v.x = math.Mod(v.x, GameWidth)
		if v.x < 0 {
			v.x += GameWidth + 1
		} else {
			v.x += -1
		}

	case v.y < 0 || v.y > GameHeight:
		v.y = math.Mod(v.y, GameHeight)
		if v.y < 0 {
			v.y += GameHeight + 1
		} else {
			v.y += -1
		}
	}

	return v

}

//multiplay a vector by a scale
func (v Vector2) Multiply(scale float64) Vector2 {
	return MultiplyVector(v, scale)
}

//multiplay a vector by a scale
func MultiplyVector(v Vector2, scale float64) Vector2 {
	v.x = v.x * scale
	v.y = v.y * scale
	return v
}

//reverse the vector, swapping x and y
func (v Vector2) Reverse() Vector2 {
	return Vector2{v.y, v.x}
}
