package dockercheck

import (
	"gomonclient/pkg/docker"
	"reflect"
)

func getStructureDiff(reference []docker.RunningContainer, currentState []docker.RunningContainer) (bool, string) {
	if reflect.DeepEqual(reference, currentState) {
		return true, "Structures are equal"
	} else {
		diagnosis := ""
		equality := false

		if len(reference) != len(currentState) {
			switch {
			case len(reference) > len(currentState):
				{
					for i, v := range reference{
					
					

					diagnosis = "Был остановлен контейнер"
					}
					}

				}
			case len(reference) < len(currentState):
				{
					diagnosis = "Был запущен контейнер"

				}
			}
		}

		return equality, diagnosis
	}

}
