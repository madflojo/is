package is

import (
	"testing"
)

func TestISO8601DateAndTime(t *testing.T) {
	tests := []TestCases{
		{"Valid Date and Time Formats", []string{
			"2023-07-25T12:34:56",
			"2023-07-25T12:34:56Z",
			"2023-07-25T12:34:56+02:00",
			"2023-07-25 12:34:56",
			"2023-07-25 12:34:56Z",
			"2023-07-25 12:34:56+02:00",
		}, true},
		{"Invalid Date and Time Formats", []string{
			"2023-07-25T25:00:00", // Invalid hour (should be 00-23).
			"2023-07-25T00:60:00", // Invalid minute (should be 00-59).
			"2023-07-25T00:00:60", // Invalid second (should be 00-59).
			"2023-07-32T12:00:00", // Invalid day (should be 01-31).
			"2023-13-01T12:00:00", // Invalid month (should be 01-12).
		}, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for _, input := range tc.inputs {
				if ISO8601DateAndTime().MatchString(input) && !tc.matches {
					t.Errorf("%s should not match", input)
				} else if !ISO8601DateAndTime().MatchString(input) && tc.matches {
					t.Errorf("%s should match", input)
				}
			}
		})
	}
}

func TestISO8601Date(t *testing.T) {
	tests := []TestCases{
		{"Valid Date Formats", []string{
			"2023-07-25",
			"2023-07-01",
			"2023-12-31",
		}, true},
		{"Invalid Date Formats", []string{
			"2023-07-32",          // Invalid day (should be 01-31).
			"2023-13-01",          // Invalid month (should be 01-12).
			"2023-07-25T12:34:56", // Contains time components.
		}, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for _, input := range tc.inputs {
				if ISO8601Date().MatchString(input) && !tc.matches {
					t.Errorf("%s should not match", input)
				} else if !ISO8601Date().MatchString(input) && tc.matches {
					t.Errorf("%s should match", input)
				}
			}
		})
	}
}

func TestISO8601Time(t *testing.T) {
	tests := []TestCases{
		{"Valid Time Formats", []string{
			"12:34:56",
			"00:00:00",
			"23:59:59",
		}, true},
		{"Invalid Time Formats", []string{
			"24:00:00",       // Invalid hour (should be 00-23).
			"00:60:00",       // Invalid minute (should be 00-59).
			"12:34:60",       // Invalid second (should be 00-59).
			"12:34:56.789",   // Fractional seconds not allowed.
			"12:34:56Z",      // Contains UTC indicator.
			"12:34:56+02:00", // Contains UTC offset.
		}, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for _, input := range tc.inputs {
				if ISO8601Time().MatchString(input) && !tc.matches {
					t.Errorf("%s should not match", input)
				} else if !ISO8601Time().MatchString(input) && tc.matches {
					t.Errorf("%s should match", input)
				}
			}
		})
	}
}
