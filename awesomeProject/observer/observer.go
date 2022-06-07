package observer

type observer interface {
	Update(temp, humidity, pressure int)
}
