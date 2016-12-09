package main

import (
	"github.com/hajimehoshi/ebiten"
	"math/rand"
)

const appleSize = 8
const appleNutrition = 20

type Apple struct {

	//this thing is heavy, so, careful
	//when you deref
	image    *ebiten.Image
	location Vector2
}

func (apple *Apple) amount() int {
	return appleNutrition
}

func (apple *Apple) avatar() *ebiten.Image {
	return apple.image
}

func (apple *Apple) position() Vector2 {
	offset := apple.location.Add(Vector2{-appleSize / 2, -appleSize / 2})
	return offset
}

func (apple *Apple) Collider() Rect {
	offset := apple.location.Add(Vector2{-appleSize / 2, -appleSize / 2})
	return Rect{offset, offset.Add(Vector2{appleSize, appleSize})}
}

func (apple *Apple) PlaceRandomly() {
	apple.location = Vector2{
		float64(rand.Intn(int(GameWidth))),
		float64(rand.Intn(int(GameHeight)))}
}
