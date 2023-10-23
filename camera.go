package pecel

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Camera struct {
	Width, Height float64
	Viewport      Vec2
	Position      Vec2
	Rotation      int
	ZoomFactor    int
}

func CreateCamera(width, height float64) *Camera {
	camera := &Camera{
		Width:    width,
		Height:   height,
		Viewport: Vec2{width, height},
	}
	return camera
}

func (c *Camera) String() string {
	return fmt.Sprintf("T: %.1f, R: %d, S: %d", c.Position, c.Rotation, c.ZoomFactor)
}

func (c *Camera) viewportCenter() Vec2 {
	return Vec2{
		c.Viewport.X * 0.5,
		c.Viewport.Y * 0.5,
	}
}

func (c *Camera) worldMatrix() ebiten.GeoM {
	m := ebiten.GeoM{}
	m.Translate(-c.Position.X, -c.Position.Y)
	// scale and rotate around center of image / screen
	// m.Translate(-c.viewportCenter().X, -c.viewportCenter().Y)
	m.Scale(
		math.Pow(1.01, float64(c.ZoomFactor)),
		math.Pow(1.01, float64(c.ZoomFactor)),
	)
	m.Rotate(float64(c.Rotation) * 2 * math.Pi / 360)
	m.Translate(c.viewportCenter().X, c.viewportCenter().Y)
	return m
}

func (c *Camera) Render(world, screen *ebiten.Image) {
	screen.DrawImage(world, &ebiten.DrawImageOptions{
		GeoM: c.worldMatrix(),
	})
}

func (c *Camera) ScreenToWorld(pos Vec2) Vec2 {
	inverseMatrix := c.worldMatrix()
	if inverseMatrix.IsInvertible() {
		inverseMatrix.Invert()
		x, y := inverseMatrix.Apply(float64(pos.X), float64(pos.Y))
		return Vec2{x, y}
	} else {
		return Vec2{0, 0}
	}
}

func (c *Camera) Reset() {
	c.Position.X = 0
	c.Position.Y = 0
	c.Rotation = 0
	c.ZoomFactor = 0
}
