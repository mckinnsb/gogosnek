// here we do our drawing of said snek

package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

//draw does not take a pointer to state, so we do not modifiy it

func Draw(game GameState, screen *ebiten.Image) error {

	//background draw
	screen.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff})

	//draw the snake
	err := DrawSnake(game.snake, screen)

	return err

}

//this draws the snake - it first draws the box and then
//figures out how to draw the tail from the last positions

//if any of the draws would go over the snakes head, we know
//we die, and we throw an "error" to show this

//kind of lazy, but it works because its snek

func DrawSnake(snake Snake, screen *ebiten.Image) error {

	//we could optimize this by only creating this image
	//once

	//these have to be outside the loop, because
	//NewImage creates a nxn array, and guess how espensive that is
	square, _ := ebiten.NewImage(16, 16, ebiten.FilterNearest)

	//especially when you do this, and it goes over every element again
	square.Fill(color.White)

	for position := range snake.GetTail() {

		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(position.x, position.y)

		screen.DrawImage(square, opts)

	}

	return nil

}
