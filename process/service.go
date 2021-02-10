package process

import (
	"multiPurposeTranslator/methods"
)


type Service interface{
	TranslateMethod(text string, origin string, destination string) (string,error)
}

type ProcessService struct{
	Operations methods.Operations
}

func NewTranslatorService(Operations methods.Operations) *ProcessService{
	return &ProcessService{Operations: Operations}
}

func (s *ProcessService) TranslateMethod(text string, origin string, destination string) (string,error){
 	var result string
	if (destination =="BINARY"){
		result  =  s.Operations.ConvertToBinary(text,origin)
 	}
 	return result,nil
}

