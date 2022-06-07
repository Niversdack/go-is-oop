package observer

import (
	"fmt"
	"testing"
)

func TestCCD_Display(t *testing.T) {

	t.Run("main", func(t *testing.T) {
		WD := NewWeatherData()
		NewCCD(&WD)
		FD := NewForecastDisplay(&WD)

		WD.notifyObservers()
		WD.RemoveObserver(FD)
		fmt.Println("space")
		WD.notifyObservers()

	})
}
