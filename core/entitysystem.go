package core

type EntitySystem interface {
	CheckProcessing() bool
	Update(timeDelta float32) error
}
