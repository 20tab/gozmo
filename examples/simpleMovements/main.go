package main

import (
	"fmt"
	goz "github.com/20tab/gozmo"
)

// A behaviour allowing movements with arrow keys
type CrossMove struct {
	speed float32
}

func (cross *CrossMove) Start(gameObject *goz.GameObject) {
	// 5 units per second
	cross.speed = 5
}

// gameObject.Position is a Vector2 struct, 0 is x, 1 is y
func (cross *CrossMove) Update(gameObject *goz.GameObject) {
	if gameObject.GetKey(goz.KeyRight) {
		gameObject.Position[0] += cross.speed * gameObject.DeltaTime
	}

	if gameObject.GetKey(goz.KeyLeft) {
		gameObject.Position[0] -= cross.speed * gameObject.DeltaTime
	}

	if gameObject.GetKey(goz.KeyUp) {
		gameObject.Position[1] += cross.speed * gameObject.DeltaTime
	}

	if gameObject.GetKey(goz.KeyDown) {
		gameObject.Position[1] -= cross.speed * gameObject.DeltaTime
	}
}

func main() {

	window := goz.OpenWindow(1024, 768, "Gozmo")

	scene001 := window.NewScene()
	scene001.Name = "Scene 1"

	// load a texture into the scene
	scene001.NewTextureFilename("spyke_red", "assets/spyke_red.png")

	fmt.Println("scene", scene001.Name, "created")

	spyke := scene001.NewGameObject("Player001")
	// add a component by name
        spyke.AddComponentName("render", "Renderer", nil)
	// set component attribute with SetAttr
	spyke.SetAttr("render", "texture", "spyke_red")

	// and add another one by reference
	spyke.AddComponent("move_with_arrows", &CrossMove{})

	window.SetScene(scene001)
	window.Run()

}
