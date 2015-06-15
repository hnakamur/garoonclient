package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/hnakamur/garoonclient"
)

func main() {
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
	folderReq := garoonclient.CabinetGetFolderInfoRequest{Header: header}
	folderRes, err := garoonclient.CabinetGetFolderInfo(&folderReq)
	if err != nil {
		panic(err)
	}
	_, err = fmt.Printf("%s\t%s\n", "hid", "location")
	if err != nil {
		panic(err)
	}
	walkFn := func(f *garoonclient.CabinetFolder, parent *garoonclient.CabinetFolder) error {
		_, err := fmt.Printf("%s\t%s\n", f.ID, strings.Join(f.Location, " > "))
		return err
	}
	err = folderRes.Root.Walk(walkFn, nil)
	if err != nil {
		panic(err)
	}
}
