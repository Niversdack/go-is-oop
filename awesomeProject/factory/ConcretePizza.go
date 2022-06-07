package factory

import (
	"fmt"
)

type NyStylePizza struct {
	pizza
}

func NewNYStylePizza() pizzaBuilder {
	return NyStylePizza{pizza{
		name:     "NY STYLE",
		dough:    "Thin Crust Dough",
		sauce:    "Marinara Sauce",
		toppings: []string{"Grated Reggiano Cheese"},
	}}
}

type ChicagoStylePizza struct {
	pizza
}

func NewChicagoStylePizza() pizzaBuilder {
	return ChicagoStylePizza{pizza{
		name:     "Chicago STYLE",
		dough:    "Extra thick Crust Dough",
		sauce:    "Plum Tomato Sauce",
		toppings: []string{"Shredded Mozzarella Cheese"},
	}}
}
func (p ChicagoStylePizza) Cut() {
	fmt.Println("Cutting the pizza into square slice")
}
