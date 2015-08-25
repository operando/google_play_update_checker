package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"
	log "github.com/Sirupsen/logrus"
	"github.com/operando/golack"
)

const (
	GOOGLE_PLAY     = "https://play.google.com/store/apps/details?id="
	DEFAULT_PACKAGE = "com.kouzoh.mercari"
)

var old_update_date string
var new_update_date string

func checkUpdate(url string) bool {
	isUpdate := false
	doc, _ := goquery.NewDocument(url)
	doc.Find("div .content").Each(func(_ int, s *goquery.Selection) {
		itemprop, _ := s.Attr("itemprop")
		log.Debug(itemprop)
		if itemprop != "datePublished" {
			return
		}
		log.Debug("Hit!!")

		if old_update_date == "" {
			old_update_date = s.Text()
			log.Info("Old update date : " + old_update_date)
		} else {
			new_update_date = s.Text()
			if old_update_date != new_update_date {
				log.Info("New update date : " + new_update_date)
				isUpdate = true
			}
		}
	})
	log.Debug(isUpdate)
	return isUpdate
}

func main() {
	var sleepTime = flag.Int("t", 1, "sleep time(minute)")
	var packageName = flag.String("p", DEFAULT_PACKAGE, "package name")
	var logLevel = flag.String("l", "info", "log level")
	var configPath = flag.String("c", "", "configuration file path")
	flag.Parse()

	var config Config
	setLogLevel(*logLevel)
	_, err := LoadConfig(*configPath, &config)
	if err != nil {
		fmt.Println(err)
		return
	}
	palyload := golack.Payload{
		config.Slack,
	}

	sleep := time.Duration(*sleepTime*60) * time.Second
	url := GOOGLE_PLAY + *packageName
	log.Info("Check Google Play URL : " + url)

	for {
		if checkUpdate(url) {
			// sh.Command("open", url).Run()
			golack.Post(palyload, config.Webhook)
			log.Info("Update!!!!!!!!!!!")
			break
		} else {
			log.Info("No Update")
		}
		time.Sleep(sleep)
	}

	log.Info("Update check end.")
}

func init() {
	log.SetLevel(log.DebugLevel)
}

func setLogLevel(lv string) {
	switch lv {
	case "debug", "d":
		log.SetLevel(log.DebugLevel)
	case "info", "i":
		log.SetLevel(log.InfoLevel)
	case "warn", "w":
		log.SetLevel(log.WarnLevel)
	case "error", "e":
		log.SetLevel(log.ErrorLevel)
	case "fatal", "f":
		log.SetLevel(log.FatalLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
}
