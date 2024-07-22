package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var (
	windRegex     = regexp.MustCompile(`\d* METAR.*EGLL \d*Z [A-Z ]*(\d{5}KT|VRB\d{2}KT).*=`)
	tafValidation = regexp.MustCompile(`.*TAF.*`)
	comment       = regexp.MustCompile(`\w*#.*`)
	metarClose    = regexp.MustCompile(`.*=`)
	variableWind  = regexp.MustCompile(`.*VRB\d{2}KT`)
	validWind     = regexp.MustCompile(`\d{5}KT`)
	windDirOnly   = regexp.MustCompile(`(\d{3})\d{2}KT`)
)

func parseToArray(textChannel chan string, metaChannel chan []string) {
	for text := range textChannel {
		lines := strings.Split(text, "\n")
		metarSlice := make([]string, 0, len(lines))
		metarStr := ""
		for _, line := range lines {
			if tafValidation.MatchString(line) {
				break
			}
			if !comment.MatchString(line) {
				metarStr += strings.Trim(line, " ")
			}
			if metarClose.MatchString(line) {
				metarSlice = append(metarSlice, metarStr)
				metarStr = ""
			}
		}
		metaChannel <- metarSlice
	}
	close(metaChannel)
}

func extractWindDirection(metarChannel chan []string, windsChannel chan []string) {
	for metars := range metarChannel {
		winds := make([]string, 0, len(metars))
		for _, metar := range metars {
			winds = append(winds, metar)
		}
		windsChannel <- winds
	}
	close(windsChannel)
}

func main() {
	textChannel := make(chan string)
	metaChannel := make(chan []string)
	windsChannel := make(chan []string)
	// resultChannel := make(chan [8]int)

	go parseToArray(textChannel, metaChannel)

	go extractWindDirection(metaChannel, windsChannel)

	absPath, err := filepath.Abs("../metarfiles")
	if err != nil {
		panic(err)
	}

	files, err := os.ReadDir(absPath)
	if err != nil {
		panic(err)
	}
	start := time.Now()
	for _, file := range files {
		dat, err := os.ReadFile(filepath.Join(absPath, file.Name()))

		if err != nil {
			panic(err)
		}
		text := string(dat)
		textChannel <- text
	}
	close(textChannel)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	fmt.Println(<-windsChannel)
}
