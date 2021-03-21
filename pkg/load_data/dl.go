package load_data

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func DlRnBase(outFileName string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://renju.net/downloads/getdatabase.php", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("DNT", "1")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "http://renju.net/downloads/games.php")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Sec-GPC", "1")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(outFileName, bodyText, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func DlRoBase(outFileName string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://renjuoffline.com/games.xml.gz", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("DNT", "1")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "https://renjuoffline.com/get_games.php")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Sec-GPC", "1")
	req.Header.Set("TE", "Trailers")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	gzipR, err := gzip.NewReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var bodyBytes bytes.Buffer
	_, err = bodyBytes.ReadFrom(gzipR)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(outFileName, bodyBytes.Bytes(), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
