/*
Package is provides a curated collection of regular expressions for Go.

This package does not replace the regexp package, but instead provides a
collection of regular expressions that are useful for validating input.

The design of this package is to return a regexp.Regexp object that can be
used directly as if you defined the regular expression yourself.

	// Validate an email address
	isEmail := is.Email()
	if isEmail.MatchString("email@example.com") {
	  fmt.Println("Valid email address")
	}
*/
package is
