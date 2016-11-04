package core

import (
	"reflect"
	"testing"
)

type dummySystem struct{}

func (s *dummySystem) Update(timeDelta float32) error {
	return nil
}

func (s *dummySystem) CheckProcessing() bool {
	return true
}

func TestSystemManager_Systems_Empty(t *testing.T) {
	systemManager := NewSystemManager()
	systemsEmpty := systemManager.Systems()
	if kind := reflect.TypeOf(systemsEmpty).Kind().String(); kind != "slice" {
		t.Fatalf("Expected empty slice, got %v", kind)
	}

	if size := len(systemsEmpty); size != 0 {
		t.Fatalf("Expected slice with 0 items, got %d items", size)
	}
}

func TestSystemManager_Add(t *testing.T) {
	dummySystem1 := &dummySystem{}
	dummySystem2 := &dummySystem{}

	systemManager := NewSystemManager()

	// 1 system in the SystemManager
	systemManager.Add(dummySystem1)
	systemsOneItem := systemManager.Systems()
	if size := len(systemsOneItem); size != 1 {
		t.Fatalf("Expected slice with 1 item, got %d items", size)
	}
	if systemsOneItem[0] != dummySystem1 {
		t.Fatalf("Expected dummySystem1, got %v", systemsOneItem[0])
	}

	// 2 systems in the SystemManager
	systemManager.Add(dummySystem2)
	systemsTwoItems := systemManager.Systems()
	if size := len(systemsTwoItems); size != 2 {
		t.Fatalf("Expected slice with 2 items, got %d items", size)
	}
	if systemsTwoItems[0] != dummySystem1 {
		t.Fatalf("Expected dummySystem1, got %v", systemsTwoItems[0])
	}
	if systemsTwoItems[1] != dummySystem2 {
		t.Fatalf("Expected dummySystem2, got %v", systemsTwoItems[1])
	}
}
