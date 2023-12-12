package maputil

import (
	"encoding/json"
)

// GetKeyFromMap golang通过map中的value获取key
func GetKeyFromMap(m map[string]string, value string) string {
	for k, v := range m {
		if v == value {
			return k
		}
	}
	return ""
}

func BytesToMap(bytes []byte) map[string]interface{} {
	m := make(map[string]interface{})
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		return nil
	}
	return m
}

func StructToMap(obj interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	j, err := json.Marshal(obj)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(j, &m)
	if err != nil {
		return nil
	}
	return m
}
