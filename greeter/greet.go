package greeter

import "fmt"

const Version = 3

func Greet(name string) {
  fmt.Printf("Hello, %v, from v%v\n", name, Version)
}
