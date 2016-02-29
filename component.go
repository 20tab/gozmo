package gozmo

type RegisteredComponent struct {
    Name string
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
