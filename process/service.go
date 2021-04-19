package process

import (
	_interface "multiPurposeTranslator/interface"
)

type Service interface {
	TranslateMethod(text string, origin string, destination string) _interface.Operation
}

type ProcessService struct {
	BinaryFactory _interface.Operation
}

func NewTranslatorService(BinaryFactory _interface.Operation) *ProcessService {
	return &ProcessService{BinaryFactory: BinaryFactory}
}

func (s *ProcessService) TranslateMethod(text string, origin string, destination string) _interface.Operation {

	if destination == "BINARY" {
		return s.BinaryFactory.Traslator(text, origin)
	}

}
