package member

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi/utils"
)

type TelegramPhoto struct {
	FilePath string `json:"file_path"`
}

type TelegramResponse struct {
	RootPath string            `json:"rootPath"`
	Photos   [][]TelegramPhoto `json:"data"`
}

func RetrieveTelegramPicture(ID string) (string, error) {
	var tresponse TelegramResponse

	botHost := os.Getenv("BOT_HOST")

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/pic/%s", botHost, ID), nil)
	if err != nil {
		log.Println(err)
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}

	err = utils.DecodeResponse(resp, &tresponse)
	if err != nil {
		log.Println(err)
		return "", err
	}

	if len(tresponse.Photos) == 0 {
		return "", errors.New("No photos")
	}
	return fmt.Sprintf("%s%s", tresponse.RootPath, tresponse.Photos[0][0].FilePath), nil
}
