// +build !android

package gozmo

import (
	"github.com/go-gl/glfw/v3.1/glfw"
	"log"
	"runtime"
)

type Window struct {
	width        int32
	height       int32
	title        string
	glfwWindow   *glfw.Window
	scenes       []*Scene
	currentScene *Scene
}

func OpenWindow(width int32, height int32, title string) *Window {
	runtime.LockOSThread()
	window := Window{width: width, height: height, title: title}

	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.DoubleBuffer, 1)

	//glfw.WindowHint(glfw.Samples, 2)

	// TODO support full screen
	//glfwin, err := glfw.CreateWindow(window.width, window.height, window.title, glfw.GetPrimaryMonitor(), nil)
	glfwin, err := glfw.CreateWindow(int(window.width), int(window.height), window.title, nil, nil)
	if err != nil {
		panic(err)
	}

	glfwin.MakeContextCurrent()

	window.glfwWindow = glfwin

	glfw.SwapInterval(1)
	//glfw.SwapInterval(0)

	fbWidth, fbHeight := glfwin.GetFramebufferSize()
	GLInit(int32(fbWidth), int32(fbHeight))

	return &window
}

func (window *Window) Run() {
	win := window.glfwWindow

	glfw.SetTime(0.0)

	for !win.ShouldClose() {

		GLClear()

		scene := window.currentScene
		if scene == nil {
			if len(window.scenes) > 0 {
				scene = window.scenes[0]
			}
		}

		if scene != nil {
			scene.Update(glfw.GetTime())
		}

		win.SwapBuffers()
		glfw.PollEvents()
	}
	glfw.Terminate()
}

func (window *Window) getKey(kc Key) bool {
	return window.glfwWindow.GetKey(glfw.Key(kc)) == glfw.Press
}
