package staircase

import (
	"fmt"
	"strings"
)

func Staircase() {
	input := int32(6)

	for i := int32(0); i < input; i++ {
		fmt.Printf("%*s%s\n", input-1-i, "", strings.Repeat("#", int(i+1)))
	}
}
