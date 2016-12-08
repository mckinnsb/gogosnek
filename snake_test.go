package main

import "testing"

func TestGetTailOverBoundry(t *testing.T) {

	snek := Snake{}
	snek.length = 3
	snek.cursor = 0

	snek.positions = []Vector2{{1, 1}, {}, {}, {3, 3}, {2, 2}}

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

	snek.positions = []Vector2{{}, {3, 3}, {2, 2}, {1, 1}, {}}

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

	snek.positions = []Vector2{{}, {}, {}, {}, {}}

	tail := snek.GetTail()

	<-tail
	<-tail
	<-tail

	_, elements_left := <-tail

	if elements_left {
		t.Error("snake tail did not end when we wanted")
	}

}
