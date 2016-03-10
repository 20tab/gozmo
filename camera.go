package gozmo

import (
	"github.com/go-gl/mathgl/mgl32"
)

// The Camera component sets the current view matrix for rendering, allowing
// different views to coexist. It ensures that cameras are always managed
// before other items by setting them at at lower place.
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
