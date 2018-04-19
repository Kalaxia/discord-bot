package utils

func CheckKeys(dict map[string]interface{}, keys ...string) bool {
	for key, _ := range dict {
		if _, ok := dict[key]; !ok {
			return false
		}
	}
	return true
}
