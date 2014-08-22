package main

import (
	"log"
	"reflect"
)

type Foo struct {
	Name string
}

func (foo *Foo) SetName() {
	foo.Name = "foo"
}

func CallMethod(value interface{}, methodName string) {
	if method := reflect.ValueOf(value).MethodByName(methodName); method.IsValid() {
		methodInterface := method.Interface()
		if m, ok := methodInterface.(func()); ok {
			m()
		}
	}
}

func main() {
	foo := Foo{}
	CallMethod(&foo, "SetName")
	log.Println(foo.Name)
}
