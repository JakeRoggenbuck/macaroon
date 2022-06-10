# Macaroon - Go Macros & Preprocessor
Add C style macros to code in Go.

# Quick Start
Use the `set` keyword to define the value of a simple macro.

#### Before
```go
package main

set NAME "jake"

func main() {
	fmt.Println(NAME)
}
```

#### After
```go
package main

func main() {
	fmt.Println("jake")
}
```

# Running
```sh
macaroon main.mac.go && go build
```
