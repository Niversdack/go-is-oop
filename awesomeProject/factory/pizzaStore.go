package factory

type pizzaStore struct {
	factory pizzaCreator
}

func NewPizzaStore(p pizzaCreator) Ordered {
	return pizzaStore{p}
}

func (s pizzaStore) OrderPizza(typePizza interface{}) {
	newPizza := s.factory.createPizza(typePizza)
	newPizza.Prepare()
	newPizza.Bake()
	newPizza.Cut()
	newPizza.Box()

}

type pizzaFactory struct {
}

type pizzaCreator interface {
	createPizza(typePizza interface{}) pizzaBuilder
}

func (p pizzaFactory) createPizza(typePizza interface{}) pizzaBuilder {
	switch typePizza.(type) {
	case NyStylePizza:
		return NewNYStylePizza()
	case ChicagoStylePizza:
		return NewChicagoStylePizza()
	}
	return pizza{}
}
