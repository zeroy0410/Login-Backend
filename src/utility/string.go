package utility

import (
	// "fmt"
	"strconv"
)

func AtoU(str string) (uint, error) {
	i, err := strconv.Atoi(str)
	return uint(i), err
}
