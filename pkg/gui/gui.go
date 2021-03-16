package gui

import (
	"github.com/webview/webview"
)

func StartUi() {
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Go Analyse")
	w.SetSize(999, 777, webview.HintNone)
	w.Navigate("file:///home/aru/Documents/repos/goanalyse/html.html")
	w.Run()
}
