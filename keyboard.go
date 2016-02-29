// +build !android

package gozmo

/*

Keyboard mappins, by default we will use the glfw names
but we need to duplicate them to support more platforms in the future

*/

import (
    "github.com/go-gl/glfw/v3.1/glfw"
)

type Key glfw.Key

const (
    KeyRight Key = Key(glfw.KeyRight)
    KeyLeft Key = Key(glfw.KeyLeft)
    KeyUp Key = Key(glfw.KeyUp)
    KeyDown Key = Key(glfw.KeyDown)

    KeyA Key = Key(glfw.KeyA)
    KeyB Key = Key(glfw.KeyB)
    KeyC Key = Key(glfw.KeyC)
    KeyP Key = Key(glfw.KeyP)
    KeyS Key = Key(glfw.KeyS)
    KeyT Key = Key(glfw.KeyT)

    KeyEsc Key = Key(glfw.KeyEscape)
)

/*

TODO decide if the input system can be managed
as attributes or we need a specific layer

*/
var KeyboardAttr map[string]Key = map[string]Key {
                   "Right": KeyRight,
               }
