package main

import (
	"bufio"
	"cron-clean/config"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jasonlvhit/gocron"
)

func main() {

	gocron.Every(10).Seconds().Do(CronClean)

	<-gocron.Start()
}

func CronClean() {
	var files []string
	root := config.DnsLog
	err := filepath.Walk(root, visit(&files))
	if err != nil {
		log.Println(err)
	} else {
		for a, getFile := range files {
			if a != 0 {
				fileName := config.DnsLog + filepath.Base(getFile)
				var extension = filepath.Ext(fileName)
				log.Println("MASUK 1");
				if extension != ".filepart" {
				log.Println("MASUK 2");
					file, err := os.Open(fileName) // For read access.
					partsFilename := strings.Split(file.Name(), "_")
					partsDate := strings.Split(partsFilename[2], "-")
					log.Println("MASUK 3");
					//get date
					date := partsDate[2] + "-" + partsDate[1] + "-" + partsDate[0]
					if err != nil {
						log.Println(err)
					}
					log.Println("MASUK 4");
					scanner := bufio.NewScanner(file)
					scanner.Split(bufio.ScanLines)
					var txtlines []string

					for scanner.Scan() {
						txtlines = append(txtlines, scanner.Text())
					}
					log.Println("MASUK 5");

					file.Close()
					currentTime := time.Now()
					dateNow := currentTime.Format("2006-01-02")
					if date != dateNow {
					log.Println("MASUK 6");
						e := os.Remove(fileName)
						if e != nil {
							log.Println(e)
						}
						return
					}
				}
			}
		}
	}
}

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
		}
		*files = append(*files, path)
		return nil
	}
}
