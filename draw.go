// here we do our drawing of said snek

package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"math"
	"math/rand"
)

type Drawable interface {
	//this the image for the edible
	avatar() *ebiten.Image
	//this is used for drawing
	position() Vector2
}

// does the level animation, where the snake flips out

func DoLevelAnimation(game *GameState, screen *ebiten.Image) {

	//we are responsible for all drawing in the coroutine- Update
	//is not handling that anymore

	DrawBackground(screen)

	//now we handle the shrink

	percentageRemaining := float64(game.tick-game.coroutineStart) / 180.0

	if percentageRemaining > 1 {
		percentageRemaining = 1
	}

	//find the closest value of 2, using the ticks to completion as a %

	//first calculate the difference of max and start using the percentage,
	//when there are all ticks reminaing, this is 0 ( the snake is the biggest ),
	//and when there are no ticks, it is 0 ( snake is smol )
	sizeDiffAsFloat := (maxSize - startSize) * percentageRemaining

	//now take the floor of the difference
	sizeDiffAsInt := int(math.Floor(sizeDiffAsFloat))

	//now find the nearest power of 2 of the difference
	sizeDiff := (2 - sizeDiffAsInt%2) + sizeDiffAsInt

	//now actually subtract it from the max size, getting our new size,
	//divisible by 2 and between max size and start size
	newSize := maxSize - sizeDiff

	//sometimes we can go over, depending on how accurate our frame
	//emulation is via ticks

	if newSize < startSize {
		newSize = startSize
	}

	i := 0

	shrunkImage := game.snake.GetImageForSize(newSize)

	//new color for the shrunken image

	randomColor := color.NRGBA{
		uint8(rand.Intn(255)),
		uint8(rand.Intn(255)),
		uint8(rand.Intn(255)),
		255}

	shrunkImage.Fill(randomColor)

	//do the "wobble animation"
	for segment := range game.snake.GetTail() {

		i++
		//find the value of cos/sin at game.tick + i(offset) multiplied by a radian
		//this creates a circular motion in which the segments appear to wobble

		diffX := math.Cos(float64(game.tick+i)*(math.Pi/24)) * (5 - (4 * percentageRemaining))
		diffY := math.Sin(float64(game.tick+i)*(math.Pi/24)) * (5 - (4 * percentageRemaining))

		newPosition := Vector2{
			diffX + segment.x,
			diffY + segment.y}

		//we create a new segment here, really only using the position of
		//the old segment

		newSegment := SnakeSegment{
			&game.snake,
			newPosition,
			float64(newSize),
			shrunkImage}

		//and draw the segment
		DrawObject(newSegment, screen)

	}

}

//draw does not take a pointer to state, so we do not modifiy it

//i had this throw an error for the parent to collect,
//but i am no longer doing that. unsure if i want to remove it

func Draw(game GameState, screen *ebiten.Image) error {

	DrawBackground(screen)

	//this might be a real data object someday
	drawable := make([]Drawable, len(game.edibles))

	for i, edible := range game.edibles {
		drawable[i] = edible
		DrawObject(drawable[i], screen)
	}

	for drawObject := range game.snake.GetDrawTail() {
		DrawObject(drawObject, screen)
	}

	return nil

}

func DrawBackground(screen *ebiten.Image) {
	//background draw
	screen.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff})
}

func DrawObject(drawable Drawable, screen *ebiten.Image) {

	if drawable == nil {
		return
	}

	if drawable.avatar() == nil {
		return
	}

	position := drawable.position()

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(position.x, position.y)

	screen.DrawImage(drawable.avatar(), opts)

}
