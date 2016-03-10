// +build !android

package gozmo

import (
	"github.com/go-gl/glfw/v3.1/glfw"
)

// A Key is a keyboard mapping. By default we use the glfw names, but override
// them to support more platforms in the future. Not bound on Android.
type Key glfw.Key

const (
	KeyRight Key = Key(glfw.KeyRight)
	KeyLeft  Key = Key(glfw.KeyLeft)
	KeyUp    Key = Key(glfw.KeyUp)
	KeyDown  Key = Key(glfw.KeyDown)

	KeyA Key = Key(glfw.KeyA)
	KeyB Key = Key(glfw.KeyB)
	KeyC Key = Key(glfw.KeyC)
	KeyD Key = Key(glfw.KeyD)
	KeyP Key = Key(glfw.KeyP)
	KeyS Key = Key(glfw.KeyS)
	KeyT Key = Key(glfw.KeyT)
	KeyW Key = Key(glfw.KeyW)

	KeyEsc Key = Key(glfw.KeyEscape)
)

// Rather ugly mapping, but it will simplify the "compositor's" life.
var KeyboardAttr map[string]Key = map[string]Key{
	"A":     KeyA,
	"D":     KeyD,
	"S":     KeyS,
	"W":     KeyW,
	"Right": KeyRight,
	"Left":  KeyLeft,
	"Up":    KeyUp,
	"Down":  KeyDown,
}

// A Keyboard works by polling at a given time.
type Keyboard struct{}

func (keyboard *Keyboard) Start(gameObject *GameObject)  {}
func (keyboard *Keyboard) Update(gameObject *GameObject) {}

func (keyboard *Keyboard) SetAttr(attr string, value interface{}) error {
	return nil
}

func (keyboard *Keyboard) GetName() string {
	return "Keyboard"
}

func (keyboard *Keyboard) GetKey(key Key) bool {
	return Engine.Window.getKey(key)
}

// TODO: what if the user specifies an unknown key?
func (keyboard *Keyboard) GetAttr(attr string) (interface{}, error) {
	key, ok := KeyboardAttr[attr]
	if !ok {
		return false, nil
	}
	return Engine.Window.getKey(key), nil
}

func NewKeyboard() *Keyboard {
	keyboard := Keyboard{}
	return &keyboard
}

func initKeyboard(args []interface{}) Component {
	return NewKeyboard()
}

func init() {
	RegisterComponent("Keyboard", initKeyboard)
}
