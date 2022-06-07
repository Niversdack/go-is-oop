package observer

import (
	"math/rand"
)

type WeatherData struct {
	temperature int
	humidity    int
	pressure    int
	observers   *[]observer
}

func NewWeatherData() WeatherData {
	return WeatherData{
		observers: new([]observer),
	}
}

func (d WeatherData) GetTemperature() int {
	return rand.Int()
}
func (d WeatherData) GetHumidity() int {
	return rand.Int()
}
func (d WeatherData) GetPressure() int {
	return rand.Int()
}
func (d *WeatherData) RegisterObserver(o observer) {
	obs := d.observers
	*obs = append(*obs, o)
	d.observers = obs
}
func (d *WeatherData) RemoveObserver(o observer) {
	obs := d.observers
	for i, oo := range *obs {
		if oo == o {
			*obs = removeIndex(*obs, i)
		}
	}
	d.observers = obs
}

func removeIndex(s []observer, index int) []observer {
	return append(s[:index], s[index+1:]...)
}
func (d WeatherData) notifyObservers() {
	obs := d.observers
	for _, observer := range *obs {
		observer.Update(d.GetTemperature(), d.GetHumidity(), d.GetPressure())
	}
}
