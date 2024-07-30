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
	if startX < 0 || endX >= CanvasWidth || startY < 0 || endY >= CanvasHeight {
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

func (canvas *Canvas) DrawLine(startX, startY, endX, endY int, c color.Color) error {
	// Make sure the line is within bounds
	if startX < 0 || endX >= CanvasWidth || startY < 0 || endY >= CanvasHeight {
		return fmt.Errorf("line %d:%d - %d:%d is out of bounds", startX, startY, endX, endY)
	}

	// Figure out how many pixels we need to set
	xDistance := endX - startX
	xStepSize := 1.0
	yDistance := endY - startY
	yStepSize := 1.0

	// Figure out how far we need to move between each pixel
	var pixelCount int
	if xDistance > yDistance {
		pixelCount = xDistance
		yStepSize = float64(yDistance) / float64(pixelCount)
	} else {
		pixelCount = yDistance
		xStepSize = float64(xDistance) / float64(pixelCount)
	}

	// Set all the pixels
	for i := 0; i < pixelCount; i++ {
		err := canvas.DrawPixel(startX+int(float64(i)*xStepSize), startY+int(float64(i)*yStepSize), c)
		if err != nil {
			return err
		}
	}

	return nil
}
