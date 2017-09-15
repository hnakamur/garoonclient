package garoonclient

import (
	"encoding/xml"
	"testing"
)

func TestMarshalScheduleGetEventsByTargetRequest(t *testing.T) {
	header := RequestHeader{
		Username: "foo",
		Password: "password",
		Created:  "2010-08-12T14:45:00Z",
		Expires:  "2037-08-12T14:45:00Z",
		Locale:   "jp",
	}
	r := ScheduleGetEventsByTargetRequest{
		Header: header,
		Start:  "2017-09-15T00:00:00",
		End:    "2017-09-15T23:59:59",
		Group:  "499",
	}
	want := `<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope"><soap:Header><Action>ScheduleGetEventsByTarget</Action><Security><UsernameToken><Username>foo</Username><Password>password</Password></UsernameToken></Security><Timestamp><Created>2010-08-12T14:45:00Z</Created><Expires>2037-08-12T14:45:00Z</Expires></Timestamp><Locale>jp</Locale></soap:Header><soap:Body><ScheduleGetEventsByTarget><parameters start="2017-09-15T00:00:00" end="2017-09-15T23:59:59"><group id="499"></group></parameters></ScheduleGetEventsByTarget></soap:Body></soap:Envelope>`
	gotBytes, err := xml.Marshal(r)
	if err != nil {
		t.Fatal("failed to marshal ScheduleGetEventsByTargetRequest")
	}
	got := string(gotBytes)
	if got != want {
		t.Errorf("marshal ScheduleGetEventsByTargetRequest unmatched.\ngot:\n%s!\nwant:\n%s!", got, want)
	}
}
