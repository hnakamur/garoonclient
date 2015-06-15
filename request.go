package garoonclient

import (
	"bytes"
	"encoding/xml"
	"io"
	"net/http"
)

type envelope struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	Xmlns   string   `xml:"xmlns:soap,attr"`
	Header  header   `xml:"soap:Header"`
	Body    body     `xml:"soap:Body"`
}

type header struct {
	Action   string
	Username string `xml:"Security>UsernameToken>Username"`
	Password string `xml:"Security>UsernameToken>Password"`
	Created  string `xml:"Timestamp>Created"`
	Expires  string `xml:"Timestamp>Expires"`
	Locale   string
}

type body struct {
	Content bodyContent
}

type bodyContent struct {
	XMLName    xml.Name
	Parameters interface{} `xml:"parameters"`
}

type RequestHeader struct {
	BaseURL  string
	Username string
	Password string
	Created  string
	Expires  string
	Locale   string
}

func buildRequestStruct(h RequestHeader, apiName string, parameters interface{}) envelope {
	return envelope{
		Xmlns: "http://www.w3.org/2003/05/soap-envelope",
		Header: header{
			Action:   apiName,
			Username: h.Username,
			Password: h.Password,
			Created:  h.Created,
			Expires:  h.Expires,
			Locale:   h.Locale,
		},
		Body: body{
			Content: bodyContent{
				XMLName:    xml.Name{Local: apiName},
				Parameters: parameters,
			},
		},
	}
}

func sendRequest(url string, r interface{}) (*http.Response, error) {
	buf := bytes.NewBufferString(xml.Header)
	err := xml.NewEncoder(buf).Encode(r)
	if err != nil {
		return nil, err
	}
	return http.Post(url, "text/xml; charset=utf-8", buf)
}

func PrettyPrintRequest(w io.Writer, v interface{}) error {
	_, err := w.Write([]byte(xml.Header))
	if err != nil {
		return err
	}
	b, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("\n"))
	return err
}
