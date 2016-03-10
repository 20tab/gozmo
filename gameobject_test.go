package gozmo

import (
	"testing"
)

func TestInitialOrdering(t *testing.T) {
	scene := NewScene("Test")
	gameObject := scene.NewGameObject("Object")
	if gameObject.order != 0 {
		t.Error("Expected 0, got ", gameObject.order)
	}
}

func TestChangeOrdering(t *testing.T) {
	scene := NewScene("Test")
	gameObject := scene.NewGameObject("Object")
	gameObject.SetOrder(17)
	if gameObject.order != 17 {
		t.Error("Expected 17, got ", gameObject.order)
	}
	// Ensure that order mappings are available.
	if len(scene.orderedGameObjects[17]) != 1 {
		t.Error("Expected 1, got ", len(scene.orderedGameObjects[17]))
	}
	if len(scene.orderedGameObjects[0]) != 0 {
		t.Error("Expected 0, got ", len(scene.orderedGameObjects[0]))
	}
	// Change the order again.
	gameObject.SetOrder(30)
	if len(scene.orderedGameObjects[17]) != 0 {
		t.Error("Expected 0, got ", len(scene.orderedGameObjects[17]))
	}
	if len(scene.orderedGameObjects[30]) != 1 {
		t.Error("Expected 1, got ", len(scene.orderedGameObjects[30]))
	}
}
