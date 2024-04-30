package regexutil

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// input: SomePrefix_value1_value2_...valueN - must be of this format & num of fields to extract
// output: value1 value2 ... valueN
// if no match found: (original input, error message)
func DropPrefixAndUnderscores(input string, numFields int) (string, error) {
	pattern := `^[A-z]*`

	for i := 1; i <= numFields; i++ {
		pattern += `_([^_]+?)`
		if i == numFields {
			pattern += `$`
		}
	}
	fmt.Printf("pattern: %s\n", pattern)

	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(input)

	// Error out if no match found and return original input
	if len(matches) == 0 {
		return input, errors.New("no match found")
	}

	extracted := strings.Join(matches[1:], " ")
	fmt.Printf("Extracted: %s\n", extracted)

	return extracted, nil
}
