package main

import (
	"cron-clean/config"
	"os"
	"path/filepath"
	"log"

	"github.com/jasonlvhit/gocron"
)

func main() {

	gocron.Every(1).Day().At("12:20").Do(CronClean)

	<-gocron.Start()
}

func CronClean() {
	log.Println("MASUK")
	err := ClearDir(config.DnsLog)
	if err != nil {
		log.Println(err)
	}
	// var files []string
	// root := config.DnsLog
	// err := filepath.Walk(root, visit(&files))
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	for a, getFile := range files {
	// 		if a != 0 {
	// 			fileName := config.DnsLog + filepath.Base(getFile)
	// 			var extension = filepath.Ext(fileName)
	// 			if extension != ".filepart" {
	// 				file, err := os.Open(fileName) // For read access.
	// 				if err != nil {
	// 					partsFilename := strings.Split(file.Name(), "_")
	// 					partsDate := strings.Split(partsFilename[2], "-")
	// 					//get date
	// 					date := partsDate[2] + "-" + partsDate[1] + "-" + partsDate[0]
	// 					if err != nil {
	// 						log.Println(err)
	// 					}
	// 					scanner := bufio.NewScanner(file)
	// 					scanner.Split(bufio.ScanLines)
	// 					var txtlines []string

	// 					for scanner.Scan() {
	// 						txtlines = append(txtlines, scanner.Text())
	// 					}

	// 					file.Close()
	// 					currentTime := time.Now()
	// 					dateNow := currentTime.Format("2006-01-02")
	// 					if date != dateNow {
	// 						e := os.Remove(fileName)
	// 						if e != nil {
	// 							log.Println(e)
	// 						}
	// 						return
	// 					}
	// 					return
	// 				}
	// 			}
	// 		}
	// 	}
	// }
}

// func visit(files *[]string) filepath.WalkFunc {
// 	return func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		*files = append(*files, path)
// 		return nil
// 	}
// }

func ClearDir(dir string) error {
	log.Println("MASUK2")
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	log.Println(files)
	if err != nil {
		log.Println(err)
		return err
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
