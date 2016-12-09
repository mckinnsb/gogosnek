// here we do our drawing of said snek

package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

type Drawable interface {
	//this the image for the edible
	avatar() *ebiten.Image
	//this is used for drawing
	position() Vector2
}

//draw does not take a pointer to state, so we do not modifiy it

//i had this throw an error for the parent to collect,
//but i am no longer doing that. unsure if i want to remove it

func Draw(game GameState, screen *ebiten.Image) error {

	//background draw
	screen.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff})

	//this might be a real data object someday
	drawable := []Drawable{game.apple}

	for _, drawObject := range drawable {
		DrawObject(drawObject, screen)
	}

	for drawObject := range game.snake.GetDrawTail() {
		DrawObject(drawObject, screen)
	}

	return nil

}

func DrawObject(drawable Drawable, screen *ebiten.Image) {

	if drawable == nil {
		return
	}

	if drawable.avatar() == nil {
		return
	}

	position := drawable.position()

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(position.x, position.y)

	screen.DrawImage(drawable.avatar(), opts)

}
