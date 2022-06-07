package starbuzz

type Espresso struct {
}

func NewEspresso() Beverages {
	return Espresso{}
}
func (e Espresso) Description() string {
	return "Эспрессо"
}
func (e Espresso) Cost() float64 {
	return 200
}

type HouseBlend struct {
	Beverage
}

func NewHouseBlend() Beverages {
	return HouseBlend{}
}
func (e HouseBlend) Description() string {
	return "House Blend"
}
func (e HouseBlend) Cost() float64 {
	return 100
}
