package crawler

import (
	"net/url"
	"path"
	"regexp"
	"strings"
)

var ENDS_WITH_FILENAME_REGEX = regexp.MustCompile("^.*\\..+$")

type CUrl struct {
	url   url.URL
	depth int
}

type Crawler struct {
	url          string
	targetFolder string
	maxDepth     int
}

func (crawler *Crawler) Crawl() {

}

/*
   Download the given url to rootDir folder
*/
func saveTo(url url.URL, body []byte, rootDir string) (string, error) {
	//file := rootDir + url.Host + url.Path

	//err = ioutil.WriteFile(fi.Name(), []byte(Value), 0644)
	return "", nil
}

func constructFilePath(url *url.URL, rootDir string) string {
	var filePath, fileName string

	if ENDS_WITH_FILENAME_REGEX.MatchString(url.Path) {
		filePath, fileName = path.Split(url.Path)
	} else {
		filePath = url.Path
		fileName = "index.html"
	}

	var file string
	if filePath == "" || filePath == "/" {
		file = rootDir + "/" + url.Host + "/" + fileName
	} else {
		if strings.HasSuffix(filePath, "/") {
			file = rootDir + "/" + url.Host + filePath + fileName
		} else {
			file = rootDir + "/" + url.Host + filePath + "/" + fileName
		}
	}

	return file
}

func main() {
	crawler := &Crawler{"http://huffpost.com", "/tmp/crawler", 5}
	crawler.Crawl()

	url, _ := url.Parse("http://www.google.com")
	println(constructFilePath(url, "/tmp"))
}
