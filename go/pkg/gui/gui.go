package gui

import (
	"github.com/webview/webview"
	"io/ioutil"
	"os"
	"path/filepath"
)

func StartUi() {
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Go Analyse")
	w.SetSize(999, 777, webview.HintNone)
	selfbinPath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	selfbinDir := filepath.Dir(selfbinPath)
	html, err := ioutil.ReadFile(selfbinDir + "/html.html")
	if err != nil {
		panic(err)
	}
	w.Navigate(string(html))
	w.Run()
}
