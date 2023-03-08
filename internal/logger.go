package internal

import (
	"log"
	"os"
)

func NewDefaultLogger() *log.Logger {
	flags := log.LstdFlags | log.Lshortfile
	return log.New(os.Stdout, "", flags)
}
