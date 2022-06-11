package redisfuncs

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/flutter_go/models"
)

func checkErr(err error) {
	if err != nil {
		log.Fatalf("Error occured with redis items functions: %v", err)
	}
}

func CreateItem(_type, id, name, description string, tags []string, createdBy string) error {

	//there is no uniqueness check required but the UI generates UUID and DB cannot have duplicate for that

	a := models.Item{
		Name:        name,
		Description: description,
		Tags:        tags,
		Status:      "active",
		CreatedBy:   createdBy,
		Type:        _type, //Optional add if key has a separator
	}
	key := fmt.Sprintf("%s%s", _type, id) //TODO: add separator if not saving type
	data, err := json.Marshal(a)
	checkErr(err)

	var yes bool
	yes, err = Exists(key)
	if yes {
		err = fmt.Errorf("item with this id already exists")
	} else {
		if err != nil {
			checkErr(err)
		}
		err = Set(key, data)
		checkErr(err)
	}

	//TODO: go through the tags array and spit them as events (Tag, ID). Asynchronously add the new tags to the DB
	//TODO: if the item is of type user or api then send an event that will later add it in PostGresDB
	return err
}

func GetItem(key string, all bool) (models.Item, []string, error) {
	var (
		a    models.Item
		data []byte
		keys []string
		err  error
	)

	if !all {
		data, err = Get(key)
	} else {
		key = "item*"
		keys, err = GetKeys(key)
	}

	if err != nil {
		err = fmt.Errorf("item with this id is not existing")
	} else if !all {
		err = json.Unmarshal(data, &a)
		checkErr(err)
	}
	return a, keys, err
}
