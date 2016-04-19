package core

import (
	"testing"

	"gopkg.in/fatih/set.v0"
)

func TestEntitiesFor(t *testing.T) {
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
		engine := NewEngine()
		for _, e := range row.entities {
			engine.AddEntity(e)
		}
		result := engine.EntitiesFor(row.family)
		if !entitiesEqual(result, row.expectedResult) {
			t.Errorf("testcase_%d failed: Expected: %V Got: %V", i, row.expectedResult, result)
		}
	}
}

func entitiesEqual(a, b []*Entity) bool {
	setA := set.New()
	for _, item := range a {
		setA.Add(item)
	}

	setB := set.New()
	for _, x := range b {
		setB.Add(x)
	}

	return set.SymmetricDifference(setA, setB).IsEmpty()
}
