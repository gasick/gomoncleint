package docker

import (
	"fmt"
	"reflect"

	"github.com/docker/docker/api/types"
)

type RunningContainer struct {
	ContId    string
	ContName  []string
	ContState string
	ContPorts []types.Port
}

//Функция сравнения списка структур с эталаном.
func StructContains(reference []RunningContainer, val RunningContainer) bool {
	value := false
	for _, item := range reference {
		if StructEquality(item, val) {
			value = true
		} else {
			StructFieldEquality(item, val)
			value = false
		}
	}
	return value
}

// Функция сравнения полей стуктуры
func StructFieldEquality(reference RunningContainer, val RunningContainer) (bool, string) {
	value := false
	rStruct := reflect.ValueOf(reference)
	vStruct := reflect.ValueOf(val)
	rValues := make([]interface{}, rStruct.NumField())
	vValues := make([]interface{}, vStruct.NumField())

	for i := 0; i < rStruct.NumField(); i++ {
		if rValues[i] == vValues[i] {
			fmt.Println(rStruct.Field(i).Interface())
		}
	}

	return value, "test"
}

// Функция сравнения двух структур
func StructEquality(reference RunningContainer, val RunningContainer) bool {
	value := false
	rStruct := reflect.ValueOf(reference)
	vStruct := reflect.ValueOf(val)
	rValues := make([]interface{}, rStruct.NumField())
	vValues := make([]interface{}, vStruct.NumField())

	for i := 0; i < rStruct.NumField(); i++ {
		rValues[i] = rStruct.Field(i).Interface()
	}

	for i := 0; i < vStruct.NumField(); i++ {
		vValues[i] = vStruct.Field(i).Interface()
	}

	fmt.Println("rValues")
	fmt.Println(rValues)
	fmt.Println("vValues")
	fmt.Println(vValues)

	return value
}
