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
	tbget.TOR_MIRROR = *flag.String("mirror", "https://dist.torproject.org/torbrowser/", "Tor mirror to use")
	tbget.TOR_DOWNLOADER_VERBOSE = *flag.Bool("verbose", false, "Be verbose")
	tbget.TOR_BROWSER_OS = *flag.String("os", tbget.OS(), "OS to get Tor Browser for")
	tbget.TOR_BROWSER_ARCH = *flag.String("arch", tbget.ARCH(), "Tor Browser architecture to download")
	flag.Parse()
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
