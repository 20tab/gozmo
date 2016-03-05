package gozmo

import (
	"testing"
)

func TestCageEscape(t *testing.T) {
	scene := NewScene("Test")
	gameObject := scene.NewGameObject("Object")
	gameObject.AddComponent("cage", NewCage(1, -1, -1, 1))
	gameObject.Position[0] = -2
	scene.Update(0)
	if gameObject.Position[0] != -1 {
		t.Error("Expected -1, got ", gameObject.Position[0])
	}
}
