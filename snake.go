package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

//this is only used in creating the storage array,
//do not use it for calculations!
const maxLength = 400

//how big can the snake get before we resize the level?
//(virtual because it is multiplied by the scale factor given to ebiten)

const maxSize = 32
const sizeIncrease = 4
const startSize = 8

//start length of snek, don't use this for calculations,
//just starting snek

const startLength = 20

//struct for snek
// the hero of the day
type Snake struct {

	//are you advanced snek?
	advanced bool

	//the image, this thing is heavy, so, careful
	//when you deref
	avatar *ebiten.Image

	//image buddy
	image *SnakeImage

	//the position and direction of snek
	position, direction Vector2

	//the positions of snek
	positions []SnakeSegment

	//the cursor of snek, or where his head position is in positions,
	//we form his tail from length of positions, the rest is considered
	//garbage
	cursor int

	//speed in pixels/update
	speed float64

	//this is not a length in pixels, but rather,
	//a length in positions ( can be thought of as length in pixels = speed*length )
	length int

	//this is the width of the snek
	size float64
}

func (snake *Snake) AddPosition(position Vector2) {

	snake.cursor++
	if snake.cursor >= cap(snake.positions) {
		snake.cursor = 0
	}

	snake.positions[snake.cursor] = SnakeSegment{snake, position, snake.size, snake.avatar}

}

func (snake *Snake) Collider() Rect {
	return snake.GetColliderAt(snake.position)
}

//eat a thing you can eat, make snake strong

func (snake *Snake) Eat(eatme Edible) {

	oldLength := snake.length

	if snake.advanced && snake.size < maxSize {

		snake.size += sizeIncrease
		snake.avatar = snake.GetImageForSize(int(snake.size))

	}

	snake.length += eatme.amount()

	if snake.length > maxLength {
		snake.length = maxLength
	}

	if snake.length != oldLength {

		//tail will now be longer, which we can
		//now go through, and "skip" our old length #-1
		//of segments, and set them all to the old length-1'th
		//position

		myTail := snake.GetMyTail().Skip(oldLength - 1)
		lastSegment := <-myTail

		for segment := range myTail {
			segment.x = lastSegment.x
			segment.y = lastSegment.y
		}

	}

}

//are you eating your own tail, snake?

func (snake *Snake) EatingTail() bool {

	remainder := snake.GetTail().Skip(int(snake.size) + 1)
	eatingSelf := false

	head := snake.Collider()

	for suspect := range remainder {

		rect := snake.GetColliderAt(suspect)

		if rect.CollidesWith(head) {
			eatingSelf = true
			break
		}

	}

	return eatingSelf

}

//this returns a channel that is an "enumerator" that
//decorates the tail positions as drawable objects

//this version of the class has no skip function
//because you must draw all of the tail, or at
//least im not going to make that easy for me to
//mess myself up on

func (snake *Snake) GetDrawTail() DrawTail {

	out := make(DrawTail, snake.length)

	beginning, ending := snake.GetSegments()

	for i := len(beginning) - 1; i >= 0; i-- {
		out <- beginning[i]
	}

	for i := len(ending) - 1; i >= 0; i-- {
		out <- ending[i]
	}

	close(out)
	return out

}

//this resturns an image for the snake avatar at its current

func (snake *Snake) GetImageForSize(size int) *ebiten.Image {

	image := snake.image.images[int(size)]

	//this is inefficient, but should only be called
	//once for non-advanced snek ( it does not change size )

	if !snake.advanced {
		fmt.Println("im not advanced")
		image.Fill(color.White)
	}

	return image

}

//this returns a channel that is an "enumerator" over
//the past positions - but as pointers
//
//this is for modifying the tail, removing positions,
//etc

func (snake *Snake) GetMyTail() MyTail {

	//we buffer so we don't have to make a goroutine
	out := make(MyTail, snake.length)

	beginning, ending := snake.GetSegments()

	for i := len(beginning) - 1; i >= 0; i-- {
		out <- &(beginning[i].center)
	}

	for i := len(ending) - 1; i >= 0; i-- {
		out <- &(ending[i].center)
	}

	close(out)
	return out

}

//this function returns slices representing the two
//segments of the tail, the beginning and the end
//keep in mind the end *might* be a zero slice,
//the beginning will always be at least cap1

func (snake *Snake) GetSegments() (beginning, end []SnakeSegment) {

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

	beginning = snake.positions[startIndex : snake.cursor+1]

	if start < 0 {

		//the max length plus the difference of the cursor and the
		//length is the start position of the remainder of the tail

		endIndex := len(snake.positions) + start
		end = snake.positions[endIndex:]

	}

	return beginning, end

}

func (snake *Snake) GetColliderAt(position Vector2) Rect {
	offset := position.Add(Vector2{-snake.size / 2, -snake.size / 2})
	return Rect{offset, offset.Add(Vector2{snake.size, snake.size})}
}

//this returns a channel that is an "enumerator" over
//the past positions recorded by the snake. we return
//a number of positions equal to the snakes length,
//which is not a length in pixels

func (snake *Snake) GetTail() Tail {

	//we buffer so we don't have to make a goroutine
	out := make(Tail, snake.length)

	beginning, ending := snake.GetSegments()

	for i := len(beginning) - 1; i >= 0; i-- {
		out <- beginning[i].center
	}

	for i := len(ending) - 1; i >= 0; i-- {
		out <- ending[i].center
	}

	close(out)
	return out

}

func (snake *Snake) IsColliding(edible Edible) bool {
	me := snake.Collider()
	other := edible.Collider()
	return me.CollidesWith(other)

}

func (snake *Snake) ResetToStartSize() {

	//temporary, for now

	snake.size = startSize

	//reset snake images to default
	snake.image.Start()

	snake.avatar = snake.GetImageForSize(int(snake.size))

	//we dont use tail here because we don't care about order,
	//everything gets reset

	for i, _ := range snake.positions {
		snake.positions[i].size = startSize
		snake.positions[i].image = snake.avatar
	}

}

//start snek,
//called by game init function

//this might be implemented by an interface,
//if i had more mobile objects in the game,
//right now its just snek

func (snake *Snake) Start(position Vector2) {

	snake.length = startLength

	snake.positions = make([]SnakeSegment, maxLength)
	snake.position = position

	snake.cursor = 0

	snake.size = startSize

	//we create our snake image helper here, who creates
	//all the iamges for sizes ahead of time

	snake.image = &SnakeImage{}
	snake.image.Start()

	//we create this just once per scale, because it is a heavy struct
	//we are going to move this into a new class soon

	snake.avatar = snake.GetImageForSize(int(snake.size))
	snake.positions[0] = SnakeSegment{snake, snake.position, snake.size, snake.avatar}

}

//update snek, move him, add new position to positions list
//and return

func (snake *Snake) Update() {
	newPosition := snake.direction.Multiply(snake.speed)
	snake.position = snake.position.Add(newPosition).ClampToWindow()
	snake.AddPosition(snake.position)
	return
}
