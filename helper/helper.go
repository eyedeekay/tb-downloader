package tbget

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	tbget "i2pgit.org/idk/tb-downloader"
)

func getenv(v, d string) string {
	m := os.Getenv("TOR_MIRROR")
	if m == "" {
		return d
	}
	return m
}

func getenvbool(v, d string) bool {
	m := getenv(v, d)
	switch m {
	case "true":
		return true
	case "TRUE":
		return true
	case "True":
		return true
	case "t":
		return true
	case "T":
		return true
	case "YES":
		return true
	case "Yes":
		return true
	case "yes":
		return true
	case "y":
		return true
	default:
		return false
	}
}

var TOR_MIRROR = getenv("TOR_MIRROR", "https://dist.torproject.org/torbrowser/")
var TOR_DOWNLOADER_VERBOSE = getenvbool("TOR_DOWNLOADER_VERBOSE", "false")
var TOR_NO_UNPACK = getenvbool("TOR_NO_UNPACK", "false")
var TOR_BROWSER_OS = getenv("TOR_BROWSER_OS", tbget.OS())
var TOR_BROWSER_ARCH = getenv("TOR_BROWSER_ARCH", tbget.ARCH())

func DownloadVerifyUnpackTorBrowser(Directory string) (*exec.Cmd, error) {
	lang := tbget.DefaultIETFLang
	os := tbget.OS()
	arch := tbget.ARCH()
	tbdownloader := tbget.NewTBDownloader(lang, os, arch, nil)
	tbdownloader.DownloadPath = Directory
	tbdownloader.Mirror = TOR_MIRROR
	tbdownloader.Verbose = TOR_DOWNLOADER_VERBOSE
	tbdownloader.NoUnpack = TOR_NO_UNPACK
	tbdownloader.Profile = &tbget.Content
	tbdownloader.MakeTBDirectory()
	tgz, sig, _, err := tbdownloader.DownloadUpdaterForLang(lang)
	if err != nil {
		return nil, err
	}
	//var home string
	if _, err := tbdownloader.CheckSignature(tgz, sig); err != nil {
		log.Fatal(err)
	} else {
		out, err := tbdownloader.UnpackUpdater(tgz)
		if err != nil {
			return nil, fmt.Errorf("unpacking updater: %v", err)
		}
		log.Printf("Signature check passed: %s %s", tgz, sig)
		bin := tbget.TBPath(out)
		log.Printf("output is: %s", bin)
		cmd := exec.Command(bin)
		return cmd, err
	}
	return nil, err
}

func OS() string {
	return tbget.OS()
}

func ARCH() string {
	return tbget.ARCH()
}
