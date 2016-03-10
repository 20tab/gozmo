// +build !android

package gozmo

import (
	_ "github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

// The Mouse component converts screen coordinates into game coordinates. It
// only handles mouse events, not touch ones, and is not bound on Android.
type Mouse struct{}

func (mouse *Mouse) Start(gameObject *GameObject)  {}
func (mouse *Mouse) Update(gameObject *GameObject) {}

func (mouse *Mouse) SetAttr(attr string, value interface{}) error {
	return nil
}

func (mouse *Mouse) GetName() string {
	return "Mouse"
}

func (mouse *Mouse) X() float32 {
	x, y := Engine.Window.glfwWindow.GetCursorPos()
	if x < 0 {
		x = 0
	}
	if x > 1024-1 {
		x = 1024 - 1
	}
	vecScreen := mgl32.Vec4{float32(2*x/1024) - 1, float32(2*y/576) - 1, 0, 1}
	vecWorld := Engine.Window.Projection.Inv().Mul4x1(vecScreen)
	return vecWorld[0]
}

func (mouse *Mouse) Y() float32 {
	return 0
}

// TODO: what if the user specifies an unknown key?
func (mouse *Mouse) GetAttr(attr string) (interface{}, error) {
	return 0, nil
}

func NewMouse() *Mouse {
	mouse := Mouse{}
	return &mouse
}

func initMouse(args []interface{}) Component {
	return NewMouse()
}

func init() {
	RegisterComponent("Mouse", initMouse)
}
