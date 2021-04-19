package binary

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/alwindoss/morse"
	"strings"
)

type binaryMorseImpl struct {
}

func NewBinaryMorseImpl(bm *binaryMorseImpl) *binaryMorseImpl {
	return &binaryMorseImpl{}
}

func (bm *binaryMorseImpl) Translator(text string, language string) (string, error) {
	h := morse.NewHacker()

	morseCode, err := h.Decode(strings.NewReader(text))
	if err != nil {
		return "", errors.New("No se pudo convertir morse a texto")
	}

	return toBinaryBytes(string(morseCode)), nil
}

func toBinaryBytes(s string) string {
	var buffer bytes.Buffer
	for i := 0; i < len(s); i++ {
		fmt.Fprintf(&buffer, "%b", s[i])
	}
	return fmt.Sprintf("%s", buffer.Bytes())
}
