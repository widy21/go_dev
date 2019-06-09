package main

import (
	"fmt"
	"reflect"
)

type Stu struct {
	Name string
	Age  int
}

func main() {
	var x float32 = 2.3
	// var x1 float64 = 3.3
	value := reflect.ValueOf(x)
	kind := value.Kind()
	fmt.Println("value = ", value, " kind = ", kind)
	// var type1 reflect.Type = reflect.TypeOf(x)
	type2 := reflect.TypeOf(x)
	fmt.Println("type = ", type2)
	fmt.Printf("\n---begin set value---\n")
	fmt.Println("x = ", x)
	reflect.ValueOf(&x).Elem().SetFloat(5.5)
	fmt.Printf("\n---after set value---\n")
	fmt.Printf("x= %v\n", x)

	fmt.Println("value.Interface(): ", value.Interface())
	y := value.Interface().(float32)
	fmt.Println("after convert, y= ", y)
	fmt.Printf("\n######################\n")
	var stu Stu
	// kind =  struct（类别）
	// type =  main.Stu（具体类型）
	fmt.Println("value = ", reflect.ValueOf(stu), " kind = ", reflect.ValueOf(stu).Kind())
	fmt.Println("type = ", reflect.TypeOf(stu))

}
