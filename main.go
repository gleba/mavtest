package main

import (
	webview "github.com/webview/webview_go"
)

func maizn() {

	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Basic Example")
	w.SetSize(480, 320, webview.HintNone)
	html := `
<meta http-equiv="refresh" content="0; url=http://localhost:3000">
	`
	w.SetHtml(html)

	//w.Navigate("http://localhost:3000")
	w.Run()
	//w.Navigate("http://localhost:3000")
	//for true {
	//
	//}
}
