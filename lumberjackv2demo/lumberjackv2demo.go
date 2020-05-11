package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/natefinch/lumberjack"
)

type stdLogger struct {
	*lumberjack.Logger
}

var logger stdLogger

type LogConfig struct {
	Logfile string
	MaxSize int `toml:"max_log_size"`
	MaxAge  int `toml:"max_log_age"`
}

func (c *LogConfig) SetLogger() {
	if c == nil || c.Logfile == "" {
		return
	}
	fmt.Printf("Sending log messages to: %s\n", c.Logfile)
	l := &lumberjack.Logger{
		Filename:   c.Logfile,
		MaxSize:    c.MaxSize, // megabytes
		MaxAge:     c.MaxAge,  //days
		MaxBackups: 3,         // files
		Compress:   true,      // disabled by default
	}
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	log.SetOutput(l)
	logger = stdLogger{l}
}

func (slog stdLogger) Shutdown() {
	// log.Printf("Closing log file...\n")
	if slog.Logger != nil {
		slog.Close()
	}
}

func main() {

	fileName := "app/kafka_error_logs/log_by_under_eiei.txt"
	configLog := LogConfig{
		Logfile: fileName,
		MaxSize: 5,
		MaxAge:  3,
	}

	configLog.SetLogger()

	txStr := "test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file end"
	for i := 1; i <= 100000; i++ {
		log.Println(txStr)
	}
	logger.Shutdown()

	fmt.Println(configLog)
	fileName = "app/kafka_error_logs/log_by_under_2.txt"
	anotherStr := "lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum lorem ipzum end"
	fmt.Println(configLog)
	fmt.Println(logger.Filename)
	logger.Filename = fileName
	fmt.Println(logger.Filename)
	for i := 1; i <= 100000; i++ {
		log.Println(anotherStr)
	}
	logger.Shutdown()
}

func GetFileNameFormat(podID, fileName string) string {
	now := time.Now()

	fileFormat := now.Format("20060201") + "_" + podID
	if strings.Contains(fileName, "reverse") == true {
		fileFormat = strings.ToLower(now.Format("Mon")) + "_" + fileFormat
	}

	cleanFileName := fmt.Sprintf(fileName, fileFormat)
	return cleanFileName
}

func WriteLogFile(fileName string, txStr string) {
	l := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    50,   // megabytes
		MaxBackups: 3,    // files
		MaxAge:     28,   // days
		Compress:   true, // disabled by default
	}

	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	log.SetOutput(l)
	log.Println(txStr)
}

type LogKafkaMessageToFileFunc func(string, string, string)

func LogKafkaMessageToFile(fileName string, txStr string, podID string) {
	cleanFileName := GetFileNameFormat(podID, fileName)
	WriteLogFile(cleanFileName, txStr)
}
