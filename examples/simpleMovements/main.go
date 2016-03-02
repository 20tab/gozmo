package main

import (
	"fmt"
	goz "github.com/20tab/gozmo"
)

// A behaviour allowing movements with arrow keys
type CrossMove struct {
	kbd   *goz.Keyboard
	speed float32
}

func (cross *CrossMove) Start(gameObject *goz.GameObject) {
	// 5 units per second
	cross.speed = 5
}

// gameObject.Position is a Vector2 struct, 0 is x, 1 is y
func (cross *CrossMove) Update(gameObject *goz.GameObject) {
	if cross.kbd.GetKey(goz.KeyRight) {
		gameObject.Position[0] += cross.speed * gameObject.DeltaTime
	}

	if cross.kbd.GetKey(goz.KeyLeft) {
		gameObject.Position[0] -= cross.speed * gameObject.DeltaTime
	}

	if cross.kbd.GetKey(goz.KeyUp) {
		gameObject.Position[1] += cross.speed * gameObject.DeltaTime
	}

	if cross.kbd.GetKey(goz.KeyDown) {
		gameObject.Position[1] -= cross.speed * gameObject.DeltaTime
	}
}

func main() {

	window := goz.OpenWindow(1024, 768, "Gozmo")

	scene001 := window.NewScene("Scene 1")

	// load a texture into the scene
	scene001.NewTextureFilename("spyke_red", "assets/spyke_red.png")

	fmt.Println("scene", scene001.Name, "created")

	spyke := scene001.NewGameObject("Player001")
	// add a component by name
	spyke.AddComponentName("render", "Renderer", nil)
	// set component attribute with SetAttr
	spyke.SetAttr("render", "texture", "spyke_red")

	keyboard := goz.NewKeyboard()
	spyke.AddComponent("kbd", keyboard)

	// and add another one by reference
	spyke.AddComponent("move_with_arrows", &CrossMove{kbd: keyboard})

	window.SetScene(scene001)
	window.Run()

}
