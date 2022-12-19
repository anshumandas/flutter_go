package databases

import (
	interfaces "github.com/flutter_go/framework/base"
)

//This provides an API for CRUD functions that the various implementations should follow
//The framework logic around disallowing deletes etc. managed here

type DB interface {
	CheckError(err error)
	Create(data interfaces.Data) (string, error)
	Read(key string) ([]interfaces.Data, error) //one or more matching the key. Read all is also via the key
	Update(key string, data interfaces.Data) (int, error)
	Delete(key string) (interfaces.Data, error) //Marks inactive
}
