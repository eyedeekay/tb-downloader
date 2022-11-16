package tbget

import "testing"

func TestXxx(t *testing.T) {
	_, err := DownloadVerifyUnpackTorBrowser(".")
	if err != nil {
		t.Fatal(err)
	}
}
