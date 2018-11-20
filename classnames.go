package classnames

import (
	"strings"
)

// Map of class names to whether they should be included or not
type Map = map[string]bool

// Aliases
var B = Build

// Takes any number of (arrays of) strings and class name maps
// Returns an HTML class="..." attribute string
func Build(args ...interface{}) string {
	var builder strings.Builder
	modArgCount := len(args) - 1

	for i, arg := range args {
		switch v := arg.(type) {
		case string:
			builder.WriteString(v)
		case []string:
			builder.WriteString(strings.Join(v, " "))
		case Map: // Include keys with the "true" value
			first := true

			for className, included := range v {
				if included {
					// Ensure no leading spaces
					if first {
						first = false
					} else {
						builder.WriteString(" ")
					}

					builder.WriteString(className)
				}
			}
		case []interface{}: // Recurse on all elements of the array
			builder.WriteString(Build(v...))
			continue
		default: // Unsupported type
			continue
		}

		// Ensure no leading spaces
		// Placed at the end to allow for "continue" to skip this
		if i < modArgCount {
			builder.WriteString(" ")
		}
	}

	return builder.String()
}
