package gozmo

/*

There is a single Engine struct in a program
it maintains global structures like the list
of registered components

*/

type EngineSingleton struct {
	registeredComponents map[string]*RegisteredComponent
}

var Engine EngineSingleton

func RegisterComponent(name string, generator func([]interface{}) Component) {
	// create the map if required
	if Engine.registeredComponents == nil {
		Engine.registeredComponents = make(map[string]*RegisteredComponent)
	}

	rc := RegisteredComponent{Name: name, Init: generator}

	Engine.registeredComponents[name] = &rc
}
