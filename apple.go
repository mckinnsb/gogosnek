package main

type Edible interface {
	amount() int
}

type Apple struct {
	nutrition int
}

func (apple *Apple) amount() int {
	return apple.nutrition
}
