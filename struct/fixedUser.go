package main

import "fmt"

type FixedUser struct {
  name string
  age uint
}

// przykladowa metoda na strukturze
func (u *FixedUser) sayHi() string {
  return fmt.Sprintf("%s says \"Hi :)\"", u.name)
}

// no i tu jest wskaznik *, czemu?
func (u *FixedUser) changeName(newName string) {
  u.name = newName
}
