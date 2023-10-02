package main

import (
	"fmt"
	"syscall/js"
	"wasmwrapper/convert"
)

type User struct {
	Name    string
	UserId  int
	Address Address
}
type Address struct {
	Street string
	City   string
}
type MyContainer struct {
	Users []User
}

type MyCollector struct {
	Results []string
}

var (
	msgCollector MyCollector
)

func (m *MyCollector) addEntry(s string) {
	m.Results = append(m.Results, s)
}

func MyAdd(this js.Value, inputs []js.Value) interface{} {
	a := inputs[0].Int()
	b := inputs[1].Int()
	msgCollector.addEntry(fmt.Sprintf("internal log: %d + %d = %d", a, b, a+b))
	return a + b
}

func GetMessages(this js.Value, inputs []js.Value) interface{} {
	return convert.Convert(msgCollector.Results)
}

func GetDummyObject(this js.Value, inputs []js.Value) interface{} {
	sample := MyContainer{Users: []User{User{Name: "John", UserId: 1, Address: Address{Street: "Main Street", City: "New York"}}, User{Name: "Jane", UserId: 2, Address: Address{Street: "Main Street", City: "New York"}}}}
	return convert.Convert(sample)

}

func main() {
	msgCollector = MyCollector{Results: []string{"MyInitMessage"}}

	c := make(chan struct{}, 0)
	js.Global().Set("myAdd", js.FuncOf(MyAdd))
	js.Global().Set("myLogs", js.FuncOf(GetMessages))
	js.Global().Set("myGetDummyObject", js.FuncOf(GetDummyObject))

	<-c
}
