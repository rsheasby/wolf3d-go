package gfx

import (
	"fmt"
	"image"
	"log"
	"sync"

	"image/color"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

const (
	WindowTitle  = "Wolf3D"
	CanvasWidth  = 640
	CanvasHeight = 480
)

// Canvas is a framebuffer that gets rendered into a window. Set pixels, draw frames, profit.
type Canvas struct {
	buffer       screen.Buffer
	texture      screen.Texture
	window       screen.Window
	keyEvents    chan key.Event
	mouseEvents  chan mouse.Event
	windowWidth  int
	windowHeight int
	sync.Mutex
}

func NewCanvas() (canvas *Canvas) {
	canvas = &Canvas{
		keyEvents:    make(chan key.Event, 1),
		mouseEvents:  make(chan mouse.Event, 1),
		windowWidth:  CanvasWidth * 2,
		windowHeight: CanvasHeight * 2,
	}
	canvas.Lock()

	return canvas
}

// Start creates the window and buffer, and actually starts the rendering process.
// This blocks forever.
func (canvas *Canvas) Start() {
	driver.Main(func(s screen.Screen) {
		var err error
		canvas.window, err = s.NewWindow(&screen.NewWindowOptions{
			Width:  canvas.windowWidth,
			Height: canvas.windowHeight,
			Title:  WindowTitle,
		})
		if err != nil {
			log.Fatalf("Error opening window: %v", err)
			return
		}
		defer canvas.window.Release()

		canvas.texture, err = s.NewTexture(image.Point{X: CanvasWidth, Y: CanvasHeight})
		if err != nil {
			log.Fatalf("Unable to create screen texture: %v", err)
		}

		canvas.buffer, err = s.NewBuffer(image.Point{X: CanvasWidth, Y: CanvasHeight})
		if err != nil {
			log.Fatalf("Unable to create frame buffer: %v", err)
		}
		canvas.Unlock()

		for {
			switch e := canvas.window.NextEvent().(type) {
			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					return
				}
			case paint.Event:
				canvas.paint()
			case size.Event:
				canvas.handleResize(e)
			case key.Event:
				canvas.keyEvents <- e
			case mouse.Event:
				// canvas.mouseEvents <- e
			case error:
				log.Printf("Error event: %v", e)
			default:
				log.Printf("Unhandled event: %v", e)
			}
		}
	})
}

func (canvas *Canvas) handleResize(e size.Event) {
	canvas.windowWidth = e.WidthPx
	canvas.windowHeight = e.HeightPx
}

func (canvas *Canvas) KeyEvents() (events <-chan key.Event) {
	return canvas.keyEvents
}

func (canvas *Canvas) MouseEvents() (events <-chan mouse.Event) {
	return canvas.mouseEvents
}

// DrawPixel sets the pixel at (x, y) in the canvas to the specified color, but doesn't render it yet.
func (canvas *Canvas) DrawPixel(x, y int, c color.Color) error {
	canvas.Lock()
	defer canvas.Unlock()

	if x < 0 || x >= CanvasWidth || y < 0 || y >= CanvasHeight {
		return fmt.Errorf("pixel %d:%d is out of bounds", x, y)
	}
	canvas.buffer.RGBA().Set(x, y, c)

	return nil
}

// PushFrame copies the pending frame to the active frame so it'll be rendered on the next draw
func (canvas *Canvas) PushFrame() {
	canvas.paint()
}

func (canvas *Canvas) paint() {
	canvas.Lock()
	defer canvas.Unlock()

	// log.Println("Painting...")
	canvas.texture.Upload(image.Point{0, 0}, canvas.buffer, canvas.buffer.Bounds())
	canvas.window.Scale(image.Rect(0, 0, canvas.windowWidth, canvas.windowHeight), canvas.texture, canvas.texture.Bounds(), screen.Over, nil)
	canvas.window.Publish()
}
