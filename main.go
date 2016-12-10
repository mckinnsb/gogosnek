// the main loop and input handling
package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"math"
	"time"
)

const GameWidth float64 = 320
const GameHeight float64 = 240
const MaxLevels int = 0

//ebiten sets this
const FPS = 60

type GameState struct {
	snake          Snake
	input          InputHandler
	coroutine      func(game *GameState, screen *ebiten.Image)
	coroutineStart int //the tick at which the coroutine started
	//its assumed that 60 ticks = 1s
	chef       Chef
	edibles    []Edible
	ended      bool
	frozen     bool
	level      int
	ticker     int
	tick       int
	score      int
	running    bool
	endMessage string
}

func (game *GameState) AdvanceLevel(screen *ebiten.Image) {

	//yeah. this should be an enum
	game.frozen = true
	game.running = false
	game.coroutine = DoLevelAnimation
	game.coroutineStart = game.tick

	go func() {

		time.Sleep(time.Second * 3)

		game.snake.ResetToStartSize()
		game.NextLevel()

		game.running = true
		game.frozen = false

	}()

}

func (game *GameState) End(msg string) {
	game.ended = true
	game.running = false
	game.endMessage = msg
}

func (game *GameState) Main(screen *ebiten.Image) error {

	game.tick += 1

	switch {
	case game.ended:
		ebitenutil.DebugPrint(screen, game.endMessage)
	case game.running:
		game.Update(screen)
	case game.frozen:
		if game.coroutine != nil {
			game.coroutine(game, screen)
		}
	default:
		ebitenutil.DebugPrint(screen, "Advanced Snek? (Y/n)")
		game.input.WaitForOption(game)
	}

	return nil

}

func (game *GameState) NextLevel() {

	game.level += 1

	//this is the kind of thing where no coersion at all
	//starts to look silly

	game.level = int(
		math.Min(
			float64(MaxLevels),
			float64(game.level)))

}

func (game *GameState) Start() {
	game.chef.Start()
	game.snake.Start(Vector2{GameWidth / 2, GameHeight / 2})
	game.running = true
}

func (game *GameState) Update(screen *ebiten.Image) error {

	if game.tick%game.ticker != 0 {
		err := Draw(*game, screen)
		return err
	}

	if game.snake.advanced && game.snake.size == maxSize {
		game.AdvanceLevel(screen)
		return nil
	}

	//handled in input.go
	err := game.input.ProcessInput(game)

	if err != nil {
		return err
	}

	//for now, will think of another way to do this later
	if game.edibles[0] == nil {
		game.edibles[0] = game.chef.MakeEdibleForLevel(game.level)
	}

	//handled in snake.go
	game.snake.Update()

	if game.snake.EatingTail() {
		//end with score message
		game.End(fmt.Sprintf("Game Over! Score: %v", game.score))
	}

	for i := 0; i < len(game.edibles); i++ {

		edible := game.edibles[i]

		if edible == nil {
			continue
		}

		//this might be in a collision handler that tracks
		//all edibles, but for now this is easiest
		if game.snake.IsColliding(edible) {
			game.snake.Eat(edible)
			game.score += 1
			game.edibles[i] = nil
		}

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
		snake:   Snake{direction: Vector2{1, 0}, speed: 2},
		input:   InputHandler{},
		chef:    Chef{},
		edibles: make([]Edible, 10),
		ticker:  1}

	ebiten.Run(
		game.Main,
		int(GameWidth),
		int(GameHeight),
		2,
		"Go Go Snek!")

}
