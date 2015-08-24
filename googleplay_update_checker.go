package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/codeskyblue/go-sh"
)

const (
	GOOGLE_PLAY     = "https://play.google.com/store/apps/details?id="
	DEFAULT_PACKAGE = "com.kouzoh.mercari"
	// GOOGLE_PLAY     = "https://dl.dropboxusercontent.com/u/97368150/test.html"
	// DEFAULT_PACKAGE = ""
)

var old_update_date string
var new_update_date string

func checkUpdate(url string) bool {
	isUpdate := false
	doc, _ := goquery.NewDocument(url)
	doc.Find("div .content").Each(func(_ int, s *goquery.Selection) {
		itemprop, _ := s.Attr("itemprop")
		fmt.Println(itemprop)
		if itemprop != "datePublished" {
			return
		}
		fmt.Println("Hit!!")

		if old_update_date == "" {
			old_update_date = s.Text()
			fmt.Println("Old update date : " + old_update_date)
		} else {
			new_update_date = s.Text()
			if old_update_date != new_update_date {
				fmt.Println("New update date : " + new_update_date)
				isUpdate = true
			}
		}
	})
	fmt.Println(isUpdate)
	return isUpdate
}

func main() {
	var sleepTime = flag.Int("t", 1, "sleep time(minute)")
	var packageName = flag.String("p", DEFAULT_PACKAGE, "package name")
	flag.Parse()

	url := GOOGLE_PLAY + *packageName
	fmt.Println("Check Google Play URL : " + url)

	for {
		if checkUpdate(url) {
			sh.Command("open", url).Run()
			fmt.Println("Update!!!!!!!!!!!")
			break
		} else {
			fmt.Println("No Update")
		}
		time.Sleep(time.Duration(*sleepTime*60) * time.Second)
	}

	fmt.Println("Update check end.")
}
