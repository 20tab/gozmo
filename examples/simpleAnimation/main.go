package main

import (
    goz "github.com/20tab/gozmo"
    "fmt"
)


func main() {

    window := goz.OpenWindow(1024, 768, "Gozmo")

    scene001 := window.NewSceneFilename("assets/scene.json")

    fmt.Println("scene", scene001.Name, "loaded")

    window.SetScene(scene001)
    window.Run()

}
