package helpers

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
	"os"
	"strings"
)

func LoadFile(pathFile string) *excelize.File {
	// clear tex pathFile
	pathFile = strings.TrimRight(pathFile, "\n")
	// Load file
	xlsx, err := excelize.OpenFile(pathFile)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	return xlsx
}
