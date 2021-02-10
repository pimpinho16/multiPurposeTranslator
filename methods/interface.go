package methods

type operations interface{
	ConvertToBinary(text string,origin string) (string)
	ConvertToText(text string,origin string) (string)
	ConvertToMorse(text string, origin string) (string)
}
