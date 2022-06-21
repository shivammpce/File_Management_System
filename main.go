package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	files := getFileFromDownload()

	for _, file := range files {
		file_name := file.Name()
		var ext string = getFileExtension(file_name)
		folder, status := extensions[ext]
		if !status && !file.IsDir() {
			fmt.Printf("New Extension Found (%s), Enter Folder Name: ", ext)
			fmt.Scan(&folder)
			extensions[ext] = folder
			updateJSON()
		}
		src := filepath.Join(BASE_DIR, file_name)
		if file.IsDir() {
			folder = extensions["folder"]
		}
		dest := filepath.Join(DEST_DIR, folder, file_name)
		os.Rename(src, dest)
		fmt.Printf("- `%s` moved to `%s` Folder...\n", file_name, folder)
	}
	finish := time.Since(start_time)
	fmt.Printf("[INFO] %v - Finished in %v\n", time.Now().Format("01-02-2006 15:04:05"), finish)
}

func getFileFromDownload() []fs.FileInfo {
	files, _ := ioutil.ReadDir(BASE_DIR)
	return files
}

func getFileExtension(FName string) string {
	i := strings.LastIndex(FName, ".") + 1
	return FName[i:]
}

func updateJSON() {
	fmt.Printf("[INFO] %v - Adding new extension\n", time.Now().Format("01-02-2006 15:04:05"))
	b, _ := json.Marshal(extensions)
	ioutil.WriteFile(extPath, b, 0644)
	createDestFolder()
}
