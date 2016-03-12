package main

import (
	"fmt"
	goz "github.com/20tab/gozmo"
)

type CollisionCheck struct{}

func (check *CollisionCheck) Start(gameObject *goz.GameObject) {}
func (check *CollisionCheck) Update(gameObject *goz.GameObject) {
	var speed float32 = 3
	if goz.IsTrue(gameObject.GetAttr("kbd", "Right")) {
		gameObject.Position[0] += speed * gameObject.DeltaTime
	}
	if goz.IsTrue(gameObject.GetAttr("kbd", "Left")) {
		gameObject.Position[0] -= speed * gameObject.DeltaTime
	}
	if goz.IsTrue(gameObject.GetAttr("kbd", "Up")) {
		gameObject.Position[1] += speed * gameObject.DeltaTime
	}
	if goz.IsTrue(gameObject.GetAttr("kbd", "Down")) {
		gameObject.Position[1] -= speed * gameObject.DeltaTime
	}
}
func (check *CollisionCheck) OnEvent(gameObject *goz.GameObject, event *goz.Event) {
	fmt.Println(event.Msg, "from", event.Sender)
}

func main() {
	window := goz.OpenWindow(1024, 576, "Collisions")

	scene := goz.NewScene("scene001")

	// really big texture, will be scaled in the gameObject
	scene.NewTextureFromFilename("gozmo", "assets/gozmo.png")

	gozmo := scene.NewGameObject("Gozmo")
	gozmo.AddComponent("check", &CollisionCheck{})
	gozmo.AddComponent("kbd", goz.NewKeyboard())
	gozmo.AddComponent("renderer", goz.NewRenderer(nil))
	gozmo.SetAttr("renderer", "texture", "gozmo")
	gozmo.SetAttr("", "scaleX", 0.25)
	gozmo.SetAttr("", "scaleY", 0.25)

	gozmo.AddComponent("box", goz.NewBoxRenderer(10, 10))
	gozmo.SetAttr("box", "red", 0)
	gozmo.SetAttr("box", "green", 1)
	gozmo.SetAttr("box", "blue", 0)
	gozmo.SetAttr("box", "alpha", 0.2)

	gozmo.AddComponent("hit", goz.NewHitBox(0, 0, 10, 10))

	obstacle := scene.NewGameObject("Block Red")
	obstacle.SetAttr("", "positionX", -3)
	obstacle.SetAttr("", "positionY", 5)
	obstacle.AddComponent("box", goz.NewBoxRenderer(3, 3))
	obstacle.SetAttr("box", "red", 1)
	obstacle.SetAttr("box", "alpha", 1)
	obstacle.AddComponent("hit", goz.NewHitBoxWithEvent(0, 0, 3, 3, "red hit"))

	obstacle2 := scene.NewGameObject("Block Green")
	obstacle2.SetAttr("", "positionX", 3)
	obstacle2.SetAttr("", "positionY", -5)
	obstacle2.AddComponent("box", goz.NewBoxRenderer(1.5, 1.5))
	obstacle2.SetAttr("box", "green", 1)
	obstacle2.SetAttr("box", "alpha", 1)
	obstacle2.AddComponent("hit", goz.NewHitBoxWithEvent(0, 0, 1.5, 1.5, "green hit"))

	obstacle3 := scene.NewGameObject("Block Blue")
	obstacle3.SetAttr("", "positionX", -5)
	obstacle3.SetAttr("", "positionY", -5)
	obstacle3.AddComponent("box", goz.NewBoxRenderer(4, 4))
	obstacle3.SetAttr("box", "blue", 1)
	obstacle3.SetAttr("box", "alpha", 1)
	obstacle3.AddComponent("hit", goz.NewHitBoxWithEvent(0, 0, 4, 4, "blue hit"))

	window.SetScene(scene)
	window.Run()
}
