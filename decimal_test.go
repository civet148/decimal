package decimal

import (
	"fmt"
	"testing"
)

func TestNewDecimal(t *testing.T) {
	dec := NewDecimal(623.3234, 3)
	dec = dec.Add("2993.325539923") //.Mul(10).Div(5)
	fmt.Printf("result: %v", dec.Float64())
}
