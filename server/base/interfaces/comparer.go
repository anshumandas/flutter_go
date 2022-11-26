package interfaces

type Comparer interface {
	Compare(b Embedded) int
}
