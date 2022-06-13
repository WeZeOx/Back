package utils

import (
	"Forum-Back-End/src/dto"
	"encoding/json"
)

func CheckFieldUser(user dto.User, array []string) bool {
	var structArray map[string]interface{}
	data, _ := json.Marshal(user)
	err := json.Unmarshal(data, &structArray)
	if err != nil {
		return false
	}

	for _, item := range array {
		if structArray[item] == "" || structArray[item] == nil {
			return false
		}
	}
	return true
}

func CheckFieldLogin(user dto.BodyLoginRequest, array []string) bool {
	var structArray map[string]interface{}
	data, _ := json.Marshal(user)
	err := json.Unmarshal(data, &structArray)
	if err != nil {
		return false
	}

	for _, item := range array {
		if structArray[item] == "" || structArray[item] == nil {
			return false
		}
	}
	return true
}

func CheckFieldPost(user dto.Post, array []string) bool {
	var structArray map[string]interface{}
	data, _ := json.Marshal(user)
	err := json.Unmarshal(data, &structArray)
	if err != nil {
		return false
	}
	for _, item := range array {
		if structArray[item] == "" || structArray[item] == nil {
			return false
		}
	}
	return true
}

func CheckFieldComment(user dto.ContentCommentCreator, array []string) bool {
	var structArray map[string]interface{}
	data, _ := json.Marshal(user)
	err := json.Unmarshal(data, &structArray)

	if err != nil {
		return false
	}
	for _, item := range array {
		if structArray[item] == "" || structArray[item] == nil {
			return false
		}
	}
	return true
}
