package starbuzz

type Beverage struct {
	description string
}
type Beverages interface {
	Description() string
	Cost() float64
}

func (b Beverage) Description() string {
	return ""
}
func (b Beverage) Cost() float64 {
	return 0
}
