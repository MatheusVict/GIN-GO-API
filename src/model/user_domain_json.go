package model

import (
	"encoding/json"
	"log"
)

func (ud *userDomain) GetJSONValue() (string, error) {
	sliceOfBytes, err := json.Marshal(ud)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(sliceOfBytes), nil
}
