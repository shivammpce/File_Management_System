package main

import (
	"os"
	"path/filepath"
	"time"
)

var HOME, _ = os.UserHomeDir()

var BASE_DIR = filepath.Join(HOME, "Downloads")
var DEST_DIR = filepath.Join(HOME, "Documents")

var extPath = filepath.Join(HOME, ".extensions.json")
var extensions = map[string]string{}

var start_time time.Time
