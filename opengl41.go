// +build !android

package gozmo

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"image"
)

func GLInit(width int32, height int32) {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	gl.Viewport(0, 0, width, height)

	gl.ClearColor(0, 0, 0, 1)
}

func GLClear() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func GLTexture(rgba *image.RGBA) uint32 {
	var texture uint32
	gl.GenTextures(1, &texture)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	//gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	//gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)

	return texture
}

func GLNewBuffer() uint32 {
	var bid uint32
	gl.GenBuffers(1, &bid)
	gl.BindBuffer(gl.ARRAY_BUFFER, bid)
	return bid
}

func GLNewArray() uint32 {
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	return vao
}

func GLBufferData(location uint32, bid uint32, data []float32) {
	gl.BindBuffer(gl.ARRAY_BUFFER, bid)
	gl.BufferData(gl.ARRAY_BUFFER, len(data)*4, gl.Ptr(data), gl.STATIC_DRAW)
	gl.EnableVertexAttribArray(location)
	gl.VertexAttribPointer(location, 2, gl.FLOAT, false, 0, gl.PtrOffset(0))
}

var boundsUniform int32 = -1
var orthoUniform int32 = -1
var uvDeltaUniform int32 = -1
var addColorUniform int32 = -1
var mulColorUniform int32 = -1

func GLDraw(mesh *Mesh, shader uint32, width float32, height float32, textureId int32, uvx, uvy, uvw, uvh float32, ortho mgl32.Mat4) {
	gl.UseProgram(shader)
	gl.Uniform2f(boundsUniform, width, height)
	gl.Uniform4f(uvDeltaUniform, uvx, uvy, uvw, uvh)
	addColor := mesh.addColor
	gl.Uniform4f(addColorUniform, addColor[0], addColor[1], addColor[2], addColor[3])
	mulColor := mesh.mulColor
	gl.Uniform4f(mulColorUniform, mulColor[0], mulColor[1], mulColor[2], mulColor[3])
	gl.UniformMatrix4fv(orthoUniform, 1, false, &ortho[0])
	gl.BindVertexArray(mesh.abid)
	if textureId > -1 {
		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, uint32(textureId))
	} else {
		// disable texture
		gl.BindTexture(gl.TEXTURE_2D, 0)
	}
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
