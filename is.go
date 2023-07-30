/*
Package is provides a curated collection of regular expressions for Go.

This package does not replace the regexp package, but instead provides a
collection of regular expressions that are useful for validating input.

The design of this package is to return a regexp.Regexp object that can be
used directly as if you defined the regular expression yourself.

	// Validate an email address
	if is.Email().MatchString("email@example.com") {
	  fmt.Println("Valid email address")
	}

Regular expressions can be complex and challenging to create for every edge case.
The goal of this project is to provide a collection that covers most common cases,
while acknowledging that not all possible edge cases are addressed.
*/
package is
