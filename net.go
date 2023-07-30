package is

import (
	"regexp"
)

const (
	IPv4Pattern         = `^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`
	IPv6Pattern         = `^([[:xdigit:]]{1,4}(?::[[:xdigit:]]{1,4}){7}|::|:(?::[[:xdigit:]]{1,4}){1,6}|[[:xdigit:]]{1,4}:(?::[[:xdigit:]]{1,4}){1,5}|(?:[[:xdigit:]]{1,4}:){2}(?::[[:xdigit:]]{1,4}){1,4}|(?:[[:xdigit:]]{1,4}:){3}(?::[[:xdigit:]]{1,4}){1,3}|(?:[[:xdigit:]]{1,4}:){4}(?::[[:xdigit:]]{1,4}){1,2}|(?:[[:xdigit:]]{1,4}:){5}:[[:xdigit:]]{1,4}|(?:[[:xdigit:]]{1,4}:){1,6}:)$`
	MACPattern          = `^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$`
	IPv4WithPortPattern = `^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?):\d{1,5}$`
	IPv6WithPortPattern = `^\[([[:xdigit:]]{1,4}(?::[[:xdigit:]]{1,4}){7}|::|:(?::[[:xdigit:]]{1,4}){1,6}|[[:xdigit:]]{1,4}:(?::[[:xdigit:]]{1,4}){1,5}|(?:[[:xdigit:]]{1,4}:){2}(?::[[:xdigit:]]{1,4}){1,4}|(?:[[:xdigit:]]{1,4}:){3}(?::[[:xdigit:]]{1,4}){1,3}|(?:[[:xdigit:]]{1,4}:){4}(?::[[:xdigit:]]{1,4}){1,2}|(?:[[:xdigit:]]{1,4}:){5}:[[:xdigit:]]{1,4}|(?:[[:xdigit:]]{1,4}:){1,6}:)\]:\d{1,5}$`
)

// IPv4 returns a *Regexp that matches IPv4 addresses.
func IPv4() *regexp.Regexp {
	return regexp.MustCompile(IPv4Pattern)
}

// IPv6 returns a *Regexp that matches IPv6 addresses.
func IPv6() *regexp.Regexp {
	return regexp.MustCompile(IPv6Pattern)
}

// MAC returns a *Regexp that matches MAC addresses.
func MAC() *regexp.Regexp {
	return regexp.MustCompile(MACPattern)
}

// IPv4WithPort returns a *Regexp that matches IPv4 addresses with port.
func IPv4WithPort() *regexp.Regexp {
	return regexp.MustCompile(IPv4WithPortPattern)
}

// IPv6WithPort returns a *Regexp that matches IPv6 addresses with port.
func IPv6WithPort() *regexp.Regexp {
	return regexp.MustCompile(IPv6WithPortPattern)
}
