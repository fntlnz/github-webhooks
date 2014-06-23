package main

import (
	"log"
	"os"
	"fmt"
	"io/ioutil"
)

const lightBlue string = "\033[94m"
const red string = "\033[91m"
const green string = "\033[92m"
const yellow string = "\033[93m"
const endColor string = "\033[0m"

type Logger struct {
	logger *log.Logger
	Coloured bool
}

func NewLogger() *Logger {
	logger := new(Logger)
	logger.logger = log.New(os.Stdout, "[github-webhooks]", 0)
	return logger
}

func (l *Logger) WriteInfo(format string) {
	l.write("[INFO]", lightBlue, format)
}

func (l *Logger) WriteSuccess(format string) {
	l.write("[SUCCESS]", green, format)
}

func (l *Logger) WriteError(format string) {
	l.write("[ERROR]", red, format)
}

func (l *Logger) WriteAlert(format string) {
	l.write("[ALERT]", yellow, format)
}

func (l *Logger) write(prefix string, color string, format string) {
	format = fmt.Sprintf("%s %s", prefix, format)
	if l.Coloured == false {
		l.logger.Printf(fmt.Sprintf("%s", format))
		return
	}
	l.logger.Printf(fmt.Sprintf("%s%s%s", color, format, endColor))
	return
}


func NewNullLogger() *log.Logger {
	return log.New(ioutil.Discard, "", 0)
}
