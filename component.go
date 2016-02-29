package gozmo

/*

A RegisteredComponent is a component that register itself
in the Engine so data-based systems (like the json loader)
can access it.

*/

type RegisteredComponent struct {
	Name string
	// this is called whenever a registered component is instantiated
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
