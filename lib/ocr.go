package lib

import (
	"strings"
	"unicode"

	"github.com/otiai10/gosseract/v2"
)

// TesseractEngine contains configuration of tesseract client
type TesseractEngine struct {
	Name      string
	Language  string
	Variables map[string]string
	Client    *gosseract.Client
	ImageBytes []byte
}

func (e *TesseractEngine) Close() {
	e.Client.Close()
	e.Client = nil
	e.ImageBytes = nil
}

func (e *TesseractEngine) ExtractText() (string, error) {
	err := e.Client.SetImageFromBytes(e.ImageBytes)

	if err != nil {
		return "", err
	}

	output, err := e.Client.Text()
	if err != nil {
		return "", err
	}

	result := strings.TrimFunc(output, func(r rune) bool {
		return unicode.IsSpace(r)
	})

	return result, nil
}

func (e *TesseractEngine) TesseractSettings() {
	e.Client.SetLanguage(e.Language)
	for k, v := range e.Variables {
		e.Client.SetVariable(gosseract.SettableVariable(k), v)
	}
}
