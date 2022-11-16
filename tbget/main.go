package main

import (
	"log"

	tbget "i2pgit.org/idk/tb-downloader/helper"
)

func main() {
	cmd, err := tbget.DownloadVerifyUnpackTorBrowser(".")
	if err != nil {
		log.Fatalln(err)
	}
	out, err := cmd.CombinedOutput()
	log.Println(out)
}
