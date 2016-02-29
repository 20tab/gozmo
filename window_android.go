// +build android

package gozmo

import (
	"golang.org/x/mobile/app"
)

type Key int

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
