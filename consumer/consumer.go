package consumer

import (
	"crypto/md5"
	"fmt"
	"github.com/cavaliercoder/grab"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

type Consumer struct {
	downloadFolder string
}

func New(downloadFolder string) *Consumer {
	if downloadFolder == "" {
		downloadFolder = "./download"
	}
	err := os.MkdirAll(downloadFolder, 0777)
	if err != nil {
		panic(err)
	}
	return &Consumer{
		downloadFolder: downloadFolder,
	}
}

// task
func (c *Consumer) Do(uri string) error {
	// TODO: block list
	if uri == "" || uri == "NULL" || uri == "\"\"" {
		return nil
	}

	targetFilename := path.Join(c.downloadFolder, genFilename(uri))

	if fileExists(targetFilename) {
		logrus.Debug("skip downloading as file exists")
		return nil
	}

	resp, err := grab.Get(targetFilename, uri)
	if err != nil {
		logrus.Warnf("failed to download %s, %v", uri, err)
		return err
	}

	logrus.Debugf("download success uri: %s, target: %s", uri, resp.Filename)
	return nil
}

// uri -> unique short name
func genFilename(uri string) string {
	return hash([]byte(uri)) + ".png"
}

// md5
func hash(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	if info == nil {
		return false
	}
	return !info.IsDir()
}
