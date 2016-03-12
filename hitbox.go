package gozmo

import (
	_ "fmt"
)

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

	x1 := hitbox.gameObject.Position[0] + hitbox.xOffset*hitbox.gameObject.Scale[0]
	x1 -= hitbox.width * hitbox.gameObject.Scale[0] / 2
	y1 := hitbox.gameObject.Position[1] + hitbox.yOffset*hitbox.gameObject.Scale[1]
	y1 += hitbox.height * hitbox.gameObject.Scale[1] / 2

	w1 := x1 + hitbox.width*hitbox.gameObject.Scale[0]
	h1 := y1 - hitbox.height*hitbox.gameObject.Scale[1]

	x2 := otherBox.gameObject.Position[0] + otherBox.xOffset*otherBox.gameObject.Scale[0]
	x2 -= otherBox.width * otherBox.gameObject.Scale[0] / 2
	y2 := otherBox.gameObject.Position[1] + otherBox.yOffset*otherBox.gameObject.Scale[1]
	y2 += otherBox.height * otherBox.gameObject.Scale[1] / 2

	w2 := x2 + otherBox.width*otherBox.gameObject.Scale[0]
	h2 := y2 - otherBox.height*otherBox.gameObject.Scale[1]

	if w1 < x2 {
		return false
	}

	if h1 > y2 {
		return false
	}

	if x1 > w2 {
		return false
	}

	if y1 < h2 {
		return false
	}

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
			if hbox.raiseEvent != "" {
				gameObject.EnqueueEvent(hbox.gameObject, hbox.raiseEvent)
			}
			if hitbox.raiseEvent != "" {
				hbox.gameObject.EnqueueEvent(gameObject, hitbox.raiseEvent)
			}
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

func NewHitBoxWithEvent(xOffset, yOffset, width, height float32, event string) *HitBox {
	hitbox := NewHitBox(xOffset, yOffset, width, height)
	hitbox.raiseEvent = event
	return hitbox
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
	event := ""
	if len(args) > 4 {
		event = args[4].(string)
	}
	hitbox := NewHitBox(x, y, width, height)
	hitbox.raiseEvent = event
	return hitbox
}

func init() {
	RegisterComponent("HitBox", initHitBox)
}
