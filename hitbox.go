package gozmo

// The HitBox component is a generic AABB checker that can be used for basic
// collisions or area triggers. It can be attached to only one GameObject.
type HitBox struct {
	gameObject *GameObject
	xOffset    float32
	yOffset    float32
	width      float32
	height     float32
	raiseEvent string
}

var hitBoxes []*HitBox

func (hitbox *HitBox) Start(gameObject *GameObject) {
	if hitbox.gameObject != nil {
		panic("HitBox component cannot be attached to multiple GameObjects")
	}
	hitbox.gameObject = gameObject
}

func (hitbox *HitBox) Intersect(otherBox *HitBox) bool {
	return true
}

func (hitbox *HitBox) Update(gameObject *GameObject) {
	// For each hitbox (excluding myself), check for intersections and generate
	// events
	for _, hbox := range hitBoxes {
		if hbox == hitbox {
			continue
		}
		if hitbox.Intersect(hbox) {
			// hitboxes are colliding, raise events
			// on both objects
		}
	}
}

func (hitbox *HitBox) SetAttr(attr string, value interface{}) error {
	return nil
}

func (hitbox *HitBox) GetName() string {
	return "HitBox"
}

func (hitbox *HitBox) GetAttr(attr string) (interface{}, error) {
	return 0, nil
}

func NewHitBox(xOffset, yOffset, width, height float32) *HitBox {
	hitbox := HitBox{}
	hitbox.xOffset = xOffset
	hitbox.yOffset = yOffset
	hitbox.width = width
	hitbox.height = height
	hitBoxes = append(hitBoxes, &hitbox)
	return &hitbox
}

func initHitBox(args []interface{}) Component {
	if len(args) < 4 {
		panic("you need to specify the hitbox size")
	}
	// TODO: check for errors?
	x, _ := CastFloat32(args[0])
	y, _ := CastFloat32(args[1])
	width, _ := CastFloat32(args[2])
	height, _ := CastFloat32(args[3])
	return NewHitBox(x, y, width, height)
}

func init() {
	RegisterComponent("HitBox", initHitBox)
}
