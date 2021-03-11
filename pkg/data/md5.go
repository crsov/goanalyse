package data

import (
	"crypto/md5"
	"io"
	"log"
	"os"
)

func FileMD5(path string) string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return string(h.Sum(nil))
}
