package user_in_package

import "fmt"

type fixedUser struct {
  name string
  age uint
}

func NewFixedUser(name string, age uint) *fixedUser {
  return &fixedUser{ name, age }
}

func (u *fixedUser) SayHi() string {
  return fmt.Sprintf("%s says \"Hi :)\"", u.name)
}

func (u *fixedUser) ChangeName(newName string) {
  u.name = newName
}
