package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

type Chef struct {
	appleImage *ebiten.Image
}

func (chef *Chef) MakeApple() Apple {
	apple := Apple{}
	apple.image = chef.appleImage
	return apple
}

func (chef *Chef) MakeEdibleForLevel(level int) Edible {
	var edible Edible

	switch level {
	case 0:
		apple := chef.MakeApple()
		apple.PlaceRandomly()
		edible = &apple
	}

	return edible
}

//start sets up the images for all the edibles

func (chef *Chef) Start() {

	//we create this just once, because it is a heavy struct
	chef.appleImage, _ = ebiten.NewImage(appleSize, appleSize, ebiten.FilterNearest)

	//and we do this just once, because it's fairly expensive
	chef.appleImage.Fill(color.Black)

}
