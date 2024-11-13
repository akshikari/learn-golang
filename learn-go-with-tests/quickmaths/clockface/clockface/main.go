package main

import (
	"os"
	"time"

	"github.com/akshikari/learn-golang/learning-go/learn-go-with-tests/quickmaths/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
