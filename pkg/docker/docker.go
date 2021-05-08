package docker

import (
	"reflect"

	"github.com/docker/docker/api/types"
)

type RunningContainer struct {
	ContId    string
	ContName  []string
	ContState string
	ContPorts []types.Port
}

func StructContains(reference []RunningContainer, val RunningContainer) bool {
	value := false
	for _, item := range reference {
		if item == val {
			value = true
		} else {
			value = false
		}
	}
	return value
}

func StructEquality(reference RunningContainer, val RunningContainer) bool {
	value := false
	for i := 0; i < reflect.ValueOf(reference).Elem().Len(); i++ {
		if reference.Type().Field(i) == val.Type().Field(i) {
			value = true
		} else {
			value = false
		}
	}

	return value
}
