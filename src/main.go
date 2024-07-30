package main

import (
	"image/color"
	"time"

	"github.com/rsheasby/wolf3d-go/gfx"
)

func main() {
	canvas := gfx.NewCanvas()
	go func() {
		canvas.DrawBox(0, 0, 639, 479, color.RGBA{255, 255, 255, 255})
		canvas.DrawLine(0, 0, 639, 479, color.RGBA{255, 0, 0, 255})
		canvas.PushFrame()
		time.Sleep(2 * time.Second)
		canvas.DrawLine(0, 479, 639, 0, color.RGBA{255, 0, 0, 255})
		canvas.PushFrame()
	}()
	canvas.Start()
}
