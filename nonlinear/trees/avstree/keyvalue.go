package avstree

type KeyValue interface {
	LessThan(KeyValue) bool
	EqualTo(KeyValue) bool
}