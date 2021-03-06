package main

import (
	"github.com/hajimehoshi/ebiten"
	"reflect"
	"testing"
)

func TestGetTailOverBoundry(t *testing.T) {

	snek := Snake{}
	snek.length = 3
	snek.cursor = 0

	snek.positions = []SnakeSegment{
		{center: Vector2{1, 1}},
		{center: Vector2{}},
		{center: Vector2{}},
		{center: Vector2{3, 3}},
		{center: Vector2{2, 2}}}

	tail := snek.GetTail()

	expected := [3]Vector2{{1, 1}, {2, 2}, {3, 3}}
	result := [3]Vector2{<-tail, <-tail, <-tail}

	if result != expected {
		t.Error("snake tail was not expected, got: ", result)
	}

}

func TestGetTailSimple(t *testing.T) {

	snek := Snake{}
	snek.length = 3
	snek.cursor = 3
	snek.positions = []SnakeSegment{
		{center: Vector2{}},
		{center: Vector2{3, 3}},
		{center: Vector2{2, 2}},
		{center: Vector2{1, 1}},
		{center: Vector2{}}}

	tail := snek.GetTail()

	expected := [3]Vector2{{1, 1}, {2, 2}, {3, 3}}
	result := [3]Vector2{<-tail, <-tail, <-tail}

	if result != expected {
		t.Error("snake tail was not expected, got: ", result)
	}

}

func TestTailEndsAfterLength(t *testing.T) {

	snek := Snake{}
	snek.length = 3
	snek.cursor = 3

	snek.positions = []SnakeSegment{{}, {}, {}, {}, {}}

	tail := snek.GetTail()

	<-tail
	<-tail
	<-tail

	_, elements_left := <-tail

	if elements_left {
		t.Error("snake tail did not end when we wanted")
	}

}

func TestTailSkip(t *testing.T) {

	tail := make(Tail, 5)

	tail <- Vector2{1, 1}
	tail <- Vector2{2, 2}
	tail <- Vector2{3, 3}
	tail <- Vector2{4, 4}
	tail <- Vector2{5, 5}

	tail.Skip(3)

	result := <-tail

	if (result != Vector2{4, 4}) {
		t.Error("tail skip failed")
	}

}

func TestMyTailSkip(t *testing.T) {

	tail := make(MyTail, 5)

	myVec := Vector2{4, 4}
	tail <- &Vector2{1, 1}
	tail <- &Vector2{2, 2}
	tail <- &Vector2{3, 3}
	tail <- &myVec
	tail <- &Vector2{5, 5}

	tail.Skip(3)

	result := <-tail

	if result != &myVec {
		t.Error("tail skip failed")
	}

}

type Snack struct{}

func (s Snack) amount() int {
	return 1
}

func (s Snack) avatar() *ebiten.Image {
	return nil
}

func (s Snack) position() Vector2 {
	return Vector2{0, 0}
}

func (s Snack) Collider() Rect {
	return Rect{}
}

func TestEatGrow(t *testing.T) {

	snake := Snake{}

	snake.positions = []SnakeSegment{
		{center: Vector2{1, 1}},
		{center: Vector2{2, 2}},
		{center: Vector2{3, 3}},
		{center: Vector2{4, 4}},
		{center: Vector2{5, 5}}}

	snake.cursor = 3
	snake.length = 3

	snake.Eat(Snack{})

	expected := []SnakeSegment{
		{center: Vector2{2, 2}},
		{center: Vector2{2, 2}},
		{center: Vector2{3, 3}},
		{center: Vector2{4, 4}},
		{center: Vector2{5, 5}}}

	if !reflect.DeepEqual(expected, snake.positions) {
		t.Error("eating did not set the new tail segment to the last tail segment's length")
	}

}
