package _interface

type Operation interface {
	Traslator(text string, language string) (string, error)
}

type Process interface {
	GetProcessTraslator(text string, language string) (Operation, error)
}
