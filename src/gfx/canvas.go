package gfx

import (
	"fmt"
	"image"
	"log"
	"os"
	"sync"

	"image/color"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/op/paint"
)

const (
	WindowTitle  = "Wolf3D"
	WindowWidth  = 640
	WindowHeight = 480
)

// Canvas is a framebuffer that gets rendered into a window. Set pixels, draw frames, profit.
type Canvas struct {
	pendingFrame *image.NRGBA
	activeFrame  *image.NRGBA
	window       *app.Window
	sync.Mutex
}

func NewCanvas() (canvas *Canvas) {
	window := &app.Window{}
	window.Option(app.Title(WindowTitle))
	window.Option(app.Size(WindowWidth, WindowHeight))
	window.Option(app.MinSize(WindowWidth, WindowHeight))
	window.Option(app.MaxSize(WindowWidth, WindowHeight))

	canvas = &Canvas{
		window:       window,
		pendingFrame: image.NewNRGBA(image.Rect(0, 0, WindowWidth, WindowHeight)),
		activeFrame:  image.NewNRGBA(image.Rect(0, 0, WindowWidth, WindowHeight)),
	}

	go func() {
		err := canvas.render()
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	return canvas
}

// render renders the current image every frame event.
func (canvas *Canvas) render() error {
	var ops op.Ops
	for {
		switch e := canvas.window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			e.Source.Execute(op.InvalidateCmd{})
			// Reset the operations back to zero.
			ops.Reset()

			// Draw image every frame
			canvas.Lock()
			imageOp := paint.NewImageOp(canvas.activeFrame)
			canvas.Unlock()
			imageOp.Add(&ops)

			paint.PaintOp{}.Add(&ops)

			// Update the display.
			e.Frame(&ops)

			// log.Println("render")
		}
	}
}

// DrawPixel sets the pixel at (x, y) in the canvas to the specified color, but doesn't render it yet.
func (canvas *Canvas) DrawPixel(x, y int, c color.Color) error {
	if x < 0 || x >= WindowWidth || y < 0 || y >= WindowHeight {
		return fmt.Errorf("pixel %d:%d is out of bounds", x, y)
	}
	canvas.pendingFrame.Set(x, y, c)

	return nil
}

// PushFrame copies the pending frame to the active frame so it'll be rendered on the next draw
func (canvas *Canvas) PushFrame() {
	canvas.Lock()
	defer canvas.Unlock()
	*canvas.activeFrame = *canvas.pendingFrame
}

// ClearFrame resets the pending frame to an empty frame.
func (canvas *Canvas) ClearFrame() {
	canvas.pendingFrame = image.NewNRGBA(image.Rect(0, 0, WindowWidth, WindowHeight))
}
