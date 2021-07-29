package utils

func Decode(item map[string]interface{}) map[string]interface{} {
	for k, v := range item {
		if b, ok := v.([]byte); ok {
			item[k] = string(b)
		}
		//item[k] = string(v) //v.decode("base64")
	}
	return item
}
