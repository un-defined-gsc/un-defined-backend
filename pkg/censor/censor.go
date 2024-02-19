package censor_service

import (
	"bufio"
	"os"

	censor "github.com/pcpratheesh/go-censorword"
)

type CensorService struct {
	detector *censor.CensorWordDetection
}

func NewCensorService(file_path string) *CensorService {

	badword, err := os.OpenFile(file_path, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer badword.Close()

	badwords := make([]string, 0)
	scanner := bufio.NewScanner(badword)
	for scanner.Scan() {
		badwords = append(badwords, scanner.Text())
	}

	dec := censor.NewDetector(
		censor.WithCensorReplaceChar(""),
		censor.WithCustomCensorList(badwords),
	)
	return &CensorService{
		detector: dec,
	}
}

func (c *CensorService) CensorText(textAddrs ...*string) (err error) {

	if len(textAddrs) > 0 {
		for idx := range textAddrs {
			*textAddrs[idx], err = c.detector.CensorWord(*textAddrs[idx])
			if err != nil {
				return
			}
		}
	}
	return
}
