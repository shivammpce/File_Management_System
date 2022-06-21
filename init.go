package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func init() {
	fmt.Println("** Moving Files/Folders from Downloads **")
	start_time = time.Now()
	readJSON()
	createDestFolder()
}

func readJSON() {
	fmt.Printf("[INFO] %v - Loading Extension types\n", time.Now().Format("01-02-2006 15:04:05"))
	fmt.Println(extPath)
	file, _ := ioutil.ReadFile(extPath)
	json.Unmarshal(file, &extensions)
}

func createDestFolder() {
	fmt.Printf("[INFO] %v - Checking/Updating for Destination folders\n", time.Now().Format("01-02-2006 15:04:05"))
	for _, name := range extensions {
		path := filepath.Join(DEST_DIR, name)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Println("- Creating", name, "folder")
			os.Mkdir(path, 0755)
		}
	}
}
