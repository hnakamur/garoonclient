package garoonclient

import (
	"encoding/xml"
	"testing"
)

func TestMarshalCabinetGetFolderInfoRequest(t *testing.T) {
	header := RequestHeader{
		Username: "foo",
		Password: "password",
		Created:  "2010-08-12T14:45:00Z",
		Expires:  "2037-08-12T14:45:00Z",
		Locale:   "jp",
	}
	r := CabinetGetFolderInfoRequest{Header: header}
	want := `<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope"><soap:Header><Action>CabinetGetFolderInfo</Action><Security><UsernameToken><Username>foo</Username><Password>password</Password></UsernameToken></Security><Timestamp><Created>2010-08-12T14:45:00Z</Created><Expires>2037-08-12T14:45:00Z</Expires></Timestamp><Locale>jp</Locale></soap:Header><soap:Body><CabinetGetFolderInfo><parameters></parameters></CabinetGetFolderInfo></soap:Body></soap:Envelope>`
	gotBytes, err := xml.Marshal(r)
	if err != nil {
		t.Fatal("failed to marshal CabinetFolderListRequest")
	}
	got := string(gotBytes)
	if got != want {
		t.Errorf("marshal CabinetFolderListRequest unmatched. got: %s; want: %s", got, want)
	}
}

func TestMarshalCabinetGetFileInfoRequest(t *testing.T) {
	header := RequestHeader{
		Username: "foo",
		Password: "password",
		Created:  "2010-08-12T14:45:00Z",
		Expires:  "2037-08-12T14:45:00Z",
		Locale:   "jp",
	}
	r := CabinetGetFileInfoRequest{Header: header, FolderID: "253"}
	want := `<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope"><soap:Header><Action>CabinetGetFileInfo</Action><Security><UsernameToken><Username>foo</Username><Password>password</Password></UsernameToken></Security><Timestamp><Created>2010-08-12T14:45:00Z</Created><Expires>2037-08-12T14:45:00Z</Expires></Timestamp><Locale>jp</Locale></soap:Header><soap:Body><CabinetGetFileInfo><parameters hid="253"></parameters></CabinetGetFileInfo></soap:Body></soap:Envelope>`
	gotBytes, err := xml.Marshal(r)
	if err != nil {
		t.Fatal("failed to marshal CabinetFolderListRequest")
	}
	got := string(gotBytes)
	if got != want {
		t.Errorf("marshal CabinetFolderListRequest unmatched. got: %s; want: %s", got, want)
	}
}

func TestMarshalCabinetFileDownloadReques(t *testing.T) {
	header := RequestHeader{
		Username: "foo",
		Password: "password",
		Created:  "2010-08-12T14:45:00Z",
		Expires:  "2037-08-12T14:45:00Z",
		Locale:   "jp",
	}
	r := CabinetFileDownloadRequest{Header: header, FileID: "5369"}
	want := `<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope"><soap:Header><Action>CabinetFileDownload</Action><Security><UsernameToken><Username>foo</Username><Password>password</Password></UsernameToken></Security><Timestamp><Created>2010-08-12T14:45:00Z</Created><Expires>2037-08-12T14:45:00Z</Expires></Timestamp><Locale>jp</Locale></soap:Header><soap:Body><CabinetFileDownload><parameters file_id="5369"></parameters></CabinetFileDownload></soap:Body></soap:Envelope>`
	gotBytes, err := xml.Marshal(r)
	if err != nil {
		t.Fatal("failed to marshal CabinetFolderListRequest")
	}
	got := string(gotBytes)
	if got != want {
		t.Errorf("marshal CabinetFolderListRequest unmatched. got: %s; want: %s", got, want)
	}
}
