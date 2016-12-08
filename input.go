package main

import "github.com/hajimehoshi/ebiten"

type InputHandler struct {
	inputDown bool
}

//the player cannot reverse direction,
//which we handle here. we basically see if x,y is
//the inverse of the current direction, in which case,
//we just return

//note that if the input is the same as last, it still
//sets it, but this has no effect

//x is the "leftright axis" , right is 1, left is -1

//this might be a "game controller" function,
//but we will keep it here for now because i think
//its the only function that will be really processed
//in said controller

func (game *GameState) HandleMovement(x int) {

	//find the intermediate direction by flipping the
	//current one
	newDirection := game.snake.direction.Reverse()

	//if we are going right, or clockwise, we check
	//to see if y is non-null in the reverse, if it is,
	//we multiply the value by -1 because it has crossed
	//an axes
	if newDirection.y != 0 && x == 1 {
		newDirection = newDirection.Multiply(-1)
	}

	//if we are going left, we check the x value,
	//and also multiply by -1 if it is non null,
	//because we have crossed the x axis
	if newDirection.x != 0 && x == -1 {
		newDirection = newDirection.Multiply(-1)
	}

	game.snake.direction = newDirection

	return

}

//process the input from Game, its understood that
//this will alter the state of game and input

func (input *InputHandler) ProcessInput(game *GameState) error {

	switch {

	case ebiten.IsKeyPressed(ebiten.KeyLeft):
		if !input.inputDown {
			game.HandleMovement(1)
		}
		input.inputDown = true

	case ebiten.IsKeyPressed(ebiten.KeyRight):
		if !input.inputDown {
			game.HandleMovement(-1)
		}
		input.inputDown = true

	default:
		input.inputDown = false

	}

	return nil
}
