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

func StringToJson(data string) map[string]interface{} {
	var jsonData map[string]interface{}
	json.Unmarshal([]byte(data), &jsonData)
	return jsonData
}
func ToIntList(data string) []int {
	var tmp = make([]int, 0)
	json.Unmarshal([]byte(data), &tmp)
	return tmp
}



