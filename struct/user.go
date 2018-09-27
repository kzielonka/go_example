package main

import "fmt"

type User struct {
  name string
  age uint
}

// to jest takie to_str, toString
func (u User) String() string {
  return fmt.Sprintf("User: %s %d", u.name, u.age)
}

// przykladowa metoda na strukturze
func (u User) sayHi() string {
  return fmt.Sprintf("%s says \"Hi :)\"", u.name)
}

// ta jest ciekawsza bo ma zmenic imie
func (u User) changeName(newName string) {
  u.name = newName
}

// no i tu jest wskaznik *, czemu?
func (u *User) changeName2(newName string) {
  u.name = newName
}
