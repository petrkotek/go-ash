package core

type SystemManager struct {
	systems []EntitySystem
}

func NewSystemManager() *SystemManager {
	return &SystemManager{
		systems: []EntitySystem{},
	}
}

func (sm *SystemManager) Add(system EntitySystem) {
	sm.systems = append(sm.systems, system)
}

func (sm *SystemManager) Systems() []EntitySystem {
	return sm.systems
}
