package main

import (
	"fmt"
	"math/rand"
)

type User struct {
	ID int
}

func (u User) print() {
	fmt.Println(u.ID)

}

func main() {
	var users = make([]int, 10)

	for i := range users {
		users[i] = rand.Int()
		fmt.Println("users", users)
	}

}
