package main

/*
反射获取方法进行回调时，
方法名必须都要大写。
*/
import (
	"fmt"
	"reflect"
)

type NotknownType struct {
	Name string
	Age  int
	Sex  string
}

func (p NotknownType) GetNotknownType() string {
	return fmt.Sprintf("%s-%d-%s", p.Name, p.Age, p.Sex)
}

func (p NotknownType) Test() {
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaa")
}

var secret interface{} = NotknownType{"ethan", 666, "man"}

func main() {
	value := reflect.ValueOf(secret)
	kind := value.Kind()
	tp := reflect.TypeOf(secret)
	fmt.Println("value=", value)
	fmt.Println("kind=", kind)
	fmt.Println("tp=", tp)

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fmt.Println("field=", field)
	}
	var params []reflect.Value
	result := value.Method(0).Call(params)
	fmt.Println(result)
	fmt.Println("2th method................")
	value.MethodByName("Test").Call(nil)
}
