package is

import "regexp"

const (
	EmailPattern                    = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	PhoneNumberNorthAmericaPattern  = `^([+]?\d{1,2}[-\.\s]?)?\(?(\d{3})\)?[-\.\s]?(\d{3})[-\.\s]?(\d{4})$`
	PhoneNumberInternationalPattern = `^([+]?\d{1,2}[-\.\s]?|)\d{3}[-\.\s]?\d{3}[-\.\s]?\d{4}$`
)

// Email returns a *regexp.Regexp that matches email addresses.
func Email() *regexp.Regexp {
	return regexp.MustCompile(EmailPattern)
}

// PhoneNumberNorthAmerica returns a *regexp.Regexp that matches North American phone numbers in various formats.
func PhoneNumberNorthAmerica() *regexp.Regexp {
	return regexp.MustCompile(PhoneNumberNorthAmericaPattern)
}

// PhoneNumberInternational returns a *regexp.Regexp that matches international phone numbers in various formats.
func PhoneNumberInternational() *regexp.Regexp {
	return regexp.MustCompile(PhoneNumberInternationalPattern)
}
