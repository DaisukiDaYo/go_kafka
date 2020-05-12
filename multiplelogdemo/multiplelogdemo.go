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

var (
	ReverseTodayLogger   stdLogger
	ForcepostTodayLogger stdLogger
)

type LogConfig struct {
	Logfile string
	MaxSize int
	MaxAge  int
}

func (c *LogConfig) SetLogger() (*log.Logger, *log.Logger) {
	fmt.Printf("Sending log messages to: %s\n", c.Logfile)
	reverseLumberjack := &lumberjack.Logger{
		Filename:   c.Logfile,
		MaxSize:    c.MaxSize, // megabytes
		MaxAge:     c.MaxAge,  //days
		MaxBackups: 3,         // files
		Compress:   true,      // disabled by default
	}
	reverseLogger := log.New(reverseLumberjack, "", 1)
	reverseLogger.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	ReverseTodayLogger = stdLogger{reverseLumberjack}

	forcepostLumberjack := &lumberjack.Logger{
		Filename:   c.Logfile,
		MaxSize:    c.MaxSize, // megabytes
		MaxAge:     c.MaxAge,  //days
		MaxBackups: 3,         // files
		Compress:   true,      // disabled by default
	}

	forcepostLogger := log.New(forcepostLumberjack, "", 1)
	forcepostLogger.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	ForcepostTodayLogger = stdLogger{forcepostLumberjack}

	return reverseLogger, forcepostLogger
}

func main() {
	configLog := LogConfig{
		Logfile: "default.txt",
		MaxSize: 5,
		MaxAge:  3,
	}
	reverseLogger, forcepostLogger := configLog.SetLogger()

	fileName := "app/kafka_error_logs/reverse_rotate_%s.txt"
	txStr := "test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file end"
	podID := "1243"
	for i := 1; i <= 105000; i++ {
		LogKafkaMessageToFile(fileName, txStr, podID, ReverseTodayLogger, reverseLogger)
	}

	fileName = "app/kafka_error_logs/force_post_rotate_2_%s.txt"
	txStr = "2test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file end"
	podID = "1245"
	for i := 1; i <= 105000; i++ {
		LogKafkaMessageToFile(fileName, txStr, podID, ForcepostTodayLogger, forcepostLogger)
	}

	t := time.Now()
	rounded := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	fmt.Println(t)
	fmt.Println(rounded)
	if true {
		ReverseTodayLogger.Close()
	}

	fileName = "app/kafka_error_logs/reverse_rotate_tmr_%s.txt"
	txStr = "3test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file end"
	podID = "1243"
	for i := 1; i <= 105000; i++ {
		LogKafkaMessageToFile(fileName, txStr, podID, ReverseTodayLogger, reverseLogger)
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

func WriteLogFile(fileName string, txStr string, stdLogger stdLogger, log *log.Logger) {
	stdLogger.Filename = fileName
	log.Println(txStr)
}

type LogKafkaMessageToFileFunc func(string, string, string)

func LogKafkaMessageToFile(fileName string, txStr string, podID string, stdLogger stdLogger, log *log.Logger) {
	cleanFileName := GetFileNameFormat(podID, fileName)
	// fmt.Println(cleanFileName)
	WriteLogFile(cleanFileName, txStr, stdLogger, log)
}
