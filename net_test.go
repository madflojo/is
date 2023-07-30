package is

import (
	"fmt"
	"testing"
)

func TestIPv4(t *testing.T) {
	tt := []TestCases{
		{"Valid IPv4", []string{"10.0.0.0", "127.0.0.1", "0.0.0.0", "255.255.255.255"}, true},
		{"Invalid IPv4", []string{"256.256.256.256", "192.168.0.256", "1192.168.0.0"}, false},
		{"Not even an IP address", []string{"not an IP address"}, false},
		{"Mixed - IPv4 and IPv6 Addresses", []string{"192.168.0.1::1"}, false},
		{"Valid IPv4 - Leading Zeros in Octets", []string{"192.168.000.001"}, true},
		{"Invalid IPv4 - Leading Zeros in Octets", []string{"192.168.000.0010"}, false},
		{"Valid IPv4 - Trailing Zeros", []string{"192.168.0.000"}, true},
		{"Invalid IPv4 - Trailing Zeros", []string{"192.168.0.0000"}, false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for _, input := range tc.inputs {
				if IPv4().MatchString(input) && !tc.matches {
					t.Errorf("%s should not match", input)
				} else if !IPv4().MatchString(input) && tc.matches {
					t.Errorf("%s should match", input)
				}
			}
		})
	}
}

func ExampleIPv4() {
	if IPv4().MatchString("10.0.0.1") {
		fmt.Println("Valid IPv4")
	}
	// Output: Valid IPv4
}

func TestIPv6(t *testing.T) {
	tt := []TestCases{
		{"Valid IPv6 - Full Address", []string{"2001:0db8:85a3:0000:0000:8a2e:0370:7334"}, true},
		{"Valid IPv6 - Compressed Address", []string{"2001:db8:85a3::8a2e:370:7334"}, true},
		{"Valid IPv6 - Loopback Address", []string{"::1"}, true},
		{"Invalid IPv6", []string{"2001:0db8:85a3:0000:0000::8a2e:0370:7334:extra", "not an IPv6 address"}, false},
		{"IPv6 - Mixed Compression", []string{"2001:db8::1:0:0:1"}, true},
		{"IPv6 - Leading Zeros in Segments", []string{"2001:0db8:0000:0000:0000:8a2e:0370:7334"}, true},
		{"IPv6 - Trailing Zeros in Segments", []string{"2001:0db8:0:0:0:8a2e:0370:7334"}, true},
		{"IPv6 - Compressed Leading and Trailing Zeros", []string{"2001:db8::8a2e:0:7334"}, true},
		{"Invalid IPv6 - Double Colon in Wrong Position", []string{"2001:db8:::8a2e:0370:7334"}, false},
		{"Invalid IPv6 - Multiple Compressed Segments", []string{"2001:::8a2e::7334"}, false},
		{"Invalid IPv6 - Incomplete Segments", []string{"2001:db8:85a3:0000:0000:8a2e:0370"}, false},
		{"Invalid IPv6 - Extra Characters", []string{"2001:db8:85a3:0000:0000:8a2e:0370:extra"}, false},
		{"Not even an IP address", []string{"not an IP address"}, false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for _, input := range tc.inputs {
				if IPv6().MatchString(input) && !tc.matches {
					t.Errorf("%s should not match", input)
				} else if !IPv6().MatchString(input) && tc.matches {
					t.Errorf("%s should match", input)
				}
			}
		})
	}
}

func ExampleIPv6() {
	if IPv6().MatchString("::1") {
		fmt.Println("Valid IPv6")
	}
	// Output: Valid IPv6
}

func TestMAC(t *testing.T) {
	tt := []TestCases{
		{"Valid MAC Address", []string{"00:1A:2B:3C:4D:5E", "11:22:33:AA:BB:CC", "aa:bb:cc:dd:ee:ff"}, true},
		{"Invalid MAC Address", []string{"00:1A:2B:3C:4D", "00:1A:2B:3C:4D:5E:6F"}, false},
		{"Not even a MAC address", []string{"not a MAC address"}, false},
		{"Mixed Case MAC Address", []string{"00:1a:2b:3c:4d:5e", "11:22:33:aa:bb:cc", "aA:bB:cC:dD:eE:fF"}, true},
		{"MAC Address with Extra Spaces", []string{"00:1A :2B:3C:4D:5E", " 11:22:33:AA:BB:CC", "aa:bb:cc:dd:ee :ff"}, false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for _, input := range tc.inputs {
				if MAC().MatchString(input) && !tc.matches {
					t.Errorf("%s should not match", input)
				} else if !MAC().MatchString(input) && tc.matches {
					t.Errorf("%s should match", input)
				}
			}
		})
	}
}

func ExampleMAC() {
	if MAC().MatchString("00:1A:2B:3C:4D:5E") {
		fmt.Println("Valid MAC Address")
	}
	// Output: Valid MAC Address
}

func TestIPv4WithPort(t *testing.T) {
	tt := []TestCases{
		{"Valid IPv4 with Port", []string{"192.168.0.1:8080", "10.0.0.1:12345", "0.0.0.0:80"}, true},
		{"Invalid IPv4 with Port", []string{"192.168.0.1", "256.256.256.256:8080"}, false},
		{"Not even an IP address", []string{"not an IP address"}, false},
		{"IPv4 with Port - Non-Numeric Port", []string{"192.168.0.1:port", "10.0.0.1:http", "0.0.0.0:xyz"}, false},
		{"IPv4 with Port - IPv6-Style Enclosed Address", []string{"[192.168.0.1]:8080", "[10.0.0.1]:12345", "[0.0.0.0]:80"}, false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for _, input := range tc.inputs {
				if IPv4WithPort().MatchString(input) && !tc.matches {
					t.Errorf("%s should not match", input)
				} else if !IPv4WithPort().MatchString(input) && tc.matches {
					t.Errorf("%s should match", input)
				}
			}
		})
	}
}

func ExampleIPv4WithPort() {
	if IPv4WithPort().MatchString("10.0.0.1:9000") {
		fmt.Println("Valid IPv4 with Port")
	}
  // Output: Valid IPv4 with Port
}

func TestIPv6WithPort(t *testing.T) {
	tt := []TestCases{
		{"Valid IPv6 with Port", []string{"[2001:db8::1]:8080", "[::1]:12345", "[2001:0db8:85a3::8a2e:0370:7334]:80"}, true},
		{"Invalid IPv6 with Port", []string{"2001:db8::1", "[::1]:80a"}, false},
		{"Not even an IP address", []string{"not an IP address"}, false},
		{"IPv6 with Port - Non-Numeric Port", []string{"[2001:db8::1]:port", "[::1]:http", "[2001:0db8:85a3::8a2e:0370:7334]:xyz"}, false},
		{"IPv6 with Port - IPv4-Style Address", []string{"192.168.0.1:8080", "10.0.0.1:12345", "0.0.0.0:80"}, false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for _, input := range tc.inputs {
				if IPv6WithPort().MatchString(input) && !tc.matches {
					t.Errorf("%s should not match", input)
				} else if !IPv6WithPort().MatchString(input) && tc.matches {
					t.Errorf("%s should match", input)
				}
			}
		})
	}
}

func ExampleIPv6WithPort() {
	if IPv6WithPort().MatchString("[::1]:8080") {
		fmt.Println("Valid IPv6 with Port")
	}
	// Output: Valid IPv6 with Port
}

func TestURL(t *testing.T) {
	tt := []TestCases{
		{"Valid URL", []string{
			"http://www.example.com",
			"https://www.example.com",
			"ftp://example.com",
			"https://sub.example.com/page",
			"http://example.com/path/to/page.html?param=value",
			"https://www.example.com#section",
		}, true},
		{"Invalid URL", []string{
			"www.example.com",
			"https://",
			"http://example.com /page",
			"https://example.com\ttab",
			"https://example.com\nnewline",
		}, false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for _, input := range tc.inputs {
				if URL().MatchString(input) && !tc.matches {
					t.Errorf("%s should not match", input)
				} else if !URL().MatchString(input) && tc.matches {
					t.Errorf("%s should match", input)
				}
			}
		})
	}
}

func ExampleURL() {
	if URL().MatchString("https://www.example.com") {
		fmt.Println("Valid URL")
	}
	// Output: Valid URL
}
