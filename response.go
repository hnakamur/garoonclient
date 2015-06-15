package garoonclient

import (
	"encoding/xml"
	"errors"
	"io"
)

var ResponseTagNotFoundError = errors.New("response tag not found")

func parseResponse(r io.Reader, localName string, v interface{}) error {
	decoder := xml.NewDecoder(r)
	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == localName {
				return decoder.DecodeElement(v, &se)
			}
		}
	}
	return ResponseTagNotFoundError
}
