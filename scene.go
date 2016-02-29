package gozmo

/*

A Scene is a group of resources (textures, animations, sounds)
and instantiated GameObjects

When a scene is destroyed, all of the allocated resources and GameObjects are destroyed.

*/

import (
    "io/ioutil"
    "encoding/json"
)

type Scene struct {
    Name string
    Window *Window
    gameObjects []*GameObject
    textures map[string]*Texture
    animations map[string]*Animation
    // this contains the last timestamp of the engine
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

func loadTextures(scene *Scene, textures []interface{}) {
    for _, texture := range textures {
        texMap := texture.(map[string]interface{})

        name, ok := texMap["name"]
        if !ok {
            panic("texture requires a name")
        }

        filename, hasFilename := texMap["filename"]

        rows, hasRows := texMap["rows"]
        cols, hasCols := texMap["cols"]

        var tex *Texture
        var err error

        if hasFilename {
            tex, err = scene.NewTextureFilename(name.(string), filename.(string))
            if err != nil {
                panic(err)
            }
        }

        if tex == nil {
            continue
        }

        if hasRows {
            tex.SetRows(uint32(rows.(float64)))
        }

        if hasCols {
            tex.SetCols(uint32(cols.(float64)))
        }
    }
}

func addComponents(gameObject *GameObject, components []interface{}) {
    for _, component := range components {
        componentMap := component.(map[string]interface{})

        componentName, ok := componentMap["name"]
        if !ok {
            panic("component requires a name")
        }
        componentType, ok := componentMap["type"]
        if !ok {
            panic("component requires a type")
        }

        gameObject.AddComponentName(componentName.(string), componentType.(string), nil)
    }
}

func setAttrs(gameObject *GameObject, attrs []interface{}) {
    for _, attr := range attrs {

        attrMap := attr.(map[string]interface{})

        component, ok := attrMap["component"]
        if !ok {
            panic("attr requires a component")
        }

        key, ok := attrMap["key"]
        if !ok {
            panic("attr requires a key")
        }

        value, ok := attrMap["value"]
        if !ok {
            panic("attr requires a value")
        }

        gameObject.SetAttr(component.(string), key.(string), value)
    }
}

func loadObjects(scene *Scene, objects []interface{}) {
    for _, obj := range objects {
        objMap := obj.(map[string]interface{})

        name, ok := objMap["name"]
        if !ok {
            panic("object requires a name")
        }

        components, hasComponents := objMap["components"]
        attrs, hasAttrs := objMap["attrs"]

        gameObject := scene.NewGameObject(name.(string))
        if hasComponents {
            addComponents(gameObject, components.([]interface{}))
        }

        if hasAttrs {
            setAttrs(gameObject, attrs.([]interface{}))
        }
    }
}

func (window *Window) NewSceneFilename(fileName string) *Scene {
    scene := window.NewScene()
    data, err := ioutil.ReadFile(fileName)
    if err != nil {
        panic(err)
    }

    var parsed map[string]interface{}

    err = json.Unmarshal(data, &parsed)
    if err != nil {
        panic(err)
    }

    for key, value := range parsed {
        switch key {
            case "name":
                scene.Name = value.(string)
            case "textures":
                textures := value.([]interface{})
                loadTextures(scene, textures)
            case "objects":
                objects := value.([]interface{})
                loadObjects(scene, objects)
        }
    }

    return scene
}
