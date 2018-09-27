// ./subpackage/some_file.go

package subpackage

import "fmt"

func SayHi(name string) string {
  return fmt.Sprintf("Hi %s! :)", name)
}
