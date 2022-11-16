package main

import (
	"flag"
	"log"
	"os"

	tbget "i2pgit.org/idk/tb-downloader/helper"
)

func wd() string {
	w, err := os.Getwd()
	if err != nil {
		return "."
	}
	return w
}

var (
	launch    = flag.Bool("launch", false, "Launch Tor Browser after downloading")
	directory = flag.String("directory", wd(), "Directory to download Tor Browser to")
)

func main() {
	cmd, err := tbget.DownloadVerifyUnpackTorBrowser(*directory)
	if err != nil {
		log.Fatalln(err)
	}
	if *launch {
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(out)
	}
}
