package main

import "fmt"

type Student struct {
	Name string
	Sex  bool
}

func Convert(b interface{}) {
	a, ok := b.(Student)
	if !ok {
		fmt.Println("convert false.")
		return
	}
	fmt.Println(a)
}

func JudgeType(items ...interface{}) {
	for index, value := range items {
		switch value.(type) {
		case bool:
			fmt.Println(index, " param is bool：", value)
		case int, int32, int64:
			fmt.Println(index, " param is int", value)
		case float32, float64:
			fmt.Println(index, " param is float", value)
		case string:
			fmt.Println(index, " param is string", value)
		case Student:
			fmt.Println(index, " param is Student", value)
		case *Student:
			fmt.Println(index, " param is *Student", value)
		default:
			fmt.Println(index, " param is not know", value)
		}
	}
}
func main() {
	// var a string = "abc"
	a := Student{
		Name: "dura",
		Sex:  false,
	}
	Convert(a)

	JudgeType(12, 3.3, "aaa", false, a, &a)
}
