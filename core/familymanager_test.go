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
		familyManager := NewFamilyManager(set.New())
		for _, e := range row.entities {
			familyManager.AddEntity(e)
		}
		result := familyManager.EntitiesFor(row.family)
		if !entitiesEqual(result, row.expectedResult) {
			t.Errorf("testcase_%d failed: Expected: %v Got: %v", i, row.expectedResult, result)
		}
	}
}
