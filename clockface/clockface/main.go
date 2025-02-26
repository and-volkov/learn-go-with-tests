package main

import (
	"os"
	"time"

	"study.go/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
