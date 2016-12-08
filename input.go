package main

import "github.com/hajimehoshi/ebiten"

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
