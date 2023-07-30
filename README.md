# is

![Actions Status](https://github.com/madflojo/is/actions/workflows/tests.yaml/badge.svg?branch=main)
[![codecov](https://codecov.io/gh/madflojo/is/branch/main/graph/badge.svg?token=H9C9B6I0AS)](https://codecov.io/gh/madflojo/is)
[![Go Report Card](https://goreportcard.com/badge/github.com/madflojo/is)](https://goreportcard.com/report/github.com/madflojo/is)
[![Go Reference](https://pkg.go.dev/badge/github.com/madflojo/is.svg)](https://pkg.go.dev/github.com/madflojo/is)
[![license](https://img.shields.io/github/license/madflojo/is.svg?maxAge=2592000)](https://github.com/madflojo/is/LICENSE)

Package is provides a curated collection of common regular expressions for Go.

```go
func main() {
  if is.IPv4().MatchString("10.0.0.1") {
    // do stuff
  }
}
```

Regular expressions are hard, this package is here to make them easier.

## Contributing

If you would like to contribute, please fork the repo and send in a pull request. All contributions are welcome and appreciated.

## License

The Apache 2 License. Please see [LICENSE](LICENSE) for more information.
