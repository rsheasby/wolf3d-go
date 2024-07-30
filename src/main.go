package main

import (
	"github.com/rsheasby/wolf3d-go/engine"
	"github.com/rsheasby/wolf3d-go/gfx"
)

func main() {
	canvas := gfx.NewCanvas()
	go engine.Run("map1.w3d", canvas)
	canvas.Start()
}
