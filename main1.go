package main

import (
	"fmt"
	"log"
)

func FindUser(name string) (*User, error) {
	user, err := findUserFromGroup(users, name)


	//nil は「何もない」「見つからない」「エラーなし」を表す。
	if err != nil {
		return nil, err
	}
	return user, nil
}

func main() {
	user, err := FindUser("Bob")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.name)
}