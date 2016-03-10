package gozmo

import (
	"testing"
)

func TestInitialOrdering(t *testing.T) {
	scene := NewScene("Test")
	gameObject := scene.NewGameObject("Object")
	if gameObject.order != 0 {
		t.Error("Expected 0, got", gameObject.order)
	}
}

func TestChangeOrdering(t *testing.T) {
	scene := NewScene("Test")
	gameObject := scene.NewGameObject("Object")
	gameObject.SetOrder(17)
	if gameObject.order != 17 {
		t.Error("Expected 17, got", gameObject.order)
	}
	// Ensure that order mappings are available.
	if len(scene.orderedGameObjects[17]) != 1 {
		t.Error("Expected 1, got", len(scene.orderedGameObjects[17]))
	}
	if len(scene.orderedGameObjects[0]) != 0 {
		t.Error("Expected 0, got", len(scene.orderedGameObjects[0]))
	}
	// Change the order again.
	gameObject.SetOrder(30)
	if len(scene.orderedGameObjects[17]) != 0 {
		t.Error("Expected 0, got", len(scene.orderedGameObjects[17]))
	}
	if len(scene.orderedGameObjects[30]) != 1 {
		t.Error("Expected 1, got", len(scene.orderedGameObjects[30]))
	}
}

func TestPositionX(t *testing.T) {
	scene := NewScene("Test")
	gameObject := scene.NewGameObject("Object")
	gameObject.Position[0] = 17
	value, _ := gameObject.GetAttr("", "positionX")
	if value.(float32) != 17 {
		t.Error("Expected 17, got", value)
	}
	gameObject.SetAttr("", "positionX", 30.0)
	if gameObject.Position[0] != 30 {
		t.Error("Expected 30, got", gameObject.Position[0])
	}
}

func TestPositionY(t *testing.T) {
	scene := NewScene("Test")
	gameObject := scene.NewGameObject("Object")
	gameObject.Position[1] = 17
	value, _ := gameObject.GetAttr("", "positionY")
	if value.(float32) != 17 {
		t.Error("Expected 17, got", value)
	}
	gameObject.SetAttr("", "positionY", 30.0)
	if gameObject.Position[1] != 30 {
		t.Error("Expected 30, got", gameObject.Position[1])
	}
}

func TestCustomValue(t *testing.T) {
	scene := NewScene("Test")
	gameObject := scene.NewGameObject("Object")
	gameObject.SetAttr("{}", "testKey", "testValue")
	value, _ := gameObject.GetAttr("{}", "testKey")
	if value.(string) != "testValue" {
		t.Error("Expected \"textValue\", got", value)
	}
}
