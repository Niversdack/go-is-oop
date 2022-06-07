package factory

import (
	"fmt"
)

type Ordered interface {
	OrderPizza(interface{})
}
type pizza struct {
	name     string
	dough    string
	sauce    string
	toppings []string
}
type pizzaBuilder interface {
	Prepare()
	Bake()
	Cut()
	Box()
}

func (p pizza) Prepare() {
	//TODO implement me
	fmt.Println("Prepare pizza")
}

func (p pizza) Bake() {
	//TODO implement me
	fmt.Println("Baking...")
}

func (p pizza) Box() {
	//TODO implement me
	fmt.Println("i'm in box")
}
func (p pizza) Cut() {
	//TODO implement me
	fmt.Println("Cutiing \\/ 6 items ")
}
