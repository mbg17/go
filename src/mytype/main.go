package main

import (
	"fmt"
	"reflect"
)

func reflectType(i interface{}) {
	obj := reflect.TypeOf(i)
	fmt.Printf("%s + %s\n", obj.Kind(), obj.Name())
	fmt.Printf("%T , %v\n", obj, obj)
}

func reflectValue(i interface{}) {
	obj := reflect.ValueOf(i)
	fmt.Printf("%T , %v\n", obj, obj)
	switch obj.Kind() {
	case reflect.Int:
		fmt.Println("值为", int(obj.Int()))
	case reflect.String:
		fmt.Println("值为", obj.String())
	case reflect.Struct:
		fmt.Println("值为", obj)
	}
}

func reflectSetValue(i interface{}) {
	obj := reflect.ValueOf(i)
	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	switch obj.Elem().Kind() {
	case reflect.Int:
		obj.Elem().SetInt(100)
	case reflect.String:
		obj.Elem().SetString("100")
	}
}

func main() {
	num := 10
	s := "haha"
	//reflectType(num)
	//reflectType(s)
	//type person struct {
	//	name string
	//	age  int
	//}
	//type book struct{ title string }
	//var d = person{
	//	name: "沙河小王子",
	//	age:  18,
	//}
	//var e = book{title: "《跟小王子学Go语言》"}
	//reflectType(d) // type:person kind:struct
	//reflectType(e) // type:book kind:struct
	//reflectValue(num)
	//reflectValue(s)
	//reflectValue(d)
	fmt.Println(num)
	fmt.Println(s)
	reflectSetValue(&num)
	reflectSetValue(&s)
	fmt.Println(num)
	fmt.Println(s)
}
