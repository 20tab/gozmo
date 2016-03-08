package main

import (
	goz "github.com/20tab/gozmo"
)

func main() {
	window := goz.OpenWindow(1024, 576, "Map")

	scene := goz.NewScene("Level 1")

	spritesheet, err := scene.NewTextureFromFilename("tilesheet", "assets/tiles_spritesheet.png")
	if err != nil {
		panic(err)
	}

	backgroundMap := scene.NewGameObject("BackgroundMap")

	backgroundMap.AddComponent("tilemap", goz.NewTileMap(spritesheet))

	window.SetScene(scene)
	window.Run()
}
