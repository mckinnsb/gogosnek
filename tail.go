package main

type Tail chan Vector2
type MyTail chan *Vector2
type DrawTail chan Drawable

//this function skips num entries and returns itself
//returning itself is more a convenience than
//anything else - channel is a struct around
//a pointer, so the value will always point
//to the same tail ( we could skip and not use the return value )

func (tail Tail) Skip(num int) Tail {

	for i := 0; i < num; i++ {
		_, result := <-tail
		if !result {
			break
		}
	}

	return tail

}

//similar to the function above,
//there's some repeated code, but avoiding it
//is more of a pain ( we would have to convert
//this to an empty interface channel )

func (tail MyTail) Skip(num int) MyTail {

	for i := 0; i < num; i++ {
		_, result := <-tail
		if !result {
			break
		}
	}

	return tail

}
