package redisdb

import (
	interfaces "github.com/flutter_go/framework/base"
)

func Create(data interfaces.Data, binary []byte) (string, error) {
	/** New Data is saved in Redis as follows:
	* 1. If PK provided then check if exists. Return error 'exists' else go to 2
	* 2. If no PK, generate PK and save the raw object in default version and return PK
	* 3. If has workflow then initiate async workflow, put state as Workflow
	* 4. Once workflow ends, then we add this object in its Type list, increment version counter
	**/
	//PK~o0 = binary data

	return "nil", nil
}

func Read(key string) ([]interfaces.Data, error) {
	// var raw, err = Get(key)
	// var data interfaces.Data = raw
	return nil, nil //data, err
}

func Update(data interfaces.Data) (string, error) {
	/**Data update is done in Redis as follows:
	* 1. If state not in workflow then save the raw object in the new version and return version Else return error (unless allowed to change in workflow)
	* 3. If has workflow then initiate async workflow, put state as Workflow
	* 4. Once workflow ends, then we add this object in its Type list, increment version counter and keep the latest object in default access
	**/
	//PK~version~o0 = binary data
	//PK~o0 = new binary data after the workflow
	return "nil", nil
}
