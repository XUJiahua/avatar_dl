package util

import (
	"github.com/sirupsen/logrus"
	"time"
)

func Elapsed() func() {
	start := time.Now()
	return func() {
		logrus.Infof("%s took \n", time.Since(start))
	}
}
