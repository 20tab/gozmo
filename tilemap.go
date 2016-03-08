package gozmo

/*

a simple tilemap tilemap using a single mesh for the whole level

*/

import (
	_ "fmt"
	"github.com/go-gl/mathgl/mgl32"
)

type TileMap struct {
	mesh    *Mesh
	texture *Texture

	pixelsPerUnit uint32
}

func NewTileMap(texture *Texture) *TileMap {
	// default 100 pixels per unit (like in Unity3D)
	tilemap := TileMap{texture: texture, pixelsPerUnit: 100}

	return &tilemap
}

func (tilemap *TileMap) Start(gameObject *GameObject) {
	if shader == -1 {
		shader = int32(GLShader())
	}
	mesh := Mesh{}

	mesh.abid = GLNewArray()
	mesh.vbid = GLNewBuffer()
	mesh.uvbid = GLNewBuffer()

	//for y := 0; y > -30; y-- {
	//for x := 0; x < 170; x++ {
	for y := float32(0); y > -5; y-- {
		for x := float32(0); x < 7; x++ {
			// x -1
			mesh.vertices = append(mesh.vertices, 2*x-1)
			// y -1
			mesh.vertices = append(mesh.vertices, 2*y-1)
			// x -1
			mesh.vertices = append(mesh.vertices, 2*x-1)
			// y 1
			mesh.vertices = append(mesh.vertices, 2*y+1)
			// x 1
			mesh.vertices = append(mesh.vertices, 2*x+1)
			// y -1
			mesh.vertices = append(mesh.vertices, 2*y-1)
			// x 1
			mesh.vertices = append(mesh.vertices, 2*x+1)
			// y -1
			mesh.vertices = append(mesh.vertices, 2*y-1)
			// x 1
			mesh.vertices = append(mesh.vertices, 2*x+1)
			// y 1
			mesh.vertices = append(mesh.vertices, 2*y+1)
			// x -1
			mesh.vertices = append(mesh.vertices, 2*x-1)
			// y 1
			mesh.vertices = append(mesh.vertices, 2*y+1)

			mesh.uvs = append(mesh.uvs, 0)
			mesh.uvs = append(mesh.uvs, 1.0/11)

			mesh.uvs = append(mesh.uvs, 0)
			mesh.uvs = append(mesh.uvs, 0)

			mesh.uvs = append(mesh.uvs, 1.0/12)
			mesh.uvs = append(mesh.uvs, 1.0/11)

			mesh.uvs = append(mesh.uvs, 1.0/12)
			mesh.uvs = append(mesh.uvs, 1.0/11)

			mesh.uvs = append(mesh.uvs, 1.0/12)
			mesh.uvs = append(mesh.uvs, 0)

			mesh.uvs = append(mesh.uvs, 0)
			mesh.uvs = append(mesh.uvs, 0)
		}
	}

	mesh.mulColor = mgl32.Vec4{1, 1, 1, 1}

	GLBufferData(0, mesh.vbid, mesh.vertices)

	GLBufferData(1, mesh.uvbid, mesh.uvs)

	tilemap.mesh = &mesh
}

func (tilemap *TileMap) Update(gameObject *GameObject) {

	texture := tilemap.texture

	// recompute mesh size based on the texture
	var width float32 = 1
	var height float32 = 1

	//width = float32(texture.Width) / float32(texture.Cols) / float32(tilemap.pixelsPerUnit) / 2
	//height = float32(texture.Height) / float32(texture.Rows) / float32(tilemap.pixelsPerUnit) / 2

	// recompute uvs based on index
	/*
		idxX := tilemap.index % texture.Cols
		idxY := tilemap.index / texture.Cols

		uvw := (1.0 / float32(texture.Cols))
		uvh := (1.0 / float32(texture.Rows))

		uvx := uvw * float32(idxX)
		uvy := uvh * float32(idxY)
	*/

	var uvx float32 = 0
	var uvy float32 = 0
	var uvw float32 = 0
	var uvh float32 = 0

	model := mgl32.Translate3D(gameObject.Position[0], gameObject.Position[1], 0)

	model = model.Mul4(mgl32.Scale3D(gameObject.Scale[0], gameObject.Scale[1], 1))

	model = model.Mul4(mgl32.HomogRotate3DZ(gameObject.Rotation))

	view := Engine.Window.View.Mul4(model)

	ortho := Engine.Window.Projection.Mul4(view)

	IncPerFrameStats("GL.DrawCalls", 1)

	GLDraw(tilemap.mesh, uint32(shader), width, height, int32(texture.tid), uvx, uvy, uvw, uvh, ortho)
}

func (tilemap *TileMap) SetPixelsPerUnit(pixels uint32) {
	tilemap.pixelsPerUnit = pixels
}

func (tilemap *TileMap) SetAttr(attr string, value interface{}) error {
	return nil
}

func (tilemap *TileMap) GetAttr(attr string) (interface{}, error) {
	return nil, nil
}

func (tilemap *TileMap) GetType() string {
	return "TileMap"
}

func initTileMap(args []interface{}) Component {
	return NewTileMap(nil)
}

func init() {
	RegisterComponent("TileMap", initTileMap)
}
