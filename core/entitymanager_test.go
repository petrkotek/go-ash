package core

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestEntityManager_AddEntity(t *testing.T) {
	entity := NewEntity()

	em := NewEntityManager(&DummyEntityListener{})
	if err := em.AddEntity(entity, false); err != nil {
		t.Fatalf("AddEntity() expected to return nil error, got %v", err)
	}

	if !em.Entities().Has(entity) {
		t.Fatal("AddEntity() expected to add entity to the set")
	}

	if err := em.AddEntity(entity, false); err != ErrAlreadyRegistered {
		t.Fatalf("AddEntity() expected to return %v error, got %v", ErrAlreadyRegistered, err)
	}
}

func TestEntityManager_AddEntity_Delayed(t *testing.T) {
	entity := NewEntity()

	em := NewEntityManager(&DummyEntityListener{})
	em.AddEntity(entity, true)

	if em.Entities().Has(entity) {
		t.Fatal("AddEntity() expected to add entity only after calling ProcessPendingOperations()")
	}

	em.ProcessPendingOperations()
	if !em.Entities().Has(entity) {
		t.Fatal("AddEntity() expected to add entity after calling ProcessPendingOperations()")
	}
}

func TestEntityManager_RemoveEntity(t *testing.T) {
	entity := NewEntity()

	em := NewEntityManager(&DummyEntityListener{})
	em.AddEntity(entity, false)
	em.RemoveEntity(entity, false)

	if em.Entities().Has(entity) {
		t.Fatal("RemoveEntity() expected to remove entity from the set")
	}
}

func TestEntityManager_RemoveEntity_Delayed(t *testing.T) {
	entity := NewEntity()

	em := NewEntityManager(&DummyEntityListener{})
	em.AddEntity(entity, false)
	em.RemoveEntity(entity, true)

	if !em.Entities().Has(entity) {
		t.Fatal("RemoveEntity() expected to remove entity from the set only after calling ProcessPendingOperations()")
	}

	em.ProcessPendingOperations()

	if em.Entities().Has(entity) {
		t.Fatal("RemoveEntity() expected to remove entity from the set after calling ProcessPendingOperations()")
	}
}

func TestEntityManager_RemoveEntity_Delayed_DoubleRemove(t *testing.T) {
	entity := NewEntity()

	em := NewEntityManager(&DummyEntityListener{})
	em.AddEntity(entity, false)
	em.RemoveEntity(entity, true)
	em.RemoveEntity(entity, true)

	if !em.Entities().Has(entity) {
		t.Fatal("RemoveEntity() expected to remove entity from the set only after calling ProcessPendingOperations()")
	}

	em.ProcessPendingOperations()

	if em.Entities().Has(entity) {
		t.Fatal("RemoveEntity() expected to remove entity from the set after calling ProcessPendingOperations()")
	}
}

func TestEntityManager_AddEntity_CallsListener(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	entity := NewEntity()
	entityListenerMock := NewMockEntityListener(ctrl)
	entityListenerMock.EXPECT().OnEntityAdded(entity).Times(1)

	em := NewEntityManager(entityListenerMock)
	em.AddEntity(entity, false)
}

func TestEntityManager_RemoveEntity_CallsListener(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	entity := NewEntity()
	entityListenerMock := NewMockEntityListener(ctrl)
	entityListenerMock.EXPECT().OnEntityAdded(gomock.Any()).AnyTimes()
	entityListenerMock.EXPECT().OnEntityRemoved(entity).Times(1)

	em := NewEntityManager(entityListenerMock)
	em.AddEntity(entity, false)
	em.RemoveEntity(entity, false)
}

func TestEntityManager_RemoveEntity_DoesntCallListenerWhenEntityIsntInTheSet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	entity := NewEntity()
	entityListenerMock := NewMockEntityListener(ctrl)

	em := NewEntityManager(entityListenerMock)
	em.RemoveEntity(entity, false)
}
