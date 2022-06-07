package observer

import (
	"fmt"
)

type Displayer interface {
	Display()
}

type CCD struct {
	temp        int
	humid       int
	pressure    int
	WeatherData *WeatherData
}

func NewCCD(WeatherData *WeatherData) *CCD {
	ccd := &CCD{
		temp:        0,
		humid:       0,
		pressure:    0,
		WeatherData: WeatherData,
	}
	WeatherData.RegisterObserver(ccd)
	return ccd
}
func (c CCD) Display() {
	fmt.Println(c.humid, c.temp, c.pressure)
}
func (c *CCD) Update(temp, humidity, pressure int) {
	c.humid = humidity
	c.temp = temp
	c.pressure = pressure
	c.Display()
}

type ForecastDisplay struct {
	name string
	CCD
}

func NewForecastDisplay(WeatherData *WeatherData) *ForecastDisplay {
	fd := &ForecastDisplay{
		"forecast",
		CCD{
			WeatherData: WeatherData,
		},
	}
	WeatherData.RegisterObserver(fd)
	return fd
}
