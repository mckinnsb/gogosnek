package main

import "testing"

func TestClampMaxX(t *testing.T) {
	v := Vector2{GameWidth + 1, 0}
	if v.ClampToWindow().x != 1 {
		t.Error("vector did not clamp past max X correctly")
	}
}

func TestClampMinX(t *testing.T) {
	v := Vector2{-1, 0}
	if v.ClampToWindow().x != GameWidth-1 {
		t.Error("vector did not clamp before min X correctly")
	}
}

func TestClampMaxY(t *testing.T) {
	v := Vector2{0, GameHeight + 1}
	if v.ClampToWindow().y != 1 {
		t.Error("vector did not clamp past max Y correctly")
	}
}

func TestClampMinY(t *testing.T) {
	v := Vector2{0, GameHeight - 1}
	if v.ClampToWindow().y != GameHeight-1 {
		t.Error("vector did not clamp before min Y correctly")
	}
}
