package user_in_package

import "fmt"

type newUser struct {
  name string
  age uint
}

func NewUser(name string, age uint) *newUser {
  return &newUser{ name, age }
}

func (u *newUser) sayHi() string {
  return fmt.Sprintf("%s says \"Hi :)\"", u.name)
}

func (u *newUser) changeName(newName string) {
  u.name = newName
}
