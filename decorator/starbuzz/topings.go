package starbuzz

type milk struct {
	Beverages Beverages
}

func WithMilk(Beverages Beverages) Beverages {
	return milk{Beverages: Beverages}
}
func (m milk) Cost() float64 {
	return m.Beverages.Cost() + 30
}
func (m milk) Description() string {
	return m.Beverages.Description() + "with milk"
}
