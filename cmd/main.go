package main

// importing our social media packages from exporter
import (
	"exporter"
	"exporter/facebook"
	"exporter/linkedin"
	"exporter/twitter"
	"fmt"
)

func main() {

	fb := new(facebook.Facebook)
	twtr := new(twitter.Twitter)
	lnkd := new(linkedin.Linkedin)

	err := exporter.Txt(fb, "fbdata.txt")
	err = exporter.Txt(twtr, "twtrdata.txt")
	err = exporter.Txt(lnkd, "lnkddata.txt")

	// checking for errors then panic if any
	if err != nil {
		panic(err)
	}

	x := exporter.XML(fb, "fbdata.xml")
	x = exporter.XML(twtr, "twtrdata.xml")
	x = exporter.XML(lnkd, "lnkddata.xml")

	// checking for errors then panic if any
	if x != nil {
		panic(err)
	}

	jsn := exporter.JSON(fb, "fbdata.json")
	jsn = exporter.JSON(twtr, "twtrdata.json")
	jsn = exporter.JSON(lnkd, "lnkddata.json")

	// checking for errors then panic if any
	if jsn != nil {
		panic(err)
	}
	ScrollFeeds(fb, twtr, lnkd)
}

//ScrollFeeds prints all SocialMedia
func ScrollFeeds(platforms ...exporter.SocialMedia) {
	for _, sm := range platforms {
		for _, fd := range sm.Feed() {
			fmt.Println(fd)
		}
		fmt.Println("=================================")
	}
}
