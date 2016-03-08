package gozmo

// the HitBox component is a generic AABB checker
// can be used for basic collisions or area triggers

type HitBox struct {
	// an hitbox can be attached only to a single GameObject
	gameObject *GameObject
	xOffset    float32
	yOffset    float32
	width      float32
	height     float32
	raiseEvent string
}

func (hitbox *HitBox) Start(gameObject *GameObject) {
	if hitbox.gameObject != nil {
		panic("HitBox component cannot be attached to multiple GameObjects")
	}
	hitbox.gameObject = gameObject
}
func (hitbox *HitBox) Update(gameObject *GameObject) {}

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
	return &hitbox
}

func initHitBox(args []interface{}) Component {
	if len(args) < 4 {
		panic("you need to specify the hitbox size")
	}
	// TODO check for errors ?
	x, _ := CastFloat32(args[0])
	y, _ := CastFloat32(args[1])
	width, _ := CastFloat32(args[2])
	height, _ := CastFloat32(args[3])
	return NewHitBox(x, y, width, height)
}

func init() {
	RegisterComponent("HitBox", initHitBox)
}
