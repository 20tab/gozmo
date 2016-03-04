package gozmo

/*

Camera component

ensure cameras are always managed befre the other items
by setting them a lower order

*/

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct{}

func (camera *Camera) Start(gameObject *GameObject) {}
func (camera *Camera) Update(gameObject *GameObject) {
	Engine.Window.View = mgl32.LookAt(gameObject.Position[0], gameObject.Position[1], 1, gameObject.Position[0], gameObject.Position[1], 0, 0, 1, 0)
}

func (camera *Camera) SetAttr(attr string, value interface{}) error {
	return nil
}

func (camera *Camera) GetName() string {
	return "Camera"
}

func (camera *Camera) GetAttr(attr string) (interface{}, error) {
	return 0, nil
}

func NewCamera() *Camera {
	camera := Camera{}
	return &camera
}

func initCamera(args []interface{}) Component {
	return NewCamera()
}

func init() {
	RegisterComponent("Camera", initCamera)
}
