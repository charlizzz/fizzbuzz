# Custom Fizzbuzz REST Server

This is a fully customizable fizzbuzz app using a Gin REST server with 5 parameters :
    - the first integer to be replaced by the "fizz"
    - the second integer to be replaced by the "buzz"
    - the limit of the count (integer)
    - the custom string that replace the "fizz"
    - the custom string that replace the "buzz"

## How to run on CLI

1. `go mod init fizzbuzz`

2. `go mod tidy`

3. `go get .`

4. `go run .`

5. `curl http://localhost:8080/custom-fizzbuzz/3/5/100/fizz/buzz`

## How to test

- main_test.go
  `go test`

## Ready for production

- Missing HELM or Kubes yaml deployment files
- Missing dockerfile
- Missing gin server settings
- Missing logging / tracing (with OTEL)
- Missing Versioning of endpoint [(semver.org)](https://semver.org/)
