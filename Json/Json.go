package json

import "encoding/json"

func Marsh(str1 interface{}) []byte {
	bol, _ := json.Marshal(str1)
	return bol
}
func UnMarsh(byt []byte) interface{} {
	var dat interface{}
	err := json.Unmarshal(byt, &dat)
	if err != nil {
		panic(err)
	}
	return dat
}
