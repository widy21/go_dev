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
}
