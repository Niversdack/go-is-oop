package starbuzz

import (
	"fmt"
	"testing"
)

func Test_milk_Cost(t *testing.T) {

	t.Run("decorator", func(t *testing.T) {
		espresso := NewEspresso()
		espresso = WithMilk(espresso)
		fmt.Println(espresso.Description())
		espresso = WithMilk(espresso)
		fmt.Println(espresso.Description())

	})

}
