package gozmo

// The Cage component constrains the position of gameObjects to a given area.
// Place it after the components you want to limit.
type Cage struct {
	top    float32
	left   float32
	bottom float32
	right  float32
}

func (cage *Cage) Start(gameObject *GameObject) {}
func (cage *Cage) Update(gameObject *GameObject) {
	if gameObject.Position[0] < cage.left {
		gameObject.Position[0] = cage.left
	}
	if gameObject.Position[0] > cage.right {
		gameObject.Position[0] = cage.right
	}
	if gameObject.Position[1] < cage.bottom {
		gameObject.Position[1] = cage.bottom
	}
	if gameObject.Position[1] > cage.top {
		gameObject.Position[1] = cage.top
	}
}

func (cage *Cage) SetAttr(attr string, value interface{}) error {
	return nil
}

func (cage *Cage) GetName() string {
	return "Cage"
}

func (cage *Cage) GetAttr(attr string) (interface{}, error) {
	return 0, nil
}

func NewCage(top, left, bottom, right float32) *Cage {
	cage := Cage{}
	cage.top = top
	cage.left = left
	cage.bottom = bottom
	cage.right = right
	return &cage
}

func initCage(args []interface{}) Component {
	if len(args) < 4 {
		panic("you need to specify the cage size")
	}
	// TODO: check for errors?
	top, _ := CastFloat32(args[0])
	left, _ := CastFloat32(args[1])
	bottom, _ := CastFloat32(args[2])
	right, _ := CastFloat32(args[3])
	return NewCage(top, left, bottom, right)
}

func init() {
	RegisterComponent("Cage", initCage)
}
