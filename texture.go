package gozmo

import (
	"image"
	"image/draw"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

type Texture struct {
	tid    uint32
	Name   string
	Width  uint32
	Height uint32
	Rows   uint32
	Cols   uint32
}

func (scene *Scene) NewTextureFilename(name string, fileName string) (*Texture, error) {
	imgFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	return scene.NewTextureFile(name, imgFile)
}

func (scene *Scene) NewTextureFile(name string, file *os.File) (*Texture, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	tid := GLTexture(rgba)

	tex := Texture{tid: tid, Name: name, Width: uint32(rgba.Bounds().Size().X), Height: uint32(rgba.Bounds().Size().Y)}

	tex.Rows = 1
	tex.Cols = 1

	scene.textures[name] = &tex

	return &tex, nil
}

func (scene *Scene) NewTexture(name string, width uint32, height uint32) {
}

func (texture *Texture) SetRows(rows uint32) {
	texture.Rows = rows
}

func (texture *Texture) SetCols(cols uint32) {
	texture.Cols = cols
}

func (texture *Texture) SetRowsCols(rows, cols uint32) {
	texture.Rows = rows
	texture.Cols = cols
}

func (texture *Texture) Destroy() {
	// TODO delete texture from the GPU
}
