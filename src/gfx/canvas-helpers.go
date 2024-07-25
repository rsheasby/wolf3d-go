package gfx

import (
	"fmt"
	"image/color"
)

func (canvas *Canvas) DrawBox(startX, startY, endX, endY int, c color.Color) error {
	// Flip the order if necessary
	if startX > endX {
		startX, endX = endX, startX
	}
	if startY > endY {
		startY, endY = endY, startY
	}

	// Make sure the box is within bounds
	if startX < 0 || endX >= WindowWidth || startY < 0 || endY >= WindowHeight {
		return fmt.Errorf("box %d:%d - %d:%d is out of bounds", startX, startY, endX, endY)
	}

	// Set the pixels
	for x := startX; x <= endX; x++ {
		for y := startY; y < endY; y++ {
			err := canvas.DrawPixel(x, y, c)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
