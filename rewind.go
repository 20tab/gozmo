package gozmo

import (
	_ "fmt"
	"github.com/go-gl/mathgl/mgl32"
)

// The Rewind component restores the previous GameObject state
// when the specific event is raised
type Rewind struct {
	event       string
	oldPosition mgl32.Vec2
	oldRotation float32
	oldScale    mgl32.Vec2
}

func (rewind *Rewind) Start(gameObject *GameObject) {}
func (rewind *Rewind) Update(gameObject *GameObject) {
	rewind.oldPosition = gameObject.Position
	rewind.oldRotation = gameObject.Rotation
	rewind.oldScale = gameObject.Scale
}
func (rewind *Rewind) OnEvent(gameObject *GameObject, event *Event) {
	if event.Msg != rewind.event {
		return
	}

	gameObject.Position = rewind.oldPosition
	gameObject.Rotation = rewind.oldRotation
	gameObject.Scale = rewind.oldScale
}

func (rewind *Rewind) SetAttr(attr string, value interface{}) error {
	return nil
}

func (rewind *Rewind) GetName() string {
	return "Rewind"
}

func (rewind *Rewind) GetAttr(attr string) (interface{}, error) {
	return 0, nil
}

func NewRewind(event string) *Rewind {
	return &Rewind{event: event}
}

func initRewind(args []interface{}) Component {
	if len(args) < 1 {
		panic("you need to specify the rewind event")
	}
	return NewRewind(args[0].(string))
}

func init() {
	RegisterComponent("Rewind", initRewind)
}
