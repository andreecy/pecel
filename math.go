package pecel

type Vec2 struct {
	X float64
	Y float64
}

func NewVec2(x, y int) Vec2 {
	return Vec2{float64(x), float64(y)}
}
