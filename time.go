package is

import (
	"regexp"
)

const (
	ISO8601DateAndTimePattern = `^(?:[1-9]\d{3}-(?:(?:0[1-9]|1[0-2])-(?:0[1-9]|1\d|2[0-8])|(?:0[13-9]|1[0-2])-(?:29|30)|(?:0[13578]|1[02])-31)|(?:[1-9]\d(?:0[48]|[2468][048]|[13579][26])|(?:[2468][048]|[13579][26])00)-02-29)(T|\s)(?:[01]\d|2[0-3]):[0-5]\d:[0-5]\d(?:\.\d{1,9})?([zZ]|[+-][01]\d:[0-5]\d)?$`
	ISO8601DatePattern        = `^([0-9]{4})-?(1[0-2]|0[1-9])-?(3[01]|0[1-9]|[12][0-9])$`
	ISO8601TimePattern        = `^(2[0-3]|[01][0-9]):?([0-5][0-9]):?([0-5][0-9])$`
)

// ISO8601DateAndTime returns a *regexp.Regexp that matches time formats in ISO 8601 format.
func ISO8601DateAndTime() *regexp.Regexp {
	return regexp.MustCompile(ISO8601DateAndTimePattern)
}

// ISO8601Date returns a *regexp.Regexp that matches date formats in ISO 8601 format.
func ISO8601Date() *regexp.Regexp {
	return regexp.MustCompile(ISO8601DatePattern)
}

// ISO8601Time returns a *regexp.Regexp that matches time formats in ISO 8601 format.
func ISO8601Time() *regexp.Regexp {
	return regexp.MustCompile(ISO8601TimePattern)
}
