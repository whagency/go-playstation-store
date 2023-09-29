package ps

import (
	"io"
	"log"
	"os"
)

const (
	logPath = "log"
	logFile = "log/ps.log"
)

const (
	LoggerFile   = 1
	LoggerStdout = 2
	LoggerOff    = 0
)

type logger struct {
	Enable bool
	Info   *log.Logger
	Error  *log.Logger
}

func newPsLogger(output uint8) logger {
	var logger logger
	var writer io.Writer
	if output == LoggerFile {
		if err := os.MkdirAll(logPath, os.ModePerm); err != nil {
			panic(err)
		}
		var file, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		writer = file
	} else if output == LoggerStdout {
		writer = os.Stdout
	} else {
		return logger
	}

	logger.Enable = true
	logger.Info = log.New(writer, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Error = log.New(writer, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)

	return logger
}
