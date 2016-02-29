// +build android

package gozmo

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"golang.org/x/mobile/gl"
	"image"
)

var glctx gl.Context

func GLInit(width int32, height int32) {

	version := glctx.GetString(gl.VERSION)
	fmt.Println("OpenGL version", version)

	glctx.Enable(gl.BLEND)
	glctx.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	glctx.Viewport(0, 0, int(width), int(height))

	glctx.ClearColor(0, 0, 0, 1)
}

func GLClear() {
	glctx.Clear(gl.COLOR_BUFFER_BIT)
}

func GLTexture(rgba *image.RGBA) uint32 {
	texture := glctx.CreateTexture()
	glctx.ActiveTexture(gl.TEXTURE0)
	glctx.BindTexture(gl.TEXTURE_2D, texture)
	glctx.TexImage2D(gl.TEXTURE_2D, 0,
		rgba.Rect.Size().X,
		rgba.Rect.Size().Y,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		rgba.Pix)

	glctx.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	glctx.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	glctx.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	glctx.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	//gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	//gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)

	return texture.Value
}

func GLNewBuffer() uint32 {
	bid := glctx.CreateBuffer()
	glctx.BindBuffer(gl.ARRAY_BUFFER, bid)
	fmt.Println(bid)
	return bid.Value
}

// OpenGL ES has no VAO :(
func GLNewArray() uint32 {
	return 0
}

func GLBufferData(location uint32, bid uint32, data []float32) {
	glctx.BindBuffer(gl.ARRAY_BUFFER, gl.Buffer{Value: bid})
	glctx.BufferData(gl.ARRAY_BUFFER, data, gl.STATIC_DRAW)
	glctx.EnableVertexAttribArray(location)
	glctx.VertexAttribPointer(location, 2, gl.FLOAT, false, 0, gl.PtrOffset(0))
}

var boundsUniform int32 = -1
var orthoUniform int32 = -1
var uvDeltaUniform int32 = -1
var addColorUniform int32 = -1
var mulColorUniform int32 = -1

func GLDraw(renderer *Renderer, shader uint32, width float32, height float32, uvx, uvy, uvw, uvh float32, ortho mgl32.Mat4) {
	mesh := renderer.mesh
	texture := renderer.texture
	gl.UseProgram(shader)
	gl.Uniform2f(boundsUniform, width, height)
	gl.Uniform4f(uvDeltaUniform, uvx, uvy, uvw, uvh)
	addColor := renderer.addColor
	gl.Uniform4f(addColorUniform, addColor[0], addColor[1], addColor[2], addColor[3])
	mulColor := renderer.mulColor
	gl.Uniform4f(mulColorUniform, mulColor[0], mulColor[1], mulColor[2], mulColor[3])
	gl.UniformMatrix4fv(orthoUniform, 1, false, &ortho[0])
	gl.BindVertexArray(mesh.abid)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture.tid)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(mesh.vertices)/2))
}

var vertexShader = `
#version 330 core

layout(location = 0) in vec2 vertex;
layout(location = 1) in vec2 uv;

uniform vec4 uvdelta;
uniform vec2 bounds;
uniform mat4 ortho;

out vec2 uvout;

void main() {
    gl_Position = ortho * vec4(vertex.xy * bounds.xy, 0.0, 1.0);

    vec2 uv2 = uv;

    if (uv2.x == 0) {
        uv2.x = uvdelta.x;
    }
    else {
        uv2.x = uvdelta.x + uvdelta.z;
    }

    if (uv2.y == 0) {
        uv2.y = uvdelta.y;
    }
    else {
        uv2.y = uvdelta.y + uvdelta.w;
    }

    uvout = uv2;
}` + "\x00"

var fragmentShader = `
#version 330 core

uniform sampler2D tex;

uniform vec4 addColor;
uniform vec4 mulColor;

in vec2 uvout;
out vec4 color;

void main() {
    color = texture(tex, uvout) * mulColor + addColor;
}` + "\x00"

func GLShader() uint32 {
	vertexShaderId := gl.CreateShader(gl.VERTEX_SHADER)
	vscstr, free := gl.Strs(vertexShader)
	gl.ShaderSource(vertexShaderId, 1, vscstr, nil)
	free()
	gl.CompileShader(vertexShaderId)

	fragmentShaderId := gl.CreateShader(gl.FRAGMENT_SHADER)
	fscstr, free := gl.Strs(fragmentShader)
	gl.ShaderSource(fragmentShaderId, 1, fscstr, nil)
	free()
	gl.CompileShader(fragmentShaderId)

	programId := gl.CreateProgram()
	gl.AttachShader(programId, vertexShaderId)
	gl.AttachShader(programId, fragmentShaderId)
	gl.LinkProgram(programId)

	gl.DetachShader(programId, vertexShaderId)
	gl.DetachShader(programId, fragmentShaderId)

	gl.DeleteShader(vertexShaderId)
	gl.DeleteShader(fragmentShaderId)

	gl.UseProgram(programId)

	boundsUniform = gl.GetUniformLocation(programId, gl.Str("bounds\x00"))
	orthoUniform = gl.GetUniformLocation(programId, gl.Str("ortho\x00"))
	uvDeltaUniform = gl.GetUniformLocation(programId, gl.Str("uvdelta\x00"))
	addColorUniform = gl.GetUniformLocation(programId, gl.Str("addColor\x00"))
	mulColorUniform = gl.GetUniformLocation(programId, gl.Str("mulColor\x00"))

	texUniform := gl.GetUniformLocation(programId, gl.Str("tex\x00"))
	gl.Uniform1i(texUniform, 0)

	return programId
}
