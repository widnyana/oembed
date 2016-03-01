package oembed

import (
	"log"
	"reflect"
	"strconv"
)

// getType will return data type as a string.
func getType(data interface{}) string {
	if data == nil {
		// a hack to prevent panic when get null pointer.
		return ""
	}
	return reflect.TypeOf(data).String()
}

// toUInt64 will convert data with type "string" and "float64" to "uint64"
// todo: add more data type.
func toUInt64(data interface{}) uint64 {
	var err error
	var item int

	switch getType(data) {
	case "string":
		item, err = strconv.Atoi(data.(string))
		if err != nil {
			log.Print("invalid data. set to 0")
			item = 0
		}
	case "float64":
		if tmp, ok := data.(float64); ok {
			return uint64(tmp)
		}
		item = 0
	}

	return uint64(item)
}
