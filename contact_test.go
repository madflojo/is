package is

import (
	"fmt"
	"testing"
)

func TestPhoneNumberNorthAmerica(t *testing.T) {
	tt := []TestCases{
		{"Valid North American Phone Numbers", []string{
			"+12345678901",
			"1234567890",
			"1-234-567-8901",
			"1.234.567.8901",
			"1 234 567 8901",
		}, true},
		{"Invalid North American Phone Numbers", []string{
			"+1234567890123456", // Too many digits after country code.
			"12345678",          // Too few digits in the number.
			"123abc4567",        // Contains non-numeric characters.
		}, false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for _, input := range tc.inputs {
				if PhoneNumberNorthAmerica().MatchString(input) && !tc.matches {
					t.Errorf("%s should not match", input)
				} else if !PhoneNumberNorthAmerica().MatchString(input) && tc.matches {
					t.Errorf("%s should match", input)
				}
			}
		})
	}
}

func ExamplePhoneNumberNorthAmerica() {
	if PhoneNumberNorthAmerica().MatchString("1234567890") {
		fmt.Println("Valid North American Phone Number")
	}
	// Output: Valid North American Phone Number
}

func TestPhoneNumberInternational(t *testing.T) {
	tt := []TestCases{
		{"Valid International Phone Numbers", []string{
			"+44.1234567890",
			"+441234567890",
			"+12345678901",
		}, true},
		{"Invalid International Phone Numbers", []string{
			"+.1234567890",         // Missing country code.
			"+1..1234",             // Missing phone number.
			"+123.12345.678901234", // Too many digits in the number.
			"+12..34567",           // Double dots in the phone number.
			"+1.1234567890123456",  // Too many digits in the phone number.
			"+1.1234567890x",       // Missing extension digits.
			"+1.1234567890x-",      // Extension digits with invalid characters.
		}, false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for _, input := range tc.inputs {
				if PhoneNumberInternational().MatchString(input) && !tc.matches {
					t.Errorf("%s should not match", input)
				} else if !PhoneNumberInternational().MatchString(input) && tc.matches {
					t.Errorf("%s should match", input)
				}
			}
		})
	}
}

func ExamplePhoneNumberInternational() {
	if PhoneNumberInternational().MatchString("+441234567890") {
		fmt.Println("Valid International Phone Number")
	}
	// Output: Valid International Phone Number
}

func TestEmail(t *testing.T) {
	tt := []TestCases{
		{"Valid Email", []string{
			"example@example.com",
			"example+alias@example.com",
			"some.name@example.com",
			"email@subdomain.example.com",
			"123@example.com",
			"user@example.co",
			"user@example.museum",
			"user.name@example.com",
			"user-name@example.com",
			"user_name@example.com",
		}, true},
		// Existing test cases
		{"Invalid Email", []string{
			"example",
			"example@example",
			"example@",
			"example@example.",
			"@example.com",
			"user@example_co.com",
		}, false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for _, input := range tc.inputs {
				if Email().MatchString(input) && !tc.matches {
					t.Errorf("%s should not match", input)
				} else if !Email().MatchString(input) && tc.matches {
					t.Errorf("%s should match", input)
				}
			}
		})
	}
}

func ExampleEmail() {
	if Email().MatchString("example@example.com") {
		fmt.Println("Valid Email")
	}
	// Output: Valid Email
}
