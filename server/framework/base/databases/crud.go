package databases

//This provides an API for CRUD functions that the various implementations should follow
//The framework logic around disallowing deletes etc. managed here

type DB interface {
	CheckError(err error)
	Get(key string) ([]byte, error)
}
