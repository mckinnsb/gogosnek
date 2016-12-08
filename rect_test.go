package main

import "testing"

func TestRectCollidesWithPositive(t *testing.T) {

	one := Rect{Vector2{0, 0}, Vector2{4, 4}}
	two := Rect{Vector2{-1, -1}, Vector2{3, 3}}

	if !one.CollidesWith(two) {
		t.Error("positive collision check failed")
	}

}

func TestRectCollidesWithNegative(t *testing.T) {

	one := Rect{Vector2{0, 0}, Vector2{4, 4}}
	two := Rect{Vector2{0, -4}, Vector2{4, -.05}}

	if one.CollidesWith(two) {
		t.Error("negative collision check failed")
	}

}
