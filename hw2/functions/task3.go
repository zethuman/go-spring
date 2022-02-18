package functions

import (
	"fmt"
)

var errMessage = func(param string, value float64) error {
	return fmt.Errorf("a %s of %0.2f is invalid", param, value)
}

func SmartPerimeter(l float64, w float64) (float64, error) {
	if l < 0 {
		return 0, errMessage("length", l)
	}
	if w < 0 {
		return 0, errMessage("width", w)
	}
	return l*2 + w*2, nil
}
