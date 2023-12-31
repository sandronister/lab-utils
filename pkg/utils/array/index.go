package array

import (
	"errors"
	"fmt"
	"reflect"
)

func printSlice(key interface{}, list reflect.Value) (string, error) {
	for i := 0; i < list.Len(); i++ {
		if i == key {
			return fmt.Sprintf("%v", key), nil
		}
	}
	return "", errors.New("nenhuma chave encontrada")
}

func printMap(key interface{}, list reflect.Value) (string, error) {
	mapKeys := list.MapKeys()
	for inx, _ := range mapKeys {
		if inx == key {
			return fmt.Sprintf("%v", key), nil
		}
	}
	return "", errors.New("nenhuma chave encontrada")
}

func SearchKey[T interface{}](key interface{}, data T) (string, error) {
	list := reflect.ValueOf(data)

	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice:
		return printSlice(key, list)
	case reflect.Map:
		return printMap(key, list)
	default:
		return "", errors.New("nenhuma chave encontrada")
	}
}
