package handlers

import "gopkg.in/kataras/iris.v6"

type FooBar struct {
	Name string
	Age  int
	City string
}

func Foo(context *iris.Context) {
	foo := &FooBar{}
	if err := context.ReadJSON(foo); err != nil {
		context.Log(iris.DevMode, err.Error())
		return
	}

	context.Writef("FooBarBaz: %#v\n", foo)
}
