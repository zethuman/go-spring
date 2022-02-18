package functions

import (
	"fmt"
	"strings"
)

const SIZE = 60

func Breaks(w string) {
	length := (SIZE - len(w)) / 2
	dashes := strings.Repeat("-", length)
	fmt.Printf("\n%s%s%s\n", dashes, w, dashes)
}
