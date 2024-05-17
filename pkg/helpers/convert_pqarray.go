package helpers

import (
	"strconv"
	"strings"
)

func ConvertBytePqArrayToActualValue(byteSlice []uint8) []uint8 {
	// Convert the byte slice to a string
	str := string(byteSlice)

	// Remove the curly braces from the string
	str = strings.Trim(str, "{}")

	// Split the string by comma to get individual elements
	elements := strings.Split(str, ",")

	// Trim spaces from each element
	for i, element := range elements {
		elements[i] = strings.TrimSpace(element)
	}

	var res []uint8
	for _, element := range elements {
		// Convert the string to an integer
		num, _ := strconv.Atoi(element)
		res = append(res, uint8(num))

	}
	return res
}
