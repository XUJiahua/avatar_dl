package producer

import (
	"avatar_dl/consumer"
	"bufio"
	"github.com/cheggaaa/pb/v3"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

type Producer struct {
	filename        string
	failureFilename string
}

func New(filename string, failureFilename string) *Producer {
	return &Producer{
		filename:        filename,
		failureFilename: failureFilename,
	}
}

func (p Producer) Do(consumer *consumer.Consumer, workers int) {
	count, err := Count(p.filename)
	if err != nil {
		logrus.Fatal(err)
		return
	}

	// for producing lines from files
	buf, err := ReadLine(p.filename)
	if err != nil {
		logrus.Fatal(err)
		return
	}

	// progress bar in console
	bar := pb.StartNew(count)
	defer bar.Finish()

	// workers and 1 special worker for handling failed task
	workersCh := make(chan int, workers+1)

	// special worker for
	failedQueue := make(chan string)
	workersCh <- 1
	go func() {
		file, err := os.OpenFile(p.failureFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			logrus.Error(err)
			os.Exit(2)
		}

		writer := bufio.NewWriter(file)
		for uri := range failedQueue {
			_, _ = writer.WriteString(uri + "\n")
		}
		_ = writer.Flush()
		_ = file.Close()

		<-workersCh
	}()

	// handling lines
	for line := range buf {
		workersCh <- 1
		// worker goroutine
		go func(uri string, ch chan int) {
			if err := consumer.Do(uri); err != nil {
				failedQueue <- uri
			}

			// mark as complete
			bar.Increment()
			<-ch
		}(strings.TrimSpace(line), workersCh)
	}

	// wait for workers
	for i := 0; i < workers; i++ {
		workersCh <- 1
	}

	// wait for special worker
	close(failedQueue)
	workersCh <- 1
}

// line count of file
func Count(filename string) (int, error) {
	buf, err := ReadLine(filename)
	if err != nil {
		return 0, errors.Wrap(err, "failed to read line")
	}

	c := 0
	for range buf {
		c++
	}
	return c, nil
}

// read line of file
func ReadLine(filename string) (chan string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open file")
	}

	buf := make(chan string)

	go func(file *os.File) {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			buf <- scanner.Text()
		}
		if err = scanner.Err(); err != nil {
			logrus.Error(err)
		}

		close(buf)
	}(file)

	return buf, nil
}
