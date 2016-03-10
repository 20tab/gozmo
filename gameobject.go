package gozmo

import (
	"fmt"
	"math"
	"sort"

	"github.com/go-gl/mathgl/mgl32"
)

// A GameObject stores all graphic parameters.
type GameObject struct {
	// TODO: is it a good idea to allow the developer to change the game object
	// name and mess with internal data?
	Name    string
	enabled bool
	order   int
	index   int

	Scene *Scene

	Position  mgl32.Vec2
	Rotation  float32
	Scale     mgl32.Vec2
	Pivot     mgl32.Vec2
	DeltaTime float32

	components     map[string]Component
	componentsKeys []string

	customAttrs map[string]interface{}

	events []*Event
}

func (scene *Scene) NewGameObject(name string) *GameObject {
	gameObject := GameObject{Name: name, order: 0}
	gameObject.components = make(map[string]Component)
	gameObject.customAttrs = make(map[string]interface{})
	// A game object always starts as enabled.
	gameObject.enabled = true
	gameObject.Scene = scene
	gameObject.Scale = mgl32.Vec2{1, 1}
	scene.gameObjects[name] = &gameObject
	// A -1 index means a not yet mapped gameObject.
	gameObject.index = -1
	gameObject.SetOrder(0)
	return &gameObject
}

func (scene *Scene) FindGameObject(name string) *GameObject {
	gameObject, ok := scene.gameObjects[name]
	if !ok {
		return nil
	}
	return gameObject
}

func (gameObject *GameObject) AddComponent(name string, component Component) Component {
	gameObject.components[name] = component
	gameObject.componentsKeys = append(gameObject.componentsKeys, name)
	component.Start(gameObject)
	return component
}

func (gameObject *GameObject) AddComponentByName(name string, componentName string, args []interface{}) Component {
	component := Engine.registeredComponents[componentName].Init(args)
	return gameObject.AddComponent(name, component)
}

func (gameObject *GameObject) SetOrder(order int) {
	scene := gameObject.Scene
	_, ok := scene.orderedGameObjects[order]
	if !ok {
		// Create a new slice of gameObjects.
		scene.orderedGameObjects[order] = make([]*GameObject, 0)
		// Cdd the new key (a number).
		scene.orderedKeys = append(scene.orderedKeys, order)
		// Sort them.
		sort.Ints(scene.orderedKeys)
	}

	if gameObject.index > -1 {
		scene.orderedGameObjects[gameObject.order] = append(scene.orderedGameObjects[gameObject.order][:gameObject.index], scene.orderedGameObjects[gameObject.order][gameObject.index+1:]...)
		// NOTE: it would be cool to remove an unused order layer, but it would
		// make things quite complex. Just remove the gameObject from the list.
	}
	// Set the index (for future removal).
	gameObject.index = len(scene.orderedGameObjects[order])
	// Set the new order.
	gameObject.order = order
	// Append it.
	scene.orderedGameObjects[order] = append(scene.orderedGameObjects[order], gameObject)
}

func (gameObject *GameObject) SetScale(x, y float32) {
	gameObject.Scale = mgl32.Vec2{x, y}
}

func (gameObject *GameObject) SetEuler(deg float32) {
	gameObject.Rotation = deg * math.Pi / 180
}

func (gameObject *GameObject) SetRotation(rad float32) {
	gameObject.Rotation = rad
}

func (gameObject *GameObject) SetPositionV(vec2 Vector2) {
	gameObject.Position = mgl32.Vec2(vec2)
}

func (gameObject *GameObject) AddPositionV(vec2 Vector2) {
	gameObject.Position = gameObject.Position.Add(mgl32.Vec2(vec2))
}

func (gameObject *GameObject) SetPosition(x float32, y float32) {
	gameObject.Position = mgl32.Vec2{x, y}
}

func (gameObject *GameObject) AddPosition(x float32, y float32) {
	gameObject.Position = gameObject.Position.Add(mgl32.Vec2{x, y})
}

func (gameObject *GameObject) SetEnabled(flag bool) {
	gameObject.enabled = flag
}

func (gameObject *GameObject) GetComponent(name string) interface{} {
	return gameObject.components[name]
}

func (gameObject *GameObject) GetComponentByType(name string) interface{} {
	for _, component := range gameObject.components {
		componentType, ok := component.(ComponentType)
		if ok {
			if componentType.GetType() == name {
				return component
			}
		}
	}
	return nil
}

// support both 32 and 64bit values
func (gameObject *GameObject) setAttr(attr string, value interface{}) error {
	switch attr {
	case "enabled":
		flag, _ := CastBool(value)
		gameObject.SetEnabled(flag)
	case "positionX":
		gameObject.Position[0], _ = CastFloat32(value)
	case "positionY":
		gameObject.Position[1], _ = CastFloat32(value)
	case "positionAddX":
		x, _ := CastFloat32(value)
		gameObject.Position[0] += x
	case "positionAddY":
		y, _ := CastFloat32(value)
		gameObject.Position[1] += y
	case "scaleX":
		gameObject.Scale[0], _ = CastFloat32(value)
	case "scaleY":
		gameObject.Scale[1], _ = CastFloat32(value)
	case "euler":
		r, _ := CastFloat32(value)
		gameObject.SetEuler(r)
	case "order":
		o, _ := CastInt(value)
		gameObject.SetOrder(o)
	// TODO: implement SetName() to maintain mappings and check for duplicates.
	case "name":
		name, ok := value.(string)
		if ok {
			gameObject.Name = name
			return nil
		}
		return fmt.Errorf("%v attribute of %T expects a string", attr, gameObject)
	}
	return nil
}

func (gameObject *GameObject) getAttr(attr string) (interface{}, error) {
	switch attr {
	case "enabled":
		return gameObject.enabled, nil
	case "positionX":
		return gameObject.Position[0], nil
	case "positionY":
		return gameObject.Position[1], nil
	case "scaleX":
		return gameObject.Scale[0], nil
	case "scaleY":
		return gameObject.Scale[1], nil
	case "euler":
		return gameObject.Rotation * 180 / math.Pi, nil
	case "deltaTime":
		return gameObject.DeltaTime, nil
	case "order":
		return gameObject.order, nil
	case "name":
		return gameObject.Name, nil
	}

	return nil, fmt.Errorf("attribute %v not found in %T", attr, gameObject)
}

func (gameObject *GameObject) SetAttr(componentName string, attr string, value interface{}) error {
	// Is it a base component?
	if componentName == "" {
		return gameObject.setAttr(attr, value)
	}

	// A custom attribute?
	if componentName == "{}" {
		gameObject.customAttrs[attr] = value
		return nil
	}

	component, ok := gameObject.components[componentName].(ComponentAttr)
	if !ok {
		return fmt.Errorf("component %v not found", componentName)
	}
	return component.SetAttr(attr, value)
}

func (gameObject *GameObject) GetAttr(componentName string, attr string) (interface{}, error) {
	// Is it a base component?
	if componentName == "" {
		return gameObject.getAttr(attr)
	}

	// A custom attribute?
	if componentName == "{}" {
		v, ok := gameObject.customAttrs[attr]
		if !ok {
			return nil, fmt.Errorf("unknown attr")
		}
		return v, nil
	}

	component, ok := gameObject.components[componentName].(ComponentAttr)
	if !ok {
		return nil, fmt.Errorf("component %v not found", componentName)
	}
	return component.GetAttr(attr)
}

func (gameObject *GameObject) Update() {
	for _, key := range gameObject.componentsKeys {
		gameObject.components[key].Update(gameObject)
	}
}

func (gameObject *GameObject) Destroy() {
	// Call Destroy() on all associated components.
}
