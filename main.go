// the main loop and input handling
package main

import "github.com/hajimehoshi/ebiten"

const GameWidth float64 = 320
const GameHeight float64 = 240

type GameState struct {
	snake Snake
	input InputHandler
}

func (game *GameState) Update(screen *ebiten.Image) error {

	//handled in input.go
	err := game.input.ProcessInput(game)

	if err != nil {
		return err
	}

	//handled in snake.go
	game.snake.Update()

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
		InputHandler{}}

	game.snake.Start(Vector2{GameWidth / 2, GameHeight / 2})

	ebiten.Run(
		game.Update,
		int(GameWidth),
		int(GameHeight),
		2,
		"Go Go Snek!")

}
