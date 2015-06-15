package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/hnakamur/garoonclient"
)

var folderID string

func init() {
	flag.StringVar(&folderID, "folderid", "", "folder ID")
}

func main() {
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
	fileListReq := garoonclient.CabinetGetFileInfoRequest{Header: header, FolderID: folderID}
	fileListRes, err := garoonclient.CabinetGetFileInfo(&fileListReq)
	if err != nil {
		panic(err)
	}
	for _, file := range fileListRes.Files {
		_, err = fmt.Println(strings.Join(
			[]string{
				file.ID,
				file.Title,
				file.Name,
				file.ModifyTime,
			},
			"\t"))
		if err != nil {
			panic(err)
		}
	}
}
