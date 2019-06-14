package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Age  int
	Sex  bool
}

func main() {
	user := &User{
		Name: "user1",
		Age:  10,
		Sex:  false,
	}
	bytes, e := json.Marshal(user)
	if e != nil {
		fmt.Println("Marshal error, ", e)
	}
	fmt.Println(string(bytes))
	var user1 User
	err := json.Unmarshal(bytes, &user1)
	if err != nil {
		fmt.Errorf("json unmarshal error: %v", err)
	}
	fmt.Println("after unmarshal,user1 is: %v", user1)
}
