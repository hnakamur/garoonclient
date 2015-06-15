package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hnakamur/garoonclient"
)

var logFileName string

func init() {
	flag.StringVar(&logFileName, "logfilename", "file_list.log", "log file name")
}

func main() {
	flag.Parse()

	logFile, err := os.Create(logFileName)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags)

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
	folderListReq := garoonclient.CabinetGetFolderInfoRequest{Header: header}
	folderListRes, err := garoonclient.CabinetGetFolderInfo(&folderListReq)
	if err != nil {
		panic(err)
	}
	_, err = fmt.Println(strings.Join(
		[]string{
			"hid",
			"location",
			"file_id",
			"title",
			"filename",
			"modify_time",
		},
		"\t"))
	if err != nil {
		panic(err)
	}
	walkFn := func(folder *garoonclient.CabinetFolder, parent *garoonclient.CabinetFolder) error {
		fileListReq := garoonclient.CabinetGetFileInfoRequest{Header: header, FolderID: folder.ID}
		fileListRes, err := garoonclient.CabinetGetFileInfo(&fileListReq)
		if err != nil {
			if err == garoonclient.ResponseTagNotFoundError {
				log.Printf("GetFileInfo response tag not found. hid=%s", folder.ID)
				// continue
				return nil
			}
			return err
		}
		for _, file := range fileListRes.Files {
			_, err = fmt.Println(strings.Join(
				[]string{
					folder.ID,
					strings.Join(folder.Location, " > "),
					file.ID,
					file.Title,
					file.Name,
					file.ModifyTime,
				},
				"\t"))
			return err
		}
		return nil
	}
	err = folderListRes.Root.Walk(walkFn, nil)
	if err != nil {
		panic(err)
	}
}
