package main

import (
	"fmt"
)

var classMap = make(map[string]*BuiltinClass)

type BuiltinClass struct {
	fields  map[string]interface{}
	methods map[string]func(...interface{})
}

type System struct {
	out PrintStream
}

type PrintStream struct {
}

type callerInfo struct {
	name  string
	field string
}

type calleeInfo struct {
	name   string
	method string
}

func initBuiltin() {
	system := BuiltinClass{fields: map[string]interface{}{"out": PrintStream{}}}
	printStream := BuiltinClass{
		methods: map[string]func(...interface{}){
			"println": func(args ...interface{}) {
				fmt.Println(args[1])
			},
		},
	}
	classMap["java/lang/System"] = &system
	classMap["java/io/PrintStream"] = &printStream
}

func callInstanceMethod(callerInfo callerInfo, calleeInfo calleeInfo, args interface{}) {
	caller := classMap[callerInfo.name].fields[callerInfo.field]
	method := classMap[calleeInfo.name].methods[calleeInfo.method]
	method(caller, args)
}
