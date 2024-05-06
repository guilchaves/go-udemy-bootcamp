package main

import "fmt"

type user struct {
    name string
    age uint8 
}

func (u user) save(){
    fmt.Printf("Saving user %s into database\n", u.name)
}

func (u *user) makeBirthday() {
    u.age++
}

func main(){
    user1 := user{"User 1", 20}
    fmt.Println(user1)

    user1.save()

    fmt.Println(user1.age)
    user1.makeBirthday()
    fmt.Println(user1.age)

}
