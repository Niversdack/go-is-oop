package observer

type subject interface {
	registerObserver(o observer)
	removeObserver(o observer)
	notifyObservers()
}
