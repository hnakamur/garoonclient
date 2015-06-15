package main

import (
	"flag"
	"io/ioutil"
	"os"

	"github.com/hnakamur/garoonclient"
)

var (
	fileID   string
	fileName string
)

func init() {
	flag.StringVar(&fileID, "fileid", "", "file ID")
	flag.StringVar(&fileName, "filename", "", "file name")
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
	req := garoonclient.CabinetFileDownloadRequest{Header: header, FileID: fileID}
	res, err := garoonclient.CabinetFileDownload(&req)
	if err != nil {
		panic(err)
	}
	b, err := res.ContentBytes()
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(fileName, b, 0644)
	if err != nil {
		panic(err)
	}
}
