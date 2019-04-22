package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func SavePic(URL string, filename string) error {
	response, err := http.Get(URL)
	if err != nil {
		log.Println(err)
		return err
	}

	defer response.Body.Close()

	file, err := os.Create(fmt.Sprintf("%s/src/github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi/public/%s", os.Getenv("GOPATH"), filename))
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Println(err)
		return err
	}

	file.Close()

	return nil
}

func GetPic(filename string) string {
	return fmt.Sprintf(fmt.Sprintf("%s/public/%s", os.Getenv("APP_HOST"), filename))
}
