package hw_listeners

// интерфейс ответа от слушателя
type Responser interface {
	String() interface{}
}

// интерфейс слушателя
type Listener interface {
	getStatus() (Responser, error) // получить объект статуса
}
