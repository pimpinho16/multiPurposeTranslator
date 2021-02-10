package main

import (
	"fmt"
	"multiPurposeTranslator/methods"
	"multiPurposeTranslator/process"
)

func main (){
	fmt.Println("Hello")
	var service process.Service

	Operation := *methods.NewOperations()

	service = process.NewTranslatorService(Operation)

	result,err := service.TranslateMethod(".-","MORSE","BINARY")
	if err != nil {
		fmt.Println("Error trying convert text","hola")
	}
	fmt.Println(result)

}
