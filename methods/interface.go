package methods

type operations interface{
	Convert(text string, origin string,destination string) (string error)
}
