package jsons

import "encoding/json"

func ToJsonStr(data map[string]interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jsonData)
}
func Struct2Json(data interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonData), err
}
