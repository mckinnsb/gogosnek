package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

//this is just used for advanced snek

type SnakeImage struct {
	images map[int]*ebiten.Image
}

func (image *SnakeImage) Start() {

	image.images = make(map[int]*ebiten.Image)

	for i := startSize; i <= maxSize; i += 2 {

		//make a new image
		newImage, _ := ebiten.NewImage(i, i, ebiten.FilterNearest)

		rb := uint8(0xff * (maxSize - i) / maxSize)

		newImage.Fill(
			color.NRGBA{
				rb,
				0xff,
				rb,
				0xff})

		image.images[i] = newImage

	}

}
