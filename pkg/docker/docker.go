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

type Message string

//Функция сравнения списка структур с эталаном.
func StructContains(reference []RunningContainer, val RunningContainer) bool {
	fmt.Println("\n\tEntering StructContains")
	value := false
	for _, item := range reference {
		if StructEquality(item, val) {
			value = true
		}
	}
	return value
}

// Функция сравнения полей стуктуры
func StructFieldEquality(reference RunningContainer, val RunningContainer) (bool, string) {
	fmt.Println("\n\tEntering StructFiledEquality")
	value := false
	rStruct := reflect.ValueOf(reference)
	vStruct := reflect.ValueOf(val)
	rValues := make([]interface{}, rStruct.NumField())
	vValues := make([]interface{}, vStruct.NumField())

	for i := 0; i < rStruct.NumField(); i++ {
		if rValues[i] == vValues[i] {
			fmt.Printf("rStruct.Field(%d).Interface(): ", i)
			fmt.Println(rStruct.Field(i).Interface())
		}
	}

	return value, "test"
}

// Функция сравнения двух структур
func StructEquality(reference RunningContainer, val RunningContainer) bool {
	fmt.Println("\n\tEntering StructEquality")
	rStruct := reflect.ValueOf(reference)
	vStruct := reflect.ValueOf(val)
	rValues := make([]interface{}, rStruct.NumField())
	vValues := make([]interface{}, vStruct.NumField())

	fmt.Println("\n\t\t---***___")
	fmt.Println("\n\t\t---***___")

	fmt.Println("rValues:")
	fmt.Println(rValues)
	fmt.Println("vValues:")
	fmt.Println(vValues)

	return reflect.DeepEqual(vValues, rValues)
}
