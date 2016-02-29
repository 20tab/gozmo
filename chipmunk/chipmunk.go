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
	shape *chipmunk.Shape
}

func (circle *ShapeCircle) Start(gameObject *goz.GameObject) {
	component := gameObject.GetComponentByType("RigidBody")
	if component != nil {
		rbody := component.(*RigidBody)
		rbody.body.AddShape(circle.shape)
		space.AddShape(circle.shape)
		return
	}

	component = gameObject.GetComponentByType("StaticBody")
	if component != nil {
		sbody := component.(*StaticBody)
		sbody.body.AddShape(circle.shape)
		space.AddShape(circle.shape)
		return
	}

	fmt.Println("ShapeCircle requires a physic body")
}

func (circle *ShapeCircle) Update(gameObject *goz.GameObject) {
}

func NewShapeCircle() goz.Component {
	circle := ShapeCircle{}
	circle.shape = chipmunk.NewCircle(vect.Vector_Zero, 3)
	return &circle
}

func initShapeCircle(args []interface{}) goz.Component {
	return NewShapeCircle()
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
	goz.RegisterUpdater(updateWorld)
}
