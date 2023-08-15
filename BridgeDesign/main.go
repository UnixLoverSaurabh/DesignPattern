package main

import "fmt"

/*
Circle, Rectangle
Raster, Vector

We will have n*m structs
RasterCircle, VectorCircle, RasterRectangle, VectorRectangle
*/

type Renderer interface {
	RenderCircle(radius int)
}

type vectorRenderer struct {
}

func (vr *vectorRenderer) RenderCircle(radius int) {
	fmt.Println("Drawing a circle of radius", radius)
}

type rasterRenderer struct {
	dpi int
}

func (rr *rasterRenderer) RenderCircle(radius int) {
	fmt.Println("Drawing a circle of radius", radius)
}

// ------------------------------------------------------------------------------------------------------------

type Circle struct {
	renderer Renderer
	radius   int
}

func NewCircle(renderer Renderer, radius int) *Circle {
	return &Circle{renderer: renderer, radius: radius}
}

func (circle *Circle) draw() {
	circle.renderer.RenderCircle(circle.radius)
}

func main() {
	vr := vectorRenderer{}
	circle := NewCircle(&vr, 5)
	circle.draw()
}
