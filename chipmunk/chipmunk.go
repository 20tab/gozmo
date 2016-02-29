package chipmunk

import (
	goz "github.com/20tab/gozmo"
	"github.com/vova616/chipmunk"
	"github.com/vova616/chipmunk/vect"
	_ "fmt"
)

var space *chipmunk.Space

var Gravity goz.Vector2 = goz.Vector2{0, -9.8}

func checkSpace() {
	if space == nil {
		space = chipmunk.NewSpace()
		space.Gravity = vect.Vect{vect.Float(Gravity[0]), vect.Float(Gravity[1])}
	}
}

type RigidBody struct {
	body *chipmunk.Body
	weight float32
}

func (rbody *RigidBody) Start(gameObject *goz.GameObject) {
	checkSpace()
	rbody.body = chipmunk.NewBody(vect.Float(rbody.weight), vect.Float(1))
	pos := gameObject.Position
	rbody.body.SetPosition(vect.Vect{vect.Float(pos[0]), vect.Float(pos[1])})
	space.AddBody(rbody.body)
}

func (rbody *RigidBody) Update(gameObject *goz.GameObject) {
	pos := rbody.body.Position()
	gameObject.SetPosition( float32(pos.X), float32(pos.Y))
}

func NewRigidBody(weight float32) goz.Component {
	body := RigidBody{weight: weight}
	return &body
}

func initRigidBody(args []interface{}) goz.Component {
        return NewRigidBody(1)
}

func updateWorld(scene *goz.Scene, deltaTime float32) {
	if space == nil {
		return
        }
	space.Step(vect.Float(deltaTime))
}

func init() {
	goz.RegisterComponent("RigidBody", initRigidBody)
	goz.RegisterUpdater(updateWorld)
}
