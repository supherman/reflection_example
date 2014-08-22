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
	func(val interface{}) {
		if method := reflect.ValueOf(value).MethodByName(methodName); method.IsValid() {
			methodInterface := method.Interface()
			if m, ok := methodInterface.(func()); ok {
				m()
			}
		}
	}(value)
}

func main() {
	foo := Foo{}
	CallMethod(&foo, "SetName")
	log.Println(foo.Name)
}
