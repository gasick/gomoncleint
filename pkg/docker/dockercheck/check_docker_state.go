package dockercheck

import (
	"fmt"
	"gomonclient/pkg/docker"
	"reflect"
	"time"
)

func CheckDockerState(delay time.Duration) {
	var Reference []docker.RunContainers
	Reference = getDockerState()

	for {
		fmt.Println("\n\nChecking docker state:")
		if Reference != nil {
			var currentDockerState []docker.RunContainers
			currentDockerState = getDockerState()
			if reflect.DeepEqual(Reference, currentDockerState) {
				for i, v := range Reference {
					if v.ContId != currentDockerState[i].ContId || v.ContState != currentDockerState[i].ContState {
						fmt.Println("\tSomewhat izmenenia in state")
					} else {
						fmt.Println("\tVse is gut")
					}
				}
			} else {
				fmt.Println("\tSomewhat izmenenia in amount")
			}

		} else {
			fmt.Println("\tThe first one izmenenia")
		}
		Reference = getDockerState()
		time.Sleep(delay)
	}
}
