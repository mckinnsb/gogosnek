package main

import (
	"github.com/hajimehoshi/ebiten"
	"math/rand"
)

const appleSize int = 8
const appleNutrition int = 20

type Edible interface {
	amount() int
	avatar() *ebiten.Image
	position() Vector2
	size() Vector2
}

type Apple struct {

	//this thing is heavy, so, careful
	//when you deref
	image    *ebiten.Image
	location Vector2
}

func (apple Apple) amount() int {
	return appleNutrition
}

func (apple Apple) avatar() *ebiten.Image {
	return apple.image
}

func (apple Apple) position() Vector2 {
	return apple.location
}

func (apple Apple) size() Vector2 {
	return Vector2{float64(appleSize), float64(appleSize)}
}

func (apple *Apple) PlaceRandomly() {
	apple.location = Vector2{
		float64(rand.Intn(int(GameWidth))),
		float64(rand.Intn(int(GameHeight)))}
}
