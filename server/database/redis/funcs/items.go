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

func CreateItem(key, name, description string, tags []string, createdBy string) error {

	//there is no uniqueness check required but the UI generates UUID and DB cannot have duplicate for that

	a := models.Item{
		Name:        name,
		Description: description,
		Tags:        tags,
		Status:      "active",
		CreatedBy:   createdBy,
	}

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
