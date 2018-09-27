// ./package_example.go

package main

import "fmt"
import "github.com/kzielonka/go_example/packages_example/subpackage"

func main() {
  fmt.Println(subpackage.SayHi("Bartek"))
}
