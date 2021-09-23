package uriget

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"

	"github.com/HasinduLanka/gocommons/console"
)

var regex_url *regexp.Regexp = regexp.MustCompile("http.*")

var RetryCount int = 3
var Download_cache_dir string = "./cache/"

func LoadURI(uri string) ([]byte, error) {
	if regex_url.MatchString(uri) {
		return DownloadFileToBytes(uri)
	} else {
		return LoadFile(uri)
	}

}

func LoadURICached(uri string) ([]byte, error) {
	if regex_url.MatchString(uri) {

		filename := CachedName(uri)
		CB, CE := LoadFile(uri)
		if CE == nil {
			return CB, nil
		}

		er := DownloadToFile(filename, uri)
		if er != nil {
			return nil, er
		}

		return LoadFile(filename)

	} else {
		return LoadFile(uri)
	}
}

func CachedName(uri string) string {
	filename := Download_cache_dir + ".cache-dwn-" + url.PathEscape(uri)
	return filename
}

func InvalidateCacheURI(uri string) {
	filename := CachedName(uri)
	DeleteFiles(filename)
}

func AppendFile(filename string, content []byte) bool {
	F, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if console.PrintError(err) {
		return false
	}
	F.Write(content)
	F.Close()
	return true
}

func WriteFile(filename string, content []byte) bool {
	F, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if console.PrintError(err) {
		return false
	}
	F.Write(content)
	F.Close()

	return true
}

func LoadFileToIOReader(filename string) (io.ReadCloser, error) {
	file, err := os.Open(filename)
	if console.PrintError(err) {
		return nil, err
	}
	return file, nil
}

func MakeDir(name string) {
	os.MkdirAll(name, os.ModePerm)
}

func DeleteFiles(name string) {
	os.RemoveAll(name)
}

func LoadURIToString(uri string) (string, error) {
	B, err := LoadURI(uri)
	return string(B), err
}

func LoadURIToStringCached(uri string) (string, error) {
	B, err := LoadURICached(uri)
	return string(B), err
}

func LoadFile(filename string) ([]byte, error) {
	file, ferr := ioutil.ReadFile(filename)
	return file, ferr
}

func LoadFileToString(filename string) (string, error) {
	F, err := LoadFile(filename)
	return string(F), err
}

func DownloadToFile(filepath string, url string) error {

	body, derr := DownloadFileToStream(url)

	if derr != nil {
		return derr
	}
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, body)
	console.PrintError(body.Close())
	return err
}

func DownloadFileToStream(url string) (io.ReadCloser, error) {
	return DownloadFileToStreamRetry(url, RetryCount)
}

func DownloadFileToStreamRetry(url string, retry int) (io.ReadCloser, error) {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		if retry > 0 {
			return DownloadFileToStreamRetry(url, retry-1)
		}
		return nil, err
	}

	return resp.Body, nil
}

func DownloadFileToBytes(url string) ([]byte, error) {
	console.Print("GET " + url)
	str, err := DownloadFileToStream(url)

	if err != nil {
		return nil, err
	}
	out := StreamToByte(str)
	closeerr := str.Close()
	if console.PrintError(closeerr) {
		return nil, closeerr
	}
	return out, nil
}

func DownloadFileToString(url string) (string, error) {
	str, err := DownloadFileToStream(url)

	if err != nil {
		return "", err
	}
	out := StreamToString(str)
	closeerr := str.Close()
	if console.PrintError(closeerr) {
		return "", closeerr
	}
	return out, nil
}

func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}

func StreamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.String()
}
