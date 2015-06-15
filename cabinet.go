package garoonclient

import (
	"encoding/base64"
	"encoding/xml"
	"io"
	"strings"

	"golang.org/x/text/transform"
)

const cabinetURL = "/cbpapi/cabinet/api"

// CabinetGetFolderInfo

type CabinetGetFolderInfoRequest struct {
	Header RequestHeader
}

func (r CabinetGetFolderInfoRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.Encode(buildRequestStruct(
		r.Header,
		"CabinetGetFolderInfo",
		struct{}{},
	))
}

type exclude struct {
	transform.NopResetter
	excluder func(byte) bool
}

func NewExclude(excluder func(byte) bool) transform.Transformer {
	return exclude{excluder: excluder}
}

func (e exclude) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	for nSrc = 0; nSrc < len(src); nSrc++ {
		b := src[nSrc]
		if !e.excluder(b) {
			if nDst >= len(dst) {
				err = transform.ErrShortDst
				return
			}
			dst[nDst] = b
			nDst++
		}
	}
	return
}

type CabinetGetFolderInfoResponse struct {
	Root CabinetFolder `xml:"returns>folder_information>root"`
}

type CabinetFolder struct {
	ID         string            `xml:"id,attr"`
	Title      string            `xml:"title"`
	ModifyTime string            `xml:"modify_time"`
	Folders    []*CabinetFolders `xml:"folders"`
	Location   []string          `xml:"-"`
}

type CabinetFolders struct {
	Folder []*CabinetFolder `xml:"folder"`
}

func (c *CabinetGetFolderInfoResponse) fillPath() {
	walkFn := func(f *CabinetFolder, parent *CabinetFolder) error {
		if parent != nil {
			f.Location = make([]string, len(parent.Location), len(parent.Location)+1)
			copy(f.Location, parent.Location)
		}
		f.Location = append(f.Location, f.Title)
		return nil
	}
	c.Root.Walk(walkFn, nil)
}

type CabinetFolderWalkFunc func(f *CabinetFolder, parent *CabinetFolder) error

func (f *CabinetFolder) Walk(walkFn CabinetFolderWalkFunc, parent *CabinetFolder) error {
	err := walkFn(f, parent)
	if err != nil {
		return err
	}
	for _, folders := range f.Folders {
		err := folders.walk(walkFn, f)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f CabinetFolders) walk(walkFn CabinetFolderWalkFunc, parent *CabinetFolder) error {
	for _, folder := range f.Folder {
		err := folder.Walk(walkFn, parent)
		if err != nil {
			return err
		}
	}
	return nil
}

func CabinetGetFolderInfo(r *CabinetGetFolderInfoRequest) (*CabinetGetFolderInfoResponse, error) {
	resp, err := sendRequest(r.Header.BaseURL+cabinetURL, r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return parseCabinetGetFolderInfoResponse(resp.Body)
}

func parseCabinetGetFolderInfoResponse(r io.Reader) (*CabinetGetFolderInfoResponse, error) {
	exclude := NewExclude(func(b byte) bool {
		return b == 0x08 || b == 0x0B
	})
	r2 := transform.NewReader(r, exclude)
	var resp CabinetGetFolderInfoResponse
	err := parseResponse(r2, "CabinetGetFolderInfoResponse", &resp)
	if err != nil {
		return nil, err
	}
	resp.fillPath()
	return &resp, err
}

// CabinetGetFileInfo

type CabinetGetFileInfoRequest struct {
	Header   RequestHeader
	FolderID string
}

func (r CabinetGetFileInfoRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.Encode(buildRequestStruct(
		r.Header,
		"CabinetGetFileInfo",
		struct {
			FolderID string `xml:"hid,attr"`
		}{
			FolderID: r.FolderID,
		},
	))
}

type CabinetGetFileInfoResponse struct {
	Files []CabinetFile `xml:"returns>file_information>files>file"`
}

type CabinetFile struct {
	ID         string `xml:"id,attr"`
	Title      string `xml:"title"`
	Name       string `xml:"name"`
	MimeType   string `xml:"mime_type"`
	ModifyTime string `xml:"modify_time"`
}

func CabinetGetFileInfo(r *CabinetGetFileInfoRequest) (*CabinetGetFileInfoResponse, error) {
	resp, err := sendRequest(r.Header.BaseURL+cabinetURL, r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return parseCabinetGetFileInfoResponse(resp.Body)
}

func parseCabinetGetFileInfoResponse(r io.Reader) (*CabinetGetFileInfoResponse, error) {
	exclude := NewExclude(func(b byte) bool {
		return b == 0x08 || b == 0x0B
	})
	r2 := transform.NewReader(r, exclude)
	var resp CabinetGetFileInfoResponse
	err := parseResponse(r2, "CabinetGetFileInfoResponse", &resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

// CabinetFileDownload

type CabinetFileDownloadRequest struct {
	Header RequestHeader
	FileID string
}

func (r CabinetFileDownloadRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.Encode(buildRequestStruct(
		r.Header,
		"CabinetFileDownload",
		struct {
			FileID string `xml:"file_id,attr"`
		}{
			FileID: r.FileID,
		},
	))
}

type CabinetFileDownloadResponse struct {
	Contents string `xml:"returns>file>content"`
}

func (r *CabinetFileDownloadResponse) ContentBytes() ([]byte, error) {
	return base64.StdEncoding.DecodeString(strings.TrimSpace(r.Contents))
}

func CabinetFileDownload(r *CabinetFileDownloadRequest) (*CabinetFileDownloadResponse, error) {
	resp, err := sendRequest(r.Header.BaseURL+cabinetURL, r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return parseCabinetFileDownloadResponse(resp.Body)
}

func parseCabinetFileDownloadResponse(r io.Reader) (*CabinetFileDownloadResponse, error) {
	var resp CabinetFileDownloadResponse
	err := parseResponse(r, "CabinetFileDownloadResponse", &resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}
