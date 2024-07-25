package main

import (
	"image/color"

	"gioui.org/app"
	"github.com/rsheasby/wolf3d-go/gfx"
)

func main() {
	canvas := gfx.NewCanvas()
	canvas.DrawBox(100, 100, 200, 200, color.NRGBA{255, 0, 0, 255})
	canvas.PushFrame()
	app.Main()
}
