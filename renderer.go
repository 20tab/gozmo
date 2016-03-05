package gozmo

/*

this is an accelerated sprite drawer component

It supports color addition and multiplication

*/

import (
	"fmt"
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

type Renderer struct {
	mesh          *Mesh
	texture       *Texture
	textureName   string
	pixelsPerUnit uint32
	index         uint32
	forceHeight   float32
}

var shader int32 = -1

// the mesh is created and uploaded into the GPU only when needed
func (renderer *Renderer) createMesh() {
	if shader == -1 {
		shader = int32(GLShader())
	}
	mesh := Mesh{}

	mesh.abid = GLNewArray()
	mesh.vbid = GLNewBuffer()
	mesh.uvbid = GLNewBuffer()

	mesh.vertices = []float32{-1, -1,
		-1, 1,
		1, -1,
		1, -1,
		1, 1,
		-1, 1}

	mesh.uvs = []float32{0, 1,
		0, 0,
		1, 1,
		1, 1,
		1, 0,
		0, 0}

	mesh.mulColor = mgl32.Vec4{1, 1, 1, 1}

	GLBufferData(0, mesh.vbid, mesh.vertices)

	GLBufferData(1, mesh.uvbid, mesh.uvs)

	renderer.mesh = &mesh
}

func NewRenderer(texture *Texture) *Renderer {
	// default 100 pixels per unit (like in Unity3D)
	renderer := Renderer{texture: texture, pixelsPerUnit: 100}

	if texture != nil {
		renderer.textureName = texture.Name
		renderer.createMesh()
	}

	return &renderer
}

func (renderer *Renderer) Start(gameObject *GameObject) {
}

func (renderer *Renderer) Update(gameObject *GameObject) {
	if renderer.textureName == "" {
		return
	}

	renderer.texture, _ = gameObject.Scene.textures[renderer.textureName]

	if renderer.texture == nil {
		return
	}

	if renderer.mesh == nil {
		renderer.createMesh()
	}

	texture := renderer.texture

	// recompute mesh size based on the texture
	var width float32
	var height float32
	if renderer.forceHeight > 0 {
		height = renderer.forceHeight / 2
		width = renderer.forceHeight * ((float32(texture.Width) / float32(texture.Cols)) / (float32(texture.Height) / float32(texture.Rows))) / 2
	} else {
		width = float32(texture.Width) / float32(texture.Cols) / float32(renderer.pixelsPerUnit) / 2
		height = float32(texture.Height) / float32(texture.Rows) / float32(renderer.pixelsPerUnit) / 2
	}

	// recompute uvs based on index
	idxX := renderer.index % texture.Cols
	idxY := renderer.index / texture.Cols

	uvw := (1.0 / float32(texture.Cols))
	uvh := (1.0 / float32(texture.Rows))

	uvx := uvw * float32(idxX)
	uvy := uvh * float32(idxY)

	model := mgl32.Translate3D(gameObject.Position[0], gameObject.Position[1], 0)

	model = model.Mul4(mgl32.Scale3D(gameObject.Scale[0], gameObject.Scale[1], 1))

	model = model.Mul4(mgl32.HomogRotate3DZ(gameObject.Rotation))

	view := Engine.Window.View.Mul4(model)

	ortho := Engine.Window.Projection.Mul4(view)

	GLDraw(renderer.mesh, uint32(shader), width, height, int32(renderer.texture.tid), uvx, uvy, uvw, uvh, ortho)
}

func (renderer *Renderer) SetPixelsPerUnit(pixels uint32) {
	renderer.pixelsPerUnit = pixels
}

func (renderer *Renderer) SetAttr(attr string, value interface{}) error {
	switch attr {
	case "index":
		index, err := CastUInt32(value)
		if err != nil {
			return fmt.Errorf("%v attribute of %T", attr, renderer, err)
		}
		renderer.index = index
		return nil
	case "texture":
		textureName, ok := value.(string)
		if ok {
			renderer.textureName = textureName
			return nil
		}
		return fmt.Errorf("%v attribute of %T expects a string", attr, renderer)
	case "addR":
		color, ok := value.(float32)
		if ok {
			renderer.mesh.addColor[0] = color
			return nil
		}
		return fmt.Errorf("%v attribute of %T expects a float32", attr, renderer)
	case "addG":
		color, ok := value.(float32)
		if ok {
			renderer.mesh.addColor[1] = color
			return nil
		}
		return fmt.Errorf("%v attribute of %T expects a float32", attr, renderer)
	case "addB":
		color, ok := value.(float32)
		if ok {
			renderer.mesh.addColor[2] = color
			return nil
		}
		return fmt.Errorf("%v attribute of %T expects a float32", attr, renderer)
	case "addA":
		color, ok := value.(float32)
		if ok {
			renderer.mesh.addColor[3] = color
			return nil
		}
		return fmt.Errorf("%v attribute of %T expects a float32", attr, renderer)
	case "mulR":
		color, ok := value.(float32)
		if ok {
			renderer.mesh.mulColor[0] = color
			return nil
		}
		return fmt.Errorf("%v attribute of %T expects a float32", attr, renderer)
	case "mulG":
		color, ok := value.(float32)
		if ok {
			renderer.mesh.mulColor[1] = color
			return nil
		}
		return fmt.Errorf("%v attribute of %T expects a float32", attr, renderer)
	case "mulB":
		color, ok := value.(float32)
		if ok {
			renderer.mesh.mulColor[2] = color
			return nil
		}
		return fmt.Errorf("%v attribute of %T expects a float32", attr, renderer)
	case "mulA":
		color, ok := value.(float32)
		if ok {
			renderer.mesh.mulColor[3] = color
			return nil
		}
		return fmt.Errorf("%v attribute of %T expects a float32", attr, renderer)
	case "forceHeight":
		height, err := CastFloat32(value)
		if err == nil {
			renderer.forceHeight = height
			return nil
		}
		return fmt.Errorf("%v attribute of %T expects a float32", attr, renderer)
	}
	return nil
}

func (renderer *Renderer) GetAttr(attr string) (interface{}, error) {
	switch attr {
	case "index":
		return renderer.index, nil
	case "texture":
		return renderer.textureName, nil
	case "addR":
		return renderer.mesh.addColor[0], nil
	case "addG":
		return renderer.mesh.addColor[1], nil
	case "addB":
		return renderer.mesh.addColor[2], nil
	case "addA":
		return renderer.mesh.addColor[3], nil
	case "mulR":
		return renderer.mesh.mulColor[0], nil
	case "mulG":
		return renderer.mesh.mulColor[1], nil
	case "mulB":
		return renderer.mesh.mulColor[2], nil
	case "mulA":
		return renderer.mesh.mulColor[3], nil
	}
	return nil, fmt.Errorf("%v attribute of %T not found", attr, renderer)
}

func (renderer *Renderer) GetType() string {
	return "Renderer"
}

func initRenderer(args []interface{}) Component {
	return NewRenderer(nil)
}

func init() {
	RegisterComponent("Renderer", initRenderer)
}
