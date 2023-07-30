package is

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
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
