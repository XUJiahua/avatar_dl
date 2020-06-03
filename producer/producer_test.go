package producer

import (
	"testing"
)

func Test_Count(t *testing.T) {
	c, err := Count("sample.csv")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(c)
}

func TestReadLine(t *testing.T) {
	buf, err := ReadLine("sample.csv")
	if err != nil {
		t.Fatal(err)
	}
	for line := range buf {
		t.Log(line)
	}
}
