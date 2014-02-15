package columnize

import (
	"fmt"
	"strings"
)

// Returns a list of elements, each representing a single item which will
// belong to a column of output.
func getElementsFromLine(line string, delim string) []interface{} {
	elements := make([]interface{}, 0)
	for _, field := range strings.Split(line, delim) {
		elements = append(elements, strings.TrimSpace(field))
	}
	return elements
}

// Examines a list of strings and determines how wide each column should be
// considering all of the elements that need to be printed within it.
func getWidthsFromLines(lines []string, delim string) []int {
	var widths []int

	for _, line := range lines {
		elems := getElementsFromLine(line, delim)
		i := 0
		for _, elem := range elems {
			if len(widths) <= i {
				widths = append(widths, len(elem.(string)))
			} else if widths[i] < len(elem.(string)) {
				widths[i] = len(elem.(string))
			}
			i++
		}
	}
	return widths
}

// Columnize is the public-facing interface that takes a list of strings and a
// delimiter, and returns nicely aligned output.
func Columnize(input []string, delim string) string {
	var stringfmt string
	var result string

	widths := getWidthsFromLines(input, delim)

	// Create the format string from the discovered widths
	for i := 0; i < len(widths); i++ {
		if i == len(widths)-1 {
			stringfmt += "%s\n"
		} else {
			stringfmt += fmt.Sprintf("%%-%ds", widths[i]+2)
		}
	}

	// Create the formatted output using the format string
	i := 0
	for _, line := range input {
		elems := getElementsFromLine(line, delim)
		for a := len(elems); a < len(widths); a++ {
			elems = append(elems, "")
		}
		result += fmt.Sprintf(stringfmt, elems...)
		i++
	}
	return strings.TrimSpace(result)
}
