package gozmo

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"math"
	"strings"
)

type GameObject struct {
	Name    string
	enabled bool
	order   int32

	Scene *Scene

	Position  mgl32.Vec2
	Rotation  float32
	Scale     mgl32.Vec2
	Pivot     mgl32.Vec2
	DeltaTime float32

	components     map[string]Component
	componentsKeys []string

	customAttrs map[string]interface{}
}

func (scene *Scene) NewGameObject(name string) *GameObject {
	gameObject := GameObject{Name: name, order: 0}
	gameObject.components = make(map[string]Component)
	gameObject.customAttrs = make(map[string]interface{})
	gameObject.Scene = scene
	gameObject.Scale = mgl32.Vec2{1, 1}
	scene.gameObjects = append(scene.gameObjects, &gameObject)
	return &gameObject
}

func (gameObject *GameObject) AddComponent(name string, component Component) Component {
	gameObject.components[name] = component
	gameObject.componentsKeys = append(gameObject.componentsKeys, name)
	component.Start(gameObject)
	return component
}

func (gameObject *GameObject) AddComponentName(name string, componentName string, args []interface{}) Component {
	component := Engine.registeredComponents[componentName].Init(args)
	return gameObject.AddComponent(name, component)
}

func (gameObject *GameObject) SetOrder(order int32) {
}

func (gameObject *GameObject) GetKey(kc Key) bool {
	return gameObject.Scene.Window.getKey(kc)
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

func (gameObject *GameObject) GetComponent(name string) interface{} {
	return gameObject.components[name]
}

// support both 32 and 64bit values
func (gameObject *GameObject) setAttr(attr string, value interface{}) error {
	switch attr {
	case "positionX":
		x, ok := value.(float32)
		if ok {
			gameObject.Position[0] = x
			return nil
		}
		x64, ok := value.(float64)
		if ok {
			gameObject.Position[0] = float32(x64)
			return nil
		}
		return fmt.Errorf("%v attribute of %T expects a float32", attr, gameObject)
	case "positionY":
		y, ok := value.(float32)
		if ok {
			gameObject.Position[1] = y
			return nil
		}
		y64, ok := value.(float64)
		if ok {
			gameObject.Position[1] = float32(y64)
			return nil
		}
		return fmt.Errorf("%v attribute of %T expects a float32", attr, gameObject)
	case "positionAddX":
		x, ok := value.(float32)
		if ok {
			gameObject.Position[0] += x
			return nil
		}
		x64, ok := value.(float64)
		if ok {
			gameObject.Position[0] += float32(x64)
			return nil
		}
		return fmt.Errorf("%v attribute of %T expects a float32", attr, gameObject)
	case "positionAddY":
		y, ok := value.(float32)
		if ok {
			gameObject.Position[1] += y
			return nil
		}
		y64, ok := value.(float64)
		if ok {
			gameObject.Position[1] = float32(y64)
			return nil
		}
		return fmt.Errorf("%v attribute of %T expects a float32", attr, gameObject)
	case "scaleX":
		x, ok := value.(float32)
		if ok {
			gameObject.Scale[0] = x
			return nil
		}
		x64, ok := value.(float64)
		if ok {
			gameObject.Scale[0] = float32(x64)
			return nil
		}
		return fmt.Errorf("%v attribute of %T expects a float32", attr, gameObject)
	case "scaleY":
		y, ok := value.(float32)
		if ok {
			gameObject.Scale[1] = y
			return nil
		}
		y64, ok := value.(float64)
		if ok {
			gameObject.Scale[1] = float32(y64)
			return nil
		}
		return fmt.Errorf("%v attribute of %T expects a float32", attr, gameObject)
	case "euler":
		r, ok := value.(float32)
		if ok {
			gameObject.SetEuler(r)
			return nil
		}
		r64, ok := value.(float64)
		if ok {
			gameObject.SetEuler(float32(r64))
			return nil
		}
		return fmt.Errorf("%v attribute of %T expects a float32", attr, gameObject)
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

func (gameObject *GameObject) keyboardAttr(key string) bool {
	return gameObject.GetKey(KeyboardAttr[key])
}

func (gameObject *GameObject) getAttr(attr string) (interface{}, error) {
	switch attr {
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
	case "name":
		return gameObject.Name, nil
	}

	if strings.HasPrefix(attr, "getKey") {
		return gameObject.keyboardAttr(attr[6:]), nil
	}

	return nil, fmt.Errorf("attribute %v not found in %T", attr, gameObject)
}

func (gameObject *GameObject) SetAttr(componentName string, attr string, value interface{}) error {
	// base component ?
	if componentName == "" {
		return gameObject.setAttr(attr, value)
	}

	// custom attrs ?
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
	// base component ?
	if componentName == "" {
		return gameObject.getAttr(attr)
	}

	// custom attrs ?
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
