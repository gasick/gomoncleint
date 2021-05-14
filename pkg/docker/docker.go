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
func StructContains(reference []RunningContainer, val RunningContainer) (bool, string) {
	fmt.Println("\n\tEntering StructContains")
	difference := ""
	value := false
	for _, item := range reference {
		if StructEquality(item, val) {
			value = true
		} else {
			value, difference = StructFieldEquality(item, val)
			fmt.Printf("*** -- " + difference)
		}
	}
	return value, difference
}

// Функция сравнения полей стуктуры
func StructFieldEquality(reference RunningContainer, val RunningContainer) (bool, string) {
	fmt.Println("\n\tEntering StructFiledEquality")
	value := false
	message := ""
	rStruct := reflect.ValueOf(reference)
	vStruct := reflect.ValueOf(val)
	rValues := make([]interface{}, rStruct.NumField())
	vValues := make([]interface{}, vStruct.NumField())

	for i := 0; i < rStruct.NumField(); i++ {
		fmt.Println(i)
		if rValues[i] != vValues[i] {
			fmt.Println("---***---")
			fmt.Println(rStruct.Field(i))
			message = fmt.Sprintf("%s", rStruct.Field(i))
			fmt.Printf(message + "\n")
			fmt.Print("message inside: ")
			fmt.Println(message)
		}
	}

	fmt.Print("message outside: ")
	fmt.Println(message)

	return value, message
}

// Функция сравнения двух структур
func StructEquality(reference RunningContainer, val RunningContainer) bool {
	fmt.Println("\n\tEntering StructEquality")

	return reflect.DeepEqual(reference, val)
}
