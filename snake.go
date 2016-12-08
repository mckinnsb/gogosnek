package main

//start length of snek, don't use this for calculations,
//just starting snek
const startLength = 20

//this is only used in creating the storage array,
//do not use it for calculations!
const maxLength = 200

//struct for snek
type Snake struct {
	position, direction Vector2
	positions           []Vector2
	cursor              int

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

//eat an edible, make snake strong
func (snake *Snake) Eat(eatme Edible) {
	snake.length += eatme.amount()
}

/**
maybe if we do resizable tails? will have to deal with circular structure
func (snake *Snake) ResizeTail() {
}
**/

//start snek
func (snake *Snake) Start(position Vector2) {

	snake.length = startLength

	snake.positions = make([]Vector2, maxLength)
	snake.position = position
	snake.positions[0] = snake.position

	snake.cursor = 0

}

//update snek, move him, add new position to positions list
//and return
func (snake *Snake) Update() {
	newPosition := snake.direction.Multiply(snake.speed)
	snake.position = snake.position.Add(newPosition)
	snake.AddPosition(snake.position)
	return
}
