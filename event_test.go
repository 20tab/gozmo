package gozmo

import (
	"testing"
)

func TestEventProcessing(t *testing.T) {
	scene := NewScene("Test")
	gameObject001 := scene.NewGameObject("Object 1")
	gameObject002 := scene.NewGameObject("Object 2")

	gameObject001.EnqueueEvent(gameObject002, "test")

	if gameObject001.events == nil {
		t.Error("Expected not nil, got nil")
	}

	scene.Update(0)

	if gameObject001.events != nil {
		t.Error("Expected nil, got", gameObject001.events)
	}
}

func TestEventEnqueueing(t *testing.T) {
	scene := NewScene("Test")
	gameObject001 := scene.NewGameObject("Object 1")
	gameObject002 := scene.NewGameObject("Object 2")

	gameObject001.EnqueueEvent(gameObject002, "test")

	scene.Update(0)

	gameObject001.EnqueueEvent(gameObject002, "test")

	scene.Update(0)

	if gameObject001.events != nil {
		t.Error("Expected nil, got", gameObject001.events)
	}
}

type TestComponentForEvent struct {
	counter int32
}

func (tt *TestComponentForEvent) Start(gameObject *GameObject)  {}
func (tt *TestComponentForEvent) Update(gameObject *GameObject) {}
func (tt *TestComponentForEvent) OnEvent(gameObject *GameObject, event *Event) {
	tt.counter += 1
}

func TestEventByComponent(t *testing.T) {
	scene := NewScene("Test")
	gameObject001 := scene.NewGameObject("Object 1")
	gameObject002 := scene.NewGameObject("Object 2")

	component := &TestComponentForEvent{}

	gameObject001.AddComponent("test", component)

	gameObject001.EnqueueEvent(gameObject002, "test1")
	gameObject001.EnqueueEvent(gameObject002, "test2")
	gameObject001.EnqueueEvent(gameObject002, "test3")
	gameObject001.EnqueueEvent(gameObject002, "test4")
	gameObject001.EnqueueEvent(gameObject002, "test5")

	scene.Update(0)

	if component.counter != 5 {
		t.Error("Expected 5, got", component.counter)
	}
}
