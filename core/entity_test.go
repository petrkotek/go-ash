package core

import "testing"

func TestComponent(t *testing.T) {
	componentA := ComponentA{}
	componentB := ComponentB{}
	entityAB := NewEntity()
	entityAB.Add(componentA)
	entityAB.Add(componentB)

	c1 := entityAB.Component(ComponentA{})
	if c1 != componentA {
		t.Error()
	}

	c2 := entityAB.Component(ComponentB{})
	if c2 != componentB {
		t.Error()
	}

	c3 := entityAB.Component(ComponentC{})
	if c3 != nil {
		t.Error()
	}

	components := entityAB.Components()
	if size := components.Size(); size != 2 {
		t.Errorf("Components() expected to get 2 components, got %v components", size)
	}
	if !components.Has(componentA) {
		t.Error("Components() was expected to return componentA, but it didn't")
	}
	if !components.Has(componentB) {
		t.Error("Components() was expected to return componentA, but it didn't")
	}
}
