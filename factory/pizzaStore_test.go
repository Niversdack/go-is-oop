package factory

import (
	"fmt"
	"testing"
)

func Test_pizzaStore_OrderPizza(t *testing.T) {

	t.Run("test-factory", func(t *testing.T) {
		factory := pizzaFactory{}
		Store := NewPizzaStore(factory)
		Store.OrderPizza(NyStylePizza{})
		fmt.Println("=================")
		Store.OrderPizza(ChicagoStylePizza{})
	})

}
