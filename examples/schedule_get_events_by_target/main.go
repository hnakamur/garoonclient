package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/hnakamur/garoonclient"
)

func main() {
	now := time.Now()
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	todayEnd := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	start := flag.String("start", todayStart.Format(time.RFC3339), "search start datetime")
	end := flag.String("end", todayEnd.Format(time.RFC3339), "search end datetime")
	user := flag.String("user", "", "target user")
	group := flag.String("group", "", "target group")
	facility := flag.String("facility", "", "target facility")
	flag.Parse()

	baseURL := os.Getenv("BASE_URL")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	header := garoonclient.RequestHeader{
		BaseURL:  baseURL,
		Username: username,
		Password: password,
		Created:  "2010-08-12T14:45:00Z",
		Expires:  "2037-08-12T14:45:00Z",
		Locale:   "jp",
	}
	req := garoonclient.ScheduleGetEventsByTargetRequest{
		Header:   header,
		Start:    *start,
		End:      *end,
		User:     *user,
		Group:    *group,
		Facility: *facility,
	}
	resp, err := garoonclient.ScheduleGetEventsByTarget(&req)
	if err != nil {
		panic(err)
	}
	log.Printf("resp=%+v", resp)
}
