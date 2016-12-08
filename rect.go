package main

//a simple rectangle class

//start and end are the top left and bottom right corners of the
//rectangle

type Rect struct {
	start, end Vector2
}

func (rect Rect) CollidesWith(other Rect) bool {

	switch {

	//one rect or the other rect is to the right or left of the other
	case rect.start.x > other.end.x || other.start.x > rect.end.x:
		return false
	//one rect or the other rect is above or below the other
	case rect.start.y > other.end.y || other.start.y > rect.end.y:
		return false
	default:
		//otherwise, its a collision
		return true

	}

}
