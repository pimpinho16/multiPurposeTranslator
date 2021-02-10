package process

import (
	"errors"
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
 	var err error = nil


	if (destination =="BINARY"){
		result  =  s.Operations.ConvertToBinary(text,origin)
 	}else if (destination == "TEXT"){
 		result = s.Operations.ConvertToText(text,origin)
	}else if (destination == "MORSE"){
		result = s.Operations.ConvertToMorse(text,origin)
	}else {
		err = errors.New("Format not recognized")
	}
 	return result,err
}

