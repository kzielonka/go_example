// ./package_example.go

package main

import "fmt"
import "github.com/kzielonka/go_example/struct/user_in_package"

func main() {
  fmt.Println("==== 1 ====")
  // możemy zadeklarować bartka z var (explicite typ)
  var bartek User
  bartek = User{ "bartek", 25 }
  fmt.Println(bartek)

  fmt.Println("==== 2 ====")
  // albo możemy zadeklarować tak, go sam zgaduje typ
  otherBartek := User { "Marcepan", 32 }
  fmt.Println(otherBartek)

  fmt.Println("==== 3 ====")
  // bartek mówi cześć, wykonujemy metode na strukturze
  fmt.Println(bartek.sayHi())
 
  fmt.Println("==== 4 ====")
  // sprobojmy zmienic imie Bartkowi
  bartek.changeName("Krzysiek")
  fmt.Println(bartek) // no i dupa :(, czemu ?

  fmt.Println("==== 5 ====")
  // dobra, druga proba
  bartek.changeName2("Basia")
  fmt.Println(bartek)

  fmt.Println("==== 6 ====")
  // po co było sie tak pałować ;)
  bartek.name = "Kobra"
  fmt.Println(bartek)

  fmt.Println("==== 7 ====")
  var wojtek GentleGuy

  // hehe :) panic
  // fmt.Println(wojtek.sayHi())

  wojtek = User{ "Wojtek", 30 }

  // no nie :/
  // fmt.Println(wojtek) 

  fmt.Println(wojtek.sayHi())

  // WTF ????
  wojtek.changeName("Tomek")
  fmt.Println(wojtek.sayHi())

  fmt.Println("==== 8 ====")
  // referenca :), to jest pojebane nie bede wam tlumaczyl ze struct to raz tak a raz srak
  wojtek = &FixedUser{ "Wojtek", 30 }

  fmt.Println(wojtek.sayHi())

  // nie ma juz WTF ????
  wojtek.changeName("Tomek")
  fmt.Println(wojtek.sayHi())


  fmt.Println("==== 9 ====")

  //externalUser := user_in_package.NewUser("External", 58)
  //fmt.Println(externalUser.sayHi())

  //wojtek = user_in_package.NewUser("Marcin", 10)
  var tomek ExternalGentleGuy

  tomek = user_in_package.NewFixedUser("Tomek", 27)
  fmt.Println(tomek.SayHi())
}
