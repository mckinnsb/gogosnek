package main

import (
	"fmt"
	"testing"
)

func TestClampMaxX(t *testing.T) {
	v := Vector2{GameWidth + 1, 0}

	if v.ClampToWindow().x != 0 {
		t.Error(fmt.Sprintf(
			"vector did not clamp past max X correctly, expected %v, got %v",
			0,
			v))
	}
}

func TestClampMinX(t *testing.T) {
	v := Vector2{-1, 0}

	if v.ClampToWindow().x != GameWidth {
		t.Error(fmt.Sprintf(
			"vector did not clamp before min X correctly, expected %v, got %v",
			GameWidth,
			v.ClampToWindow().x))
	}
}

func TestClampMaxY(t *testing.T) {
	v := Vector2{0, GameHeight + 1}

	if v.ClampToWindow().y != 0 {
		t.Error(fmt.Sprintf(
			"vector did not clamp past max Y correctly, expected %v, got %v",
			0,
			v.ClampToWindow().y))
	}
}

func TestClampMinY(t *testing.T) {
	v := Vector2{0, -1}

	if v.ClampToWindow().y != GameHeight {
		t.Error(fmt.Sprintf(
			"vector did not clamp before min Y correctly, expected %v, got %v",
			GameHeight,
			v.ClampToWindow().y))
	}

}
