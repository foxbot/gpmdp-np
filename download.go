// Credits to https://github.com/thbar/golang-playground/blob/master/download-files.go

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func downloadFromUrl(url string, outputFile string) {
	fmt.Println("Downloading", url, "to", outputFile)

	// TODO: check file existence first with io.IsExist
	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error while creating", outputFile, "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}

	fmt.Println(n, "bytes downloaded.")
}