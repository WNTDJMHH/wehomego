package main

import "fmt"
import "flag"
import "reflect"

var strFunName string
type TestFunc struct{
} 

func (tf * TestFunc)SayHelloWord(){
	var strName string
	flag.StringVar(&strName, "i", "", "UserName")
	flag.Parse()
	fmt.Println("SayHelloWordFunction", strName);
}

func (tf * TestFunc)SayHelloWord2(){
	fmt.Println("SayHelloWordFunction2");
}

func main() {
	flag.StringVar(&strFunName, "f", "", "InputMethodName")
	flag.Parse()
	var testFunc TestFunc
	vauleReflect := reflect.ValueOf(&testFunc)
	vauleReflectType := vauleReflect.Type()
	fmt.Println(vauleReflectType.NumMethod())

	methodList := make([]string, 0, 50)		
	for i := 0; i < vauleReflect.NumMethod(); i++{
			fmt.Println("ToolsFunc", strFunName, "Method:", vauleReflectType.Method(i).Name);
			if strFunName == vauleReflectType.Method(i).Name{
				parms := []reflect.Value{}
				vauleReflect.Method(i).Call(parms)
				return ;
			}
			methodList = append(methodList, vauleReflectType.Method(i).Name)
	}
	flag.Usage()
	fmt.Println("MethodList:")
	for _, strMethod := range methodList{
		fmt.Println("\t", strMethod);
	}	
	
}
