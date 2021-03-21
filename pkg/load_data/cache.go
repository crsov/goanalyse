package load_data

import (
	"bytes"
	"encoding/gob"
	"github.com/djherbis/atime"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func filenameWithoutExtension(file string) string {
	return strings.TrimSuffix(file, path.Ext(file))
}

func LoadAccordingToCache(inputFile string) ro {
	fileHash, localDir := getFileAdditions(inputFile)

	// If there no gob file we get an error but we dont worry about that, and just gets nil
	gobFile, _ := os.Open(filepath.Join(localDir, "cache", filenameWithoutExtension(inputFile)+".gob"))
	defer gobFile.Close()
	// removeOldCache checks atime and we need to check it after openning, but and after closing gobFile, and there only one way to do it - defer it after gobFile.close
	defer removeOldCache(filepath.Join(localDir, "cache"))

	if gobFile != nil {
		// If there no sum file we get an error but we dont worry about that, and just gets nil
		sumFile, _ := ioutil.ReadFile(filepath.Join(localDir, "cache", filenameWithoutExtension(inputFile)+".sum"))

		if string(sumFile) != fileHash {
			writeInCache(inputFile)

		}
		return useGob(gobFile)
	} else {
		return writeInCache(inputFile)

	}
}

func writeInCache(inputFile string) ro {
	fileHash, localDir := getFileAdditions(inputFile)
	err := os.WriteFile(filepath.Join(localDir, "cache", filenameWithoutExtension(inputFile)+".sum"), []byte(fileHash), 0666)
	if err != nil {
		panic(err)
	}

	games := ParseXml(inputFile)
	var buff bytes.Buffer
	encoder := gob.NewEncoder(&buff)
	err = encoder.Encode(games)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filepath.Join(localDir, "cache", filenameWithoutExtension(inputFile)+".gob"), buff.Bytes(), 0666)
	if err != nil {
		panic(err)
	}
	return games
}

func getFileAdditions(inputFile string) (string, string) {
	fileHash := FileMD5(inputFile)
	executablePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	localDir := filepath.Dir(executablePath)
	return fileHash, localDir
}

func useGob(gobFile io.Reader) ro {

	decoder := gob.NewDecoder(gobFile)
	var games ro
	err := decoder.Decode(&games)
	if err != nil {
		panic(err)
	}
	return games
}

func removeOldCache(pathToCache string) {
	cacheInfo, err := ioutil.ReadDir(pathToCache)
	if err != nil {
		panic(err)
	}

	for _, cacheFile := range cacheInfo {
		cacheFilePath := filepath.Join(pathToCache, cacheFile.Name())
		atime, err := atime.Stat(cacheFilePath)
		if err != nil {
			panic(err)
		}
		// 7776000 = 3 months
		if atime.Unix() < time.Now().Unix()-7776000 {
			err := os.Remove(cacheFilePath)
			if err != nil {
				panic(err)
			}
		}
	}
}
