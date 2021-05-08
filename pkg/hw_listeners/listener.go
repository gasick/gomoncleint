package hw_listeners

import "time"

// интерфейс ответа от слушателя
type Responser interface {
	String() interface{}
}

// интерфейс слушателя
type Listener interface {
	getStatus() (Responser, error) // получить объект статуса
}

func getHWStatus(delay int) {
	for {
		CPUListener.getStatus()
		DiskListener.getStatus()
		RAMListener.getStatus()
		time.Sleep(delay * time.Second)
	}
}
