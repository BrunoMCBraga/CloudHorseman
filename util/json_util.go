package util

import (
	"encoding/json"
	"errors"
)

func LoadJSONFileIntoMap(filePath string) (map[string]interface{}, error) {

	readFileResult, readFileError := ReadFile(filePath)
	if readFileError != nil {
		return nil, errors.New("|CloudHorseman->util->file_util->CreateTemporaryFile->util.ReadFile:" + readFileError.Error() + "|")
	}

	var jsonMap map[string]interface{}
	unmarshallErr := json.Unmarshal(readFileResult, &jsonMap)

	if unmarshallErr != nil {
		return nil, errors.New("|CloudHorseman->util->file_util->CreateTemporaryFile->util.ReadFile:" + unmarshallErr.Error() + "|")
	}

	return jsonMap, nil
}
