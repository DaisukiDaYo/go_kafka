package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/robfig/cron"
)

type StdLogger struct {
	*lumberjack.Logger
}

var (
	ReverseWriter   StdLogger
	ForcepostWriter StdLogger

	ReverseLogger   *log.Logger
	ForcepostLogger *log.Logger
	txStr           string
)

type LogConfig struct {
	Logfile string
	MaxSize int
	MaxAge  int
}

func generateLumberjack(c *LogConfig) (*log.Logger, StdLogger) {
	l := &lumberjack.Logger{
		Filename: c.Logfile,
		MaxSize:  c.MaxSize, // megabytes
		MaxAge:   c.MaxAge,  // days
	}
	logger := log.New(l, "", 1)
	logger.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	writer := StdLogger{l}
	return logger, writer
}

func (c *LogConfig) NewSwitchFileWriterFunc(podID, reverseFile, forcepostFile string) func() {
	ReverseLogger, ReverseWriter = generateLumberjack(c)

	ForcepostLogger, ForcepostWriter = generateLumberjack(c)

	ReverseWriter.Filename = getReverseFileNameFormat(podID, reverseFile)
	fmt.Println(ReverseWriter.Filename)
	ForcepostWriter.Filename = getForcepostFileNameFormat(podID, forcepostFile)

	return func() {
		ReverseWriter.Close()
		ForcepostWriter.Close()
		ReverseWriter.Filename = "kafka_error_logs/reverse_rotate_2.txt"
		ForcepostWriter.Filename = getForcepostFileNameFormat(podID, forcepostFile)
	}
}

type WriteReverseMessageToFileFunc func(string)

func WriteReverseMessageToFile(txStr string) {
	ReverseLogger.Println(txStr)
}

type WriteForcepostMessageToFileFunc func(string)

func WriteForcepostMessageToFile(txStr string) {
	ForcepostLogger.Println(txStr)
}

func getReverseFileNameFormat(podID, fileName string) string {
	return fmt.Sprintf(fileName, strings.ToLower(time.Now().Format("Mon_20060201_15:04:05"))+podID)
}

func getForcepostFileNameFormat(podID, fileName string) string {
	return fmt.Sprintf(fileName, time.Now().Format("20060201_")+podID)
}

func main() {
	configLog := LogConfig{
		Logfile: "default.txt",
		MaxSize: 5,
		MaxAge:  3,
	}

	c := cron.New()
	c.AddFunc("0/5 * * * *", func() { fmt.Println("Every second.") })
	c.AddFunc("0/5 * * * *", configLog.NewSwitchFileWriterFunc(
		"1",
		"kafka_error_logs/reverse_rotate_%s.txt",
		"kafka_error_logs/forcepost_rotate_%s.txt",
	))
	c.Start()

	txStr = "test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file end"
	for i := 1; i <= 1005000; i++ {
		WriteReverseMessageToFile(txStr)
	}
	for {
	}

	// switchFileFunc()

	// txStr = "test2 rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file end"
	// for i := 1; i <= 105000; i++ {
	// 	WriteReverseMessageToFile(txStr)
	// }
}
