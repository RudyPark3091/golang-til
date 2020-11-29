package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func DownloadImage(url, filename, targetPath string) error {
	fmt.Println("Downloading image from :", url)
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		return errors.New("Http Get Failed")
	}

	file, err := os.Create(targetPath + "/" + filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, r.Body)
	if err != nil {
		return err
	}

	fmt.Println("\x1b[31mDownloaded image in", path.Join(targetPath, "/", filename), "\x1b[0m")

	return nil
}
