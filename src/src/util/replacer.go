package util

import (
	"reflect"
	"strings"
)

func ReplaceSubstrings(message string, params map[string]string) string {
	for key, value := range params {
		message = strings.Replace(message, key, value, -1)
	}
	return message
}

func ReplaceString(prefix string, message string, obj interface{}) string {
	objValue := reflect.ValueOf(obj)
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}

	objType := reflect.TypeOf(objValue.Interface())
	newMessage := message

	for i := 0; i < objValue.NumField(); i++ {
		fieldValue := objValue.Field(i).String()
		fieldType := objType.Field(i).Name
		key := "{" + prefix + "." + strings.ToLower(fieldType) + "}"
		newMessage = strings.Replace(newMessage, key, fieldValue, -1)
	}

	return newMessage
}
