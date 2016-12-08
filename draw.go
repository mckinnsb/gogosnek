// here we do our drawing of said snek

package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

//draw does not take a pointer to state, so we do not modifiy it

//i had this throw an error for the parent to collect,
//but i am no longer doing that. unsure if i want to remove it

func Draw(game GameState, screen *ebiten.Image) error {

	//background draw
	screen.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff})

	//draw the snake
	DrawSnake(game.snake, screen)

	if game.apple != nil {
		//draw the apple
		DrawEdible(game.apple, screen)
	}

	return nil

}

func DrawEdible(edible Edible, screen *ebiten.Image) {

	position := edible.position()

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(position.x, position.y)

	screen.DrawImage(edible.avatar(), opts)

}

//this draws the snake - it first draws the box and then
//figures out how to draw the tail from the last positions

//if any of the draws would go over the snakes head, we know
//we die, and we throw an "error" to show this

//kind of lazy, but it works because its snek

func DrawSnake(snake Snake, screen *ebiten.Image) {

	for position := range snake.GetTail() {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(position.x, position.y)

		screen.DrawImage(snake.avatar, opts)
	}

	return

}
