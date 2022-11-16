package main

import (
	"log"

	tbget "i2pgit.org/idk/tb-downloader/helper"
)

func main() {
	_, err := tbget.DownloadVerifyUnpackTorBrowser(".")
	if err != nil {
		log.Fatalln(err)
	}
}
