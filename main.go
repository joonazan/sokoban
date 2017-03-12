package main

import (
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/joonazan/closedgl"
)

func main() {
	window := closedgl.NewWindow(640, 640, "Sokoban")
	window.SetKeyCallback(keyPressed)

	LataaKuvat()

	closedgl.RunInWindow(render, window)
}

func render(dt float64) {
	Piirrä()
}

func keyPressed(_ *glfw.Window, k glfw.Key, _ int, action glfw.Action, _ glfw.ModifierKey) {
	if action == glfw.Press {
		switch k {
		case glfw.KeyLeft:
			Liiku(-1, 0)
		case glfw.KeyRight:
			Liiku(1, 0)
		case glfw.KeyUp:
			Liiku(0, -1)
		case glfw.KeyDown:
			Liiku(0, 1)
		}
	}
}
