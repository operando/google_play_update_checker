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
	GOOGLE_PLAY = "https://play.google.com/store/apps/details?id="
)

var old_update_date string
var new_update_date string

func checkUpdate(url string) bool {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
		return false
	}
	isUpdate := false
	doc.Find("div[itemprop=\"datePublished\"]").Each(func(_ int, s *goquery.Selection) {
		log.Debug(s.Text())
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
	var configPath = flag.String("c", "", "configuration file path")
	flag.Parse()

	var config Config
	_, err := LoadConfig(*configPath, &config)
	if err != nil {
		fmt.Println(err)
		return
	}
	palyload := golack.Payload{
		config.Slack,
	}

	setLogLevel(config.Log)
	sleep := time.Duration(config.SleepTime*60) * time.Second
	url := GOOGLE_PLAY + config.Package
	log.Info("Check Google Play URL : " + url)
	log.Info("Slack Post Message : " + config.Slack.Text)

	for {
		if checkUpdate(url) {
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
	log.SetLevel(log.InfoLevel)
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
