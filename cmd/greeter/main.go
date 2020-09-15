package main

import (
  greeterv2 "github.com/chrisdoherty4/go-helloworld/v2/greeter"
  greeterv3 "github.com/chrisdoherty4/go-helloworld/v3/greeter"
)

func main() {
  greeterv2.Greet("Chris")
  greeterv3.Greet("Chris")
}
