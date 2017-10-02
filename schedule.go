package garoonclient

import (
	"encoding/xml"
	"io"
)

const scheduleURL = "/cbpapi/schedule/api"

// ScheduleGetEventsByTarget

type ScheduleGetEventsByTargetRequest struct {
	Header   RequestHeader
	Start    string
	End      string
	User     string
	Group    string
	Facility string
}

type baseID struct {
	ID string `xml:"id,attr"`
}

type ScheduleGetEventsByTargetResponse struct {
	Events []ScheduleEvent `xml:"returns>schedule_event"`
}

type ScheduleEvent struct {
	ID          string                  `xml:"id,attr"`
	EventType   string                  `xml:"event_type,attr"`
	PublicType  string                  `xml:"public_type,attr"`
	Detail      string                  `xml:"detail,attr"`
	Description string                  `xml:"description,attr"`
	Version     string                  `xml:"version,attr"`
	Timezone    string                  `xml:"timezone,attr"`
	EndTimezone string                  `xml:"end_timezone,attr"`
	AllDay      string                  `xml:"all_day,attr"`
	Members     []Member                `xml:"members>member"`
	DateTimes   []ScheduleEventDateTime `xml:"when>datetime"`
	Dates       []ScheduleEventDate     `xml:"when>date"`
}

type ScheduleEventDateTime struct {
	Start        string `xml:"start,attr"`
	End          string `xml:"end,attr"`
	FacilityCode string `xml:"facility_code,attr"`
}

type ScheduleEventDate struct {
	Start string `xml:"start,attr"`
	End   string `xml:"end,attr"`
}

type Member struct {
	User     User     `xml:"user,omitempty"`
	Group    Group    `xml:"group,omitempty"`
	Facility Facility `xml:"facility,omitempty"`
}

type User struct {
	ID   string `xml:"id,attr"`
	Name string `xml:"name,attr"`
}

type Group struct {
	ID   string `xml:"id,attr"`
	Name string `xml:"name,attr"`
}

type Facility struct {
	ID   string `xml:"id,attr"`
	Name string `xml:"name,attr"`
}

func (r ScheduleGetEventsByTargetRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	p := struct {
		Start    string  `xml:"start,attr"`
		End      string  `xml:"end,attr"`
		User     *baseID `xml:"user,omitempty"`
		Group    *baseID `xml:"group,omitempty"`
		Facility *baseID `xml:"facility,omitempty"`
	}{
		Start: r.Start,
		End:   r.End,
	}
	if r.User != "" {
		p.User = &baseID{ID: r.User}
	} else if r.Group != "" {
		p.Group = &baseID{ID: r.Group}
	} else if r.Facility != "" {
		p.Facility = &baseID{ID: r.Facility}
	}
	return e.Encode(buildRequestStruct(
		r.Header,
		"ScheduleGetEventsByTarget",
		p,
	))
}

func ScheduleGetEventsByTarget(r *ScheduleGetEventsByTargetRequest) (*ScheduleGetEventsByTargetResponse, error) {
	resp, err := sendRequest(r.Header.BaseURL+scheduleURL, r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return parseScheduleGetEventsByTargetResponse(resp.Body)
}

func parseScheduleGetEventsByTargetResponse(r io.Reader) (*ScheduleGetEventsByTargetResponse, error) {
	var resp ScheduleGetEventsByTargetResponse
	err := parseResponse(r, "ScheduleGetEventsByTargetResponse", &resp)
	if err != nil && err != ResponseTagNotFoundError {
		return nil, err
	}
	return &resp, nil
}
