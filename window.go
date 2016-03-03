// +build !android

package gozmo

import (
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"log"
	"runtime"
)

type Window struct {
	width        int32
	height       int32
	title        string
	glfwWindow   *glfw.Window
	currentScene *Scene
	Projection   mgl32.Mat4
}

func OpenWindowVersion(width int32, height int32, title string, major int, minor int) *Window {
	if Engine.Window != nil {
		panic("a window is already active")
	}
	runtime.LockOSThread()
	window := Window{width: width, height: height, title: title}

	if Engine.scenes == nil {
		Engine.scenes = make(map[string]*Scene)
	}

	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, major)
	glfw.WindowHint(glfw.ContextVersionMinor, minor)
	if major >= 4 || (major >= 3 && minor >= 2) {
		glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	}
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

	ratio := float32(width) / float32(height)

	window.Projection = mgl32.Ortho2D(-10*ratio, 10*ratio, -10.0, 10.0)
	window.glfwWindow = glfwin

	glfw.SwapInterval(1)
	//glfw.SwapInterval(0)

	fbWidth, fbHeight := glfwin.GetFramebufferSize()
	GLInit(int32(fbWidth), int32(fbHeight))

	Engine.Window = &window

	return &window
}

func OpenWindow(width int32, height int32, title string) *Window {
	return OpenWindowVersion(width, height, title, 3, 3)
}

func (window *Window) Run() {
	win := window.glfwWindow

	glfw.SetTime(0.0)

	for !win.ShouldClose() {

		GLClear()

		scene := window.currentScene
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

func (window *Window) SetScene(scene *Scene) {
        window.currentScene = scene
}

func (window *Window) SetSceneByName(sceneName string) {
        window.currentScene = Engine.scenes[sceneName]
}
