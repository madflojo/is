package is

import (
	"regexp"
)

const (
	// EmailPattern is the regular expression pattern for email addresses.
	EmailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
)

func Email() *regexp.Regexp {
	return regexp.MustCompile(EmailPattern)
}
