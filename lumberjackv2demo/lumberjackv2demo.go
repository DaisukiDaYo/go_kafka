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

func main() {
	configLog := LogConfig{
		Logfile: "default.txt",
		MaxSize: 5,
		MaxAge:  3,
	}
	configLog.SetLogger()

	fileName := "app/kafka_error_logs/test_rotate_%s.txt"
	txStr := "test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file end"
	podID := "1243"
	fmt.Println(logger)
	for i := 1; i <= 105000; i++ {
		LogKafkaMessageToFile(fileName, txStr, podID)
	}

	fileName = "app/kafka_error_logs/test_rotate_2_%s.txt"
	txStr = "2test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file end"
	podID = "1245"
	for i := 1; i <= 105000; i++ {
		LogKafkaMessageToFile(fileName, txStr, podID)
	}
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
	logger.Filename = fileName
	log.Println(txStr)
	defer logger.Close()
}

type LogKafkaMessageToFileFunc func(string, string, string)

func LogKafkaMessageToFile(fileName string, txStr string, podID string) {
	cleanFileName := GetFileNameFormat(podID, fileName)
	fmt.Println(cleanFileName)
	WriteLogFile(cleanFileName, txStr)
}
