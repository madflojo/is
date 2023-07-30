# is

![Actions Status](https://github.com/madflojo/is/actions/workflows/tests.yml/badge.svg?branch=main)
[![codecov](https://codecov.io/gh/madflojo/is/branch/main/graph/badge.svg?token=H9C9B6I0AS)](https://codecov.io/gh/madflojo/is)
[![Go Report Card](https://goreportcard.com/badge/github.com/madflojo/is)](https://goreportcard.com/report/github.com/madflojo/is)
[![Go Reference](https://pkg.go.dev/badge/github.com/madflojo/is.svg)](https://pkg.go.dev/github.com/madflojo/is)
[![license](https://img.shields.io/github/license/madflojo/is.svg?maxAge=2592000)](https://github.com/madflojo/is/LICENSE)

## Introduction

Package `is` provides a curated collection of common regular expressions for Go. It aims to simplify the process of 
validating common patterns, such as IP addresses, URLs, email addresses, and more, using easy-to-use regular expressions.

```go
package main

import (
  "github.com/madflojo/is"
)

func main() {
  // Validate an IPv4 address
  if is.IPv4().MatchString("10.0.0.1") {
    // do stuff
  }

  // Validate a web URL
  if is.URL().MatchString("https://example.com") {
    // do stuff
  }

  // Validate an email address
  if is.Email().MatchString("madflojo@example.com") {
    // do stuff
  }
}
```

Regular expressions can be complex and challenging to create for every edge case. The goal of this project is to provide a 
collection that covers most common cases, while acknowledging that not all possible edge cases are addressed. Contributions 
and improvements to the regular expressions are always welcome to enhance the package's reliability and usefulness.

## Contributing

If you would like to contribute, please fork the repo and send in a pull request. All contributions are welcome and 
appreciated.

## License

The Apache 2 License. Please see [LICENSE](LICENSE) for more information.
