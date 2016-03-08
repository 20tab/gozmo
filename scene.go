package gozmo

/*

A Scene is a group of resources (textures, animations, sounds)
and instantiated GameObjects

When a scene is destroyed, all of the allocated resources and GameObjects are destroyed.

*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Scene struct {
	// TODO is it a good idea to expose Name ?
	Name        string
	gameObjects map[string]*GameObject
	textures    map[string]*Texture
	animations  map[string]*Animation
	// this contains the last timestamp of the engine
	lastTime           float64
	orderedGameObjects map[int][]*GameObject
	orderedKeys        []int
}

func (scene *Scene) Update(now float64) {
	deltaTime := float32(now - scene.lastTime)
	scene.lastTime = now

	for _, order := range scene.orderedKeys {
		for _, gameObject := range scene.orderedGameObjects[order] {
			if !gameObject.enabled {
				continue
			}
			gameObject.DeltaTime = deltaTime
			for _, key := range gameObject.componentsKeys {
				gameObject.components[key].Update(gameObject)
			}
		}
	}

	for _, updater := range Engine.registeredUpdaters {
		updater(scene, deltaTime)
	}

	UpdatePerFrameStats()
}

func NewScene(name string) *Scene {
	scene := Scene{Name: name}
	scene.gameObjects = make(map[string]*GameObject)
	scene.textures = make(map[string]*Texture)
	scene.animations = make(map[string]*Animation)

	scene.orderedGameObjects = make(map[int][]*GameObject)

	if Engine.scenes == nil {
		Engine.scenes = make(map[string]*Scene)
	}

	Engine.scenes[name] = &scene
	return &scene
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
			tex, err = scene.NewTextureFromFilename(name.(string), filename.(string))
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

		var args []interface{} = nil
		argsItem, ok := componentMap["args"]
		if ok {
			argsList, ok := argsItem.([]interface{})
			if ok {
				args = argsList
			}
		}

		gameObject.AddComponentName(componentName.(string), componentType.(string), args)
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

		err := gameObject.SetAttr(component.(string), key.(string), value)
		if err != nil {
			fmt.Println(err)
		}
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

func loadAnimations(scene *Scene, animations []interface{}) {
	for _, anim := range animations {
		animMap := anim.(map[string]interface{})

		name, ok := animMap["name"]
		if !ok {
			panic("animation requires a name")
		}

		fps, hasFps := animMap["fps"]
		loop, hasLoop := animMap["loop"]

		frames, hasFrames := animMap["frames"]

		animation := scene.AddAnimation(name.(string), 0, false)

		if hasFps {
			animation.Fps = int(fps.(float64))
		}

		if hasLoop {
			animation.Loop = loop.(bool)
		}

		if hasFrames {
			var actionsList []*AnimationAction
			for _, frame := range frames.([]interface{}) {
				actions := frame.([]interface{})
				for _, action := range actions {
					actionMap := action.(map[string]interface{})
					component, ok := actionMap["component"]
					if !ok {
						panic("animation action requires a component")
					}
					key, ok := actionMap["key"]
					if !ok {
						panic("animation action requires a key")
					}
					value, ok := actionMap["value"]
					if !ok {
						panic("animation action requires a value")
					}
					actionsList = append(actionsList, &AnimationAction{ComponentName: component.(string), Attr: key.(string), Value: value})
				}
				animation.AddFrame(actionsList)
			}
		}

	}
}

func NewSceneFromFilename(fileName string) *Scene {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	var parsed map[string]interface{}

	err = json.Unmarshal(data, &parsed)
	if err != nil {
		panic(err)
	}

	name, ok := parsed["name"]
	if !ok {
		panic("a scene requires a name")
	}

	scene := NewScene(name.(string))

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
		case "animations":
			animations := value.([]interface{})
			loadAnimations(scene, animations)
		}
	}

	return scene
}

func (scene *Scene) Destroy() {
	// first of all destroy objects
	for _, gameObject := range scene.gameObjects {
		gameObject.Destroy()
	}

	// the destroy textures
	for _, texture := range scene.textures {
		texture.Destroy()
	}

	// remove the scene from the map, so the GC can make its job
	delete(Engine.scenes, scene.Name)
}
