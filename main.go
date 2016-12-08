// the main loop and input handling
package main

import "github.com/hajimehoshi/ebiten"

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
	game := GameState{Snake{direction: Vector2{1, 0}, speed: 2}, InputHandler{}}
	game.snake.Start(Vector2{320 / 2, 240 / 2})
	ebiten.Run(game.Update, 320, 240, 2, "Go Go Snek!")
}
