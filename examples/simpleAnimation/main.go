package main

import (
	"fmt"
	goz "github.com/20tab/gozmo"
)

func main() {

	window := goz.OpenWindow(1024, 768, "Gozmo")

	scene001 := goz.NewSceneFromFilename("assets/scene.json")

	fmt.Println("scene", scene001.Name, "loaded")

	window.SetScene(scene001)
	window.Run()

}
