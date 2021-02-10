package methods

import (
	"bytes"
	"fmt"
	"github.com/gSpera/morse"
)

type Operations struct{

}

func NewOperations() *Operations{
	return &Operations{}
}


//convet to Binary format, if the origin format is Morse then its transform to text first
func (op *Operations)ConvertToBinary(text string,origin string) (string){
	var result string
	if (origin == "MORSE") {
		text =  morse.ToText(text)
	}
	var buffer bytes.Buffer
	for _, runeValue := range text {
		fmt.Fprintf(&buffer, "%b", runeValue)
	}
	return fmt.Sprintf("%s", buffer.Bytes())
	return result

}

func (op *Operations)ConvertToText(text string,origin string) (string){
	var result string

	if (origin == "MORSE") {
		result =  morse.ToText(text)
	}
	return result

}


func (op *Operations)ConvertToMorse(text string,origin string) (string){
	var result string
	if (origin == "TEXT") {
		result =  morse.ToMorse(text)
	}
	return result

}





