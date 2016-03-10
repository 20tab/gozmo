package gozmo

// A RegisteredComponent is a component that registers itself in the Engine
// so that data-based systems, like the JSON loader, can access it.
type RegisteredComponent struct {
	Name string
	// Called whenever a registered component is instantiated.
	Init func(args []interface{}) Component
}

type Component interface {
	Start(gameObject *GameObject)
	Update(gameObject *GameObject)
}

type ComponentAttr interface {
	SetAttr(attr string, value interface{}) error
	GetAttr(attr string) (interface{}, error)
}

type ComponentDestroy interface {
	Destroy(gameObject *GameObject)
}

type ComponentType interface {
	GetType() string
}
