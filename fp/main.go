package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	path := filepath.Join("dir1", "dir2", "text.txt")

	fmt.Println("This is the path", path)

	//lets check the directory of the application
	fmt.Println("This is the directory of the file", filepath.Dir(path))

	//TIme to check what the base function filepath does ?

	fmt.Println("Base function returened", filepath.Base(path))

	//Time to check if the path is absolute or not (IsAbs)?

	fmt.Println("Our IsAbs returned", filepath.IsAbs(path))

	//Time to check what does the extension (Ext) does

	fmt.Println("Extension function returned", filepath.Ext(path))

	//We can make use oft the Stat functions ffrom the is package to get the metadata infomation about the file
	gfile := "/media/prashant/sda1/kdev/cryptit/fp/tempFile.txt"

	fPath, err := os.Stat("/media/prashant/sda1/kdev/cryptit/fp/tempFile.txt")
	if err != nil {
		fmt.Println("Encounterd Error duing os Stat", err)
	}
	//Checking file size
	fmt.Println("File size is", fPath.Size())

	//checking file name

	fmt.Println("File name : ", fPath.Name())

	//checking file permissions

	fmt.Println("Permissions asscoiated with the file :", fPath.Mode())

	// Checking weather a directory or nor

	fmt.Println("Do we have a directory ?", fPath.IsDir())

	// Reading and Writing file  opeations

	fileData, err := os.ReadFile(gfile)

	if err != nil {
		fmt.Println("Experiencing the following eror", err)
	}

	fmt.Println(string(fileData))

	wfile, err := os.OpenFile(gfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		fmt.Println("Error during opening thr file for write operations", err)
	}

	defer wfile.Close()
	_, err = wfile.WriteString(" We are writing into the file")

	if err != nil {
		fmt.Println("Error when writing")
	} else {
		fmt.Println("Finished writing to the file", wfile.Name())
	}
}
