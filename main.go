// the main loop and input handling
package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const GameWidth float64 = 320
const GameHeight float64 = 240

type GameState struct {
	snake      Snake
	input      InputHandler
	apple      Edible
	chef       Chef
	ended      bool
	ticker     int
	tick       int
	endMessage string
}

func (game *GameState) End(msg string) {
	game.ended = true
	game.endMessage = msg
}

func (game *GameState) Update(screen *ebiten.Image) error {

	game.tick += 1

	if game.tick%game.ticker != 0 {
		err := Draw(*game, screen)
		return err
	}

	//handled in input.go
	err := game.input.ProcessInput(game)

	if err != nil {
		return err
	}

	//for now, will think of another way to do this later
	if game.apple == nil {
		apple := game.chef.MakeApple()
		apple.PlaceRandomly()

		//apple has pointer recievers, so we must reference
		//it here
		game.apple = &apple
	}

	//handled in snake.go
	game.snake.Update()

	if game.snake.EatingTail() {
		//we will do a score later
		game.End("Game Over!")
	}

	//this might be in a collision handler that tracks
	//all edibles, but for now this is easiest
	if game.snake.IsColliding(game.apple) {
		game.snake.Eat(game.apple)
		game.apple = nil
	}

	//handled in draw.go, this is not a pointer fn
	//because it can't change state
	err = Draw(*game, screen)

	if err != nil {
		return err
	}

	return nil

}

func main() {

	game := GameState{
		Snake{direction: Vector2{1, 0}, speed: 2},
		InputHandler{},
		nil,
		Chef{},
		false,
		1,
		0,
		""}

	game.chef.Start()
	game.snake.Start(Vector2{GameWidth / 2, GameHeight / 2})

	updateFunction := func(screen *ebiten.Image) error {

		if game.ended {
			//we will do a score later
			ebitenutil.DebugPrint(screen, game.endMessage)
		} else {
			game.Update(screen)
		}

		return nil

	}

	ebiten.Run(
		updateFunction,
		int(GameWidth),
		int(GameHeight),
		2,
		"Go Go Snek!")

}
