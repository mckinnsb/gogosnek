// the main loop and input handling
package main

import "github.com/hajimehoshi/ebiten"

type GameState struct {
	snake Snake
}

func (game *GameState) HandleMovement(x, y float64) {
	game.snake.direction = Vector2{x, y}

	//game.snake.position.Add(Vector2{x, y})
	return
}

func (game *GameState) ProcessInput() error {

	switch {

	case ebiten.IsKeyPressed(ebiten.KeyUp):
		game.HandleMovement(0, -1)

	case ebiten.IsKeyPressed(ebiten.KeyDown):
		game.HandleMovement(0, 1)

	case ebiten.IsKeyPressed(ebiten.KeyLeft):
		game.HandleMovement(-1, 0)

	case ebiten.IsKeyPressed(ebiten.KeyRight):
		game.HandleMovement(1, 0)

	}

	return nil
}

func (game *GameState) Update(screen *ebiten.Image) error {

	err := game.ProcessInput()

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
	game := GameState{Snake{speed: 2}}
	game.snake.Start(Vector2{320 / 2, 240 / 2})
	ebiten.Run(game.Update, 320, 240, 2, "Go Go Snek!")
}
