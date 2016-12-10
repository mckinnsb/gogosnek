package main

import "github.com/hajimehoshi/ebiten"

type SnakeSegment struct {
	snake  *Snake
	center Vector2
	size   float64
	image  *ebiten.Image
}

func (seg SnakeSegment) avatar() *ebiten.Image {
	return seg.image
}

func (seg SnakeSegment) position() Vector2 {
	return seg.center.Add(Vector2{-seg.size / 2, -seg.size / 2})
}
