package interfaces

type BaseObject struct {
}

type Comparer interface {
	Compare(b BaseObject) int
}
