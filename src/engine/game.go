package engine

import (
	"image/color"

	"github.com/rsheasby/wolf3d-go/gfx"
)

// Run creates a fresh game state, and starts the game engine. Blocks forever.
func Run(mapFilename string, canvas *gfx.Canvas) {
	gameMap := ReadMap(mapFilename)
	for {
		render(gameMap, canvas)
	}
}

func render(gameMap Map, canvas *gfx.Canvas) {
	renderTopDown(gameMap, canvas)
}

func renderTopDown(gameMap Map, canvas *gfx.Canvas) {
	canvas.DrawBox(0, 0, gfx.CanvasWidth, gfx.CanvasHeight, color.NRGBA{255, 255, 255, 255})
	mapWidth, mapHeight := gameMap.Dimensions()
	tileWidth := gfx.CanvasWidth / mapWidth
	tileHeight := gfx.CanvasHeight / mapHeight
	for y := 0; y < len(gameMap); y++ {
		for x := 0; x < len(gameMap[y]); x++ {
			if gameMap[y][x] == 0 {
				continue
			}
			boxPositionX := x * tileWidth
			boxPositionY := y * tileHeight
			canvas.DrawBox(boxPositionX,
				boxPositionY,
				boxPositionX+tileWidth,
				boxPositionY+tileHeight,
				color.RGBA{50, 50, 50, 255},
			)
		}
	}
	canvas.PushFrame()
}
