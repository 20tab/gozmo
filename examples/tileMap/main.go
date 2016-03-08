package main

import (
	goz "github.com/20tab/gozmo"
)

var cameraSpeed float32 = 5

type CameraMover struct{}

func (mover *CameraMover) Start(gameObject *goz.GameObject) {}
func (mover *CameraMover) Update(gameObject *goz.GameObject) {
	var x float32 = 0
	var y float32 = 0
	if goz.IsTrue(gameObject.GetAttr("kbd", "A")) {
		x = -cameraSpeed * gameObject.DeltaTime
	}

	if goz.IsTrue(gameObject.GetAttr("kbd", "D")) {
		x = cameraSpeed * gameObject.DeltaTime
	}

	if goz.IsTrue(gameObject.GetAttr("kbd", "W")) {
		y = cameraSpeed * gameObject.DeltaTime
	}

	if goz.IsTrue(gameObject.GetAttr("kbd", "S")) {
		y = -cameraSpeed * gameObject.DeltaTime
	}

	gameObject.AddPosition(x, y)
}

func main() {
	window := goz.OpenWindow(1024, 576, "Map")

	scene := goz.NewScene("Level 1")

	camera := scene.NewGameObject("camera")
	camera.AddComponentByName("camera", "Camera", nil)
	camera.AddComponentByName("kbd", "Keyboard", nil)
	camera.AddComponent("cross", &CameraMover{})

	spritesheet, err := scene.NewTextureFromFilename("tilesheet", "assets/tiles_spritesheet.png")
	if err != nil {
		panic(err)
	}
	spritesheet.SetCols(12)
	spritesheet.SetRows(11)

	backgroundMap := scene.NewGameObject("BackgroundMap")

	backgroundMap.AddComponent("tilemap", goz.NewTileMapFromCSVFilename("assets/map001.csv", spritesheet))

	window.SetScene(scene)
	window.Run()
}
