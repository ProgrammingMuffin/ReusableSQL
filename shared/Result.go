package shared

type Error struct {
	Message string
}

type Result[V any] struct {
	Value   V
	Err     Error
	Errored bool
}
