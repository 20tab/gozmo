// +build android

package gozmo

import (
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/gl"
)

type Key int

// The Window type interfaces with the display hardware using OpenGL.
// Coordinates are 0, 0 at screen center.
type Window struct {
	width        int32
	height       int32
	title        string
	scenes       []*Scene
	currentScene *Scene
}

func (window *Window) getKey(kc Key) bool {
	return false
}

var KeyboardAttr map[string]Key = map[string]Key{}

func OpenWindow(width int32, height int32, title string) *Window {
	window := Window{width: width, height: height, title: title}
	return &window
}

func (window *Window) Run() {
	app.Main(func(a app.App) {
		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
			case lifecycle.Event:
				switch e.Crosses(lifecycle.StageVisible) {
				case lifecycle.CrossOn:
					glctx, _ = e.DrawContext.(gl.Context)
					GLInit(40, 20)
				}
			case paint.Event:
				window.redraw()
				a.Publish()
				a.Send(paint.Event{})
			}
		}
	})
}
