package pushover

import (
	"fmt"
	"strings"
)

// Errors represents the errors returned by pushover.
type Errors []string

// Error represents the error as a string.
func (e Errors) Error() string {
	ret := ""
	if len(e) > 0 {
		ret = fmt.Sprintf("Errors:\n")
		ret += strings.Join(e, "\n")
	}
	return ret
}
