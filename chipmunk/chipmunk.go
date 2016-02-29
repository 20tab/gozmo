package chipmunk

import (
	"fmt"
	goz "github.com/20tab/gozmo"
	"github.com/vova616/chipmunk"
	"github.com/vova616/chipmunk/vect"
)

var space *chipmunk.Space

var Gravity goz.Vector2 = goz.Vector2{0, -9.8}

func checkSpace() {
	if space == nil {
		space = chipmunk.NewSpace()
		space.Gravity = vect.Vect{vect.Float(Gravity[0]), vect.Float(Gravity[1])}
	}
}

// Rigid Body
type RigidBody struct {
	body        *chipmunk.Body
	weight      float32
	initialized bool
}

func (rbody *RigidBody) Start(gameObject *goz.GameObject) {
	checkSpace()
	rbody.body = chipmunk.NewBody(vect.Float(rbody.weight), vect.Float(1))
	space.AddBody(rbody.body)
}

func (rbody *RigidBody) Update(gameObject *goz.GameObject) {
	if !rbody.initialized {
		pos := gameObject.Position
		rbody.body.SetPosition(vect.Vect{vect.Float(pos[0]), vect.Float(pos[1])})
		rbody.initialized = true
	}
	pos := rbody.body.Position()
	gameObject.SetPosition(float32(pos.X), float32(pos.Y))
}

func (rbody *RigidBody) GetType() string {
	return "RigidBody"
}

func (rbody *RigidBody) SetAttr(attr string, value interface{}) error {
	switch attr {
	case "velocityX":
		x, _ := goz.CastFloat32(value)
		oldV := rbody.body.Velocity()
		rbody.body.SetVelocity(x, float32(oldV.Y))
	}
	return nil
}

func (rbody *RigidBody) GetAttr(attr string) (interface{}, error) {
	switch attr {
	case "velocityX":
		return float32(rbody.body.Velocity().X), nil
	case "velocityY":
		return float32(rbody.body.Velocity().Y), nil
	}
	return nil, fmt.Errorf("%v attribute of %T not found", attr, rbody)
}

func NewRigidBody(weight float32) goz.Component {
	body := RigidBody{weight: weight}
	return &body
}

func initRigidBody(args []interface{}) goz.Component {
	return NewRigidBody(1)
}

// Static Body
type StaticBody struct {
	body        *chipmunk.Body
	initialized bool
}

func (sbody *StaticBody) Start(gameObject *goz.GameObject) {
	checkSpace()
	sbody.body = chipmunk.NewBodyStatic()
	space.AddBody(sbody.body)
}

func (sbody *StaticBody) Update(gameObject *goz.GameObject) {
	if !sbody.initialized {
		pos := gameObject.Position
		sbody.body.SetPosition(vect.Vect{vect.Float(pos[0]), vect.Float(pos[1])})
		sbody.initialized = true
	}
	pos := sbody.body.Position()
	gameObject.SetPosition(float32(pos.X), float32(pos.Y))
}

func (sbody *StaticBody) GetType() string {
	return "StaticBody"
}

func NewStaticBody() goz.Component {
	body := StaticBody{}
	return &body
}

func initStaticBody(args []interface{}) goz.Component {
	return NewStaticBody()
}

// Shape Circle
type ShapeCircle struct {
	shape       *chipmunk.CircleShape
	initialized bool
}

func (circle *ShapeCircle) Start(gameObject *goz.GameObject) {
	component := gameObject.GetComponentByType("RigidBody")
	if component != nil {
		return
	}

	component = gameObject.GetComponentByType("StaticBody")
	if component != nil {
		return
	}

	fmt.Println("ShapeCircle requires a physic body")
}

func (circle *ShapeCircle) Update(gameObject *goz.GameObject) {
	if circle.initialized {
		return
	}

	circle.initialized = true

	component := gameObject.GetComponentByType("RigidBody")
	if component != nil {
		rbody := component.(*RigidBody)
		moment := circle.shape.Moment(float32(rbody.body.Mass()))
		rbody.body.SetMoment(moment)
		rbody.body.AddShape(circle.shape.Shape)
		space.AddShape(circle.shape.Shape)
		return
	}

	component = gameObject.GetComponentByType("StaticBody")
	if component != nil {
		sbody := component.(*StaticBody)
		sbody.body.AddShape(circle.shape.Shape)
		space.AddShape(circle.shape.Shape)
		return
	}
}

func (circle *ShapeCircle) SetAttr(attr string, value interface{}) error {
	switch attr {
	case "radius":
		radius, _ := goz.CastFloat32(value)
		circle.shape.Radius = vect.Float(radius)
		circle.shape.Shape.Update()
	}
	return nil
}

func (circle *ShapeCircle) GetAttr(attr string) (interface{}, error) {
	switch attr {
	case "radius":
		return float32(circle.shape.Radius), nil
	}
	return nil, fmt.Errorf("%v attribute of %T not found", attr, circle)
}

func NewShapeCircle() goz.Component {
	circle := ShapeCircle{}
	circle.shape = chipmunk.NewCircle(vect.Vector_Zero, 0).ShapeClass.(*chipmunk.CircleShape)
	return &circle
}

// TODO pass the radius as argument
func initShapeCircle(args []interface{}) goz.Component {
	return NewShapeCircle()
}

// Shape Box
type ShapeBox struct {
	shape       *chipmunk.BoxShape
	initialized bool
}

func (box *ShapeBox) Start(gameObject *goz.GameObject) {
	component := gameObject.GetComponentByType("RigidBody")
	if component != nil {
		return
	}

	component = gameObject.GetComponentByType("StaticBody")
	if component != nil {
		return
	}

	fmt.Println("ShapeBox requires a physic body")
}

func (box *ShapeBox) Update(gameObject *goz.GameObject) {
	if box.initialized {
		return
	}

	box.initialized = true

	component := gameObject.GetComponentByType("RigidBody")
	if component != nil {
		rbody := component.(*RigidBody)
		moment := box.shape.Moment(float32(rbody.body.Mass()))
		rbody.body.SetMoment(moment)
		rbody.body.AddShape(box.shape.Shape)
		space.AddShape(box.shape.Shape)
		return
	}

	component = gameObject.GetComponentByType("StaticBody")
	if component != nil {
		sbody := component.(*StaticBody)
		sbody.body.AddShape(box.shape.Shape)
		space.AddShape(box.shape.Shape)
		return
	}
}

func (box *ShapeBox) SetAttr(attr string, value interface{}) error {
	switch attr {
	case "width":
		w, _ := goz.CastFloat32(value)
		box.shape.Width = vect.Float(w)
		box.shape.UpdatePoly()
	case "height":
		h, _ := goz.CastFloat32(value)
		box.shape.Height = vect.Float(h)
		box.shape.UpdatePoly()
	}
	return nil
}

func (box *ShapeBox) GetAttr(attr string) (interface{}, error) {
	switch attr {
	case "width":
		return float32(box.shape.Width), nil
	case "height":
		return float32(box.shape.Height), nil
	}
	return nil, fmt.Errorf("%v attribute of %T not found", attr, box)
}

func NewShapeBox() goz.Component {
	box := ShapeBox{}
	box.shape = chipmunk.NewBox(vect.Vector_Zero, 0, 0).ShapeClass.(*chipmunk.BoxShape)
	return &box
}

// TODO pass width and height
func initShapeBox(args []interface{}) goz.Component {
	return NewShapeBox()
}

// this will be called at every world update
func updateWorld(scene *goz.Scene, deltaTime float32) {
	if space == nil {
		return
	}
	space.Step(vect.Float(deltaTime))
}

func init() {
	goz.RegisterComponent("RigidBody", initRigidBody)
	goz.RegisterComponent("StaticBody", initStaticBody)
	goz.RegisterComponent("ShapeCircle", initShapeCircle)
	goz.RegisterComponent("ShapeBox", initShapeBox)
	goz.RegisterUpdater(updateWorld)
}
