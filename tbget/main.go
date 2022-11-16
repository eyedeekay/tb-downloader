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
	launch                 = flag.Bool("launch", false, "Launch Tor Browser after downloading")
	directory              = flag.String("directory", wd(), "Directory to download Tor Browser to")
	TOR_MIRROR             = flag.String("mirror", "https://dist.torproject.org/torbrowser/", "Tor mirror to use")
	TOR_DOWNLOADER_VERBOSE = flag.Bool("verbose", false, "Be verbose")
	TOR_BROWSER_OS         = flag.String("os", tbget.OS(), "OS to get Tor Browser for")
	TOR_BROWSER_ARCH       = flag.String("arch", tbget.ARCH(), "Tor Browser architecture to download")
)

func main() {
	flag.Parse()
	tbget.TOR_MIRROR = *TOR_MIRROR
	tbget.TOR_DOWNLOADER_VERBOSE = *TOR_DOWNLOADER_VERBOSE
	tbget.TOR_BROWSER_OS = *TOR_BROWSER_OS
	tbget.TOR_BROWSER_ARCH = *TOR_BROWSER_ARCH
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
