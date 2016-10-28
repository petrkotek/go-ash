package core

import (
	"testing"

	"gopkg.in/fatih/set.v0"
)

func TestFamilyManager_EntitiesFor(t *testing.T) {
	entityA := NewEntity()
	entityA.Add(ComponentA{})
	entityB := NewEntity()
	entityB.Add(ComponentB{})
	entityAB := NewEntity()
	entityAB.Add(ComponentA{})
	entityAB.Add(ComponentB{})
	entityBA := NewEntity()
	entityBA.Add(ComponentB{})
	entityBA.Add(ComponentA{})
	entityNone := NewEntity()
	allEntities := []*Entity{entityA, entityB, entityAB, entityBA, entityNone}
	table := []struct {
		family         *Family
		entities       []*Entity
		expectedResult []*Entity
	}{

		{
			NewFamily(ComponentA{}),
			[]*Entity{},
			[]*Entity{},
		}, {
			NewFamily(ComponentA{}),
			[]*Entity{entityB},
			[]*Entity{},
		},
		{
			NewFamily(ComponentA{}),
			allEntities,
			[]*Entity{entityA, entityAB, entityBA},
		},
		{
			NewFamily(ComponentB{}),
			allEntities,
			[]*Entity{entityB, entityAB, entityBA},
		},
		{
			NewFamily(ComponentA{}, ComponentB{}),
			allEntities,
			[]*Entity{entityAB, entityBA},
		},
		{
			NewFamily(ComponentC{}),
			allEntities,
			[]*Entity{},
		},
	}
	for i, row := range table {
		entities := set.New()
		familyManager := NewFamilyManager(entities)
		for _, e := range row.entities {
			entities.Add(e)
		}
		result := familyManager.EntitiesFor(row.family)
		if !entitiesEqual(result, row.expectedResult) {
			t.Errorf("testcase_%d failed: Expected: %v Got: %v", i, row.expectedResult, result)
		}
	}
}

func TestFamilyManager_EntitiesFor_Removing(t *testing.T) {
	entityA := NewEntity()
	entityA.Add(ComponentA{})
	entityB := NewEntity()
	entityB.Add(ComponentB{})
	entityAB := NewEntity()
	entityAB.Add(ComponentA{})
	entityAB.Add(ComponentB{})

	family := NewFamily(ComponentA{})

	familyManager := NewFamilyManager(set.New())
	familyManager.EntitiesFor(family)

	familyManager.UpdateFamilyMembership(entityA, true)
	familyManager.UpdateFamilyMembership(entityB, true)
	familyManager.UpdateFamilyMembership(entityAB, true)

	familyManager.EntitiesFor(family)

	familyManager.UpdateFamilyMembership(entityAB, false)
	result := familyManager.EntitiesFor(family)

	expectedResult := []*Entity{entityA}
	if !entitiesEqual([]*Entity{entityA}, result) {
		t.Errorf("Expected: %v Got: %v", expectedResult, result)
	}
}
