package staircase

import (
	"fmt"
	"strings"
)

func Staircase(n int32) {
	for i := int32(0); i < n; i++ {
		fmt.Printf("%*s%s\n", n-1-i, "", strings.Repeat("#", int(i+1)))
	}
}
