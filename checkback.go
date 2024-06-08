package main

import (
    "flag"
	"fmt"
	"os"
    "time"
)

func main() {
	//file name
    backupPath:= flag.String("path", "", "Check the age of files in this path")
    warnAge:= flag.Int("age", 72, "Warn if older than this many hours")
    flag.Parse()

	//get file info
	fileInfo, err := os.Stat(*backupPath)
	//handle error
	if err != nil {
        if os.IsNotExist(err) {
            fmt.Printf("[ERROR]: No File")
            os.Exit(1)
        } else {
            panic(err)
        }
	}

    modTime := fileInfo.ModTime()

    // Age Check
    currentTime := time.Now()
    fileAge := currentTime.Sub(modTime)
    fileAgeHours := int(fileAge.Hours())

    // Date
    fileDate := modTime.Format("2006-01-02")

    if(fileAgeHours > *warnAge){
        fmt.Printf("[WARN]: %s - %s", *backupPath, fileDate)
        os.Exit(1)
    }

	// fmt.Printf("File Age: %+v\n", fileInfo.ModTime().Unix().Format("2006-01-02"))
	fmt.Printf("[OK]: %s - %s", *backupPath, fileDate)
    os.Exit(0)
}