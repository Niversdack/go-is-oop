package inOutput

import (
	"fmt"
	"testing"
)

func Test_milk_Cost(t *testing.T) {

	t.Run("decorator", func(t *testing.T) {
		cc := ClassicCase("Any String")
		lc := LowerCase(cc)
		fmt.Println(lc)
		upc := UpCase(lc)
		fmt.Println(upc)

	})

}
