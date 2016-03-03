package main

import (
	"fmt"
	goz "github.com/20tab/gozmo"
)

type ShowPos struct {
	mouse *goz.Mouse
}

func (showPos *ShowPos) Start(gameObject *goz.GameObject) {}
func (showPos *ShowPos) Update(gameObject *goz.GameObject) {
	fmt.Println("x =", showPos.mouse.X())
}

func main() {
	window := goz.OpenWindow(1024, 576, "Mouse")

	scene := window.NewScene("Mouse Manager")

	cursor := scene.NewGameObject("Cursor")

	mouse := goz.NewMouse()

	cursor.AddComponent("mouse", mouse)

	cursor.AddComponent("showPos", &ShowPos{mouse: mouse})

	window.SetScene(scene)
	window.Run()
}
