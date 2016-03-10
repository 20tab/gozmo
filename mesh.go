package gozmo

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Mesh struct {
	vertices []float32
	uvs      []float32

	vbid  uint32
	uvbid uint32

	abid uint32

	addColor mgl32.Vec4
	mulColor mgl32.Vec4
}

// Points to the shader id.
var shader int32 = -1
