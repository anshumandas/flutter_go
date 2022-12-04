package redisfuncs

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/flutter_go/app/models"
	db "github.com/flutter_go/framework/base/databases/redis"
)

func checkAPIErr(err error) {
	if err != nil {
		log.Fatalf("Error occured with redis apis functions: %v", err)
	}
}

func CreateAPI(name, description string, children []string) error {
	//TODO: all subs should also get created as separate categories
	a := models.API{
		Name:        name,
		Description: description,
		Children:    children,
	}

	data, err := json.Marshal(a)
	checkAPIErr(err)

	var yes bool
	yes, err = db.Exists(name)
	if yes {
		err = fmt.Errorf("post with this title already exists")
	} else {
		if err != nil {
			checkAPIErr(err)
		}
		err = db.Set(name, data)
		checkAPIErr(err)
	}

	return err
}

func GetAPI(key string, all bool) (models.API, []string, error) {
	var (
		a    models.API
		data []byte
		keys []string
		err  error
	)

	if !all {
		data, err = db.Get(key)
	} else {
		key = "api*"
		keys, err = db.GetKeys(key)
	}

	if err != nil {
		err = fmt.Errorf("api with this name does not exists")
	} else if !all {
		err = json.Unmarshal(data, &a)
		checkAPIErr(err)
	}
	return a, keys, err
}

func DeleteAPI(key string) error {
	return db.Delete(key)
}
