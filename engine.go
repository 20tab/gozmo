package gozmo

// An EngineSingleton is the one struct in the whole program that holds global
// structures like the list of registered components.
// TODO: possibly merge with window.go.
type EngineSingleton struct {
	Window               *Window
	registeredComponents map[string]*RegisteredComponent
	registeredUpdaters   []func(scene *Scene, deltaTime float32)
	scenes               map[string]*Scene
	perFrameStats        map[string]float64
}

var Engine EngineSingleton

func RegisterComponent(name string, generator func([]interface{}) Component) {
	// Create the map if required.
	if Engine.registeredComponents == nil {
		Engine.registeredComponents = make(map[string]*RegisteredComponent)
	}

	rc := RegisteredComponent{Name: name, Init: generator}

	Engine.registeredComponents[name] = &rc
}

func RegisterUpdater(updater func(scene *Scene, deltaTime float32)) {
	Engine.registeredUpdaters = append(Engine.registeredUpdaters, updater)
}
