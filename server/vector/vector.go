package vector

type Vector2f struct {
	X, Y float64
}

func MakeVector2f(x, y float64) Vector2f {
	return Vector2f{x, y}
}

type Vector2 struct {
	X, Y int
}

func MakeVector2(x, y int) Vector2 {
	return Vector2{x, y}
}
