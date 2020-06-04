package util

import (
	"fmt"
	"time"
)

func Elapsed() func() {
	start := time.Now()
	return func() {
		fmt.Printf("time elapsed: %s \n", time.Since(start))
	}
}
