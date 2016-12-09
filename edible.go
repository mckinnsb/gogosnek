package main

import "github.com/hajimehoshi/ebiten"

//you might notice if you look at draw.go,
//that this interface overlaps with
//drawable, edible has to be drawable

type Edible interface {
	//this is the amount that the edible makes the snake grow by
	amount() int
	//this the image for the edible
	avatar() *ebiten.Image
	//this is used for drawing
	position() Vector2
	//this is used for collision
	Collider() Rect
}
