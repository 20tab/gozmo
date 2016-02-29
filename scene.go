package gozmo

type Scene struct {
    Name string
    Window *Window
    gameObjects []*GameObject
    textures map[string]*Texture
    animations map[string]*Animation
    lastTime float64
}

func (scene *Scene) Update(now float64) {
    deltaTime := float32(now - scene.lastTime)
    scene.lastTime = now

    for _, gameObject := range scene.gameObjects {
        gameObject.DeltaTime = deltaTime
        for _, key := range gameObject.componentsKeys {
            gameObject.components[key].Update(gameObject)
        }
    }
}

func (window *Window) NewScene() *Scene {
    scene := Scene{Window: window}
    scene.textures = make(map[string]*Texture)
    scene.animations = make(map[string]*Animation)
    window.scenes = append(window.scenes, &scene)
    return &scene
}

func (window *Window) SetScene(scene *Scene) {
    window.currentScene = scene
}

