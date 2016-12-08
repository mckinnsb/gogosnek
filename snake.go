package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

//start length of snek, don't use this for calculations,
//just starting snek
const startLength = 20

//this is only used in creating the storage array,
//do not use it for calculations!
const maxLength = 200

//struct for snek
type Snake struct {

	//the image, this thing is heavy, so, careful
	//when you deref
	avatar *ebiten.Image

	//the position and direction of snek
	position, direction Vector2

	//the positions of snek
	positions []Vector2

	//the cursor of snek, or where his head position is in positions,
	//we form his tail from length of positions, the rest is considered
	//garbage
	cursor int

	//speed in pixels/update
	speed float64

	//this is not a length in pixels, but rather,
	//a length in positions ( can be thought of as length in pixels = speed*length )
	length int
}

func (snake *Snake) AddPosition(position Vector2) {

	snake.cursor++
	if snake.cursor >= cap(snake.positions) {
		snake.cursor = 0
	}

	snake.positions[snake.cursor] = position

}

//this returns a channel that is an "enumerator" over
//the past positions recorded by the snake. we return
//a number of positions equal to the snakes length,
//which is not a length in pixels

func (snake *Snake) GetTail() <-chan Vector2 {

	//we buffer so we don't have to make a goroutine
	out := make(chan Vector2, snake.length)

	//this is also "diff", and is used for calcuating
	//the tail end
	start := snake.cursor - snake.length + 1

	var startIndex int

	if start < 0 {
		startIndex = 0
	} else {
		startIndex = start
	}

	//ill be honest, i don't like this about go; it excludes
	//the end of the range, so we have to add one here

	beginning := snake.positions[startIndex : snake.cursor+1]

	for i := len(beginning) - 1; i >= 0; i-- {
		out <- beginning[i]
	}

	if start < 0 {

		//the max length plus the difference of the cursor and the
		//length is the start position of the remainder of the tail

		endIndex := len(snake.positions) + start
		ending := snake.positions[endIndex:]

		for i := len(ending) - 1; i >= 0; i-- {
			out <- ending[i]
		}

	}

	close(out)
	return out

}

func (snake *Snake) IsColliding(edible Edible) bool {
	me := Rect{snake.position, snake.position.Add(Vector2{8, 8})}
	other := Rect{
		edible.position(),
		edible.position().Add(edible.size())}

	return me.CollidesWith(other)

}

//eat an edible, make snake strong
func (snake *Snake) Eat(eatme Edible) {
	snake.length += eatme.amount()

	if snake.length > maxLength {
		snake.length = maxLength
	}
}

//start snek,
//called by game init function

//this might be implemented by an interface,
//if i had more mobile objects in the game,
//right now its just snek

func (snake *Snake) Start(position Vector2) {

	snake.length = startLength

	snake.positions = make([]Vector2, maxLength)
	snake.position = position
	snake.positions[0] = snake.position

	snake.cursor = 0

	//we create this just once, because it is a heavy struct
	snake.avatar, _ = ebiten.NewImage(8, 8, ebiten.FilterNearest)

	//and we do this just once, because it's fairly expensive
	snake.avatar.Fill(color.White)

}

//update snek, move him, add new position to positions list
//and return

func (snake *Snake) Update() {
	newPosition := snake.direction.Multiply(snake.speed)
	snake.position = snake.position.Add(newPosition).ClampToWindow()
	snake.AddPosition(snake.position)
	return
}
