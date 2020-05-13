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
	ReverseFirstWriter    stdLogger
	ReverseSecondWriter   stdLogger
	ReverseThirdWriter    stdLogger
	ForcepostFirstWriter  stdLogger
	ForcepostSecondWriter stdLogger
	ForcepostThirdWriter  stdLogger
	reverseWriter         stdLogger
	forcepostWriter       stdLogger

	reverseFirstLogger    *log.Logger
	reverseSecondLogger   *log.Logger
	reverseThirdLogger    *log.Logger
	forcepostFirstLogger  *log.Logger
	forcepostSecondLogger *log.Logger
	forcepostThirdLogger  *log.Logger
	reverseLogger         *log.Logger
	forcepostLogger       *log.Logger
)

type LogConfig struct {
	Logfile string
	MaxSize int
	MaxAge  int
}

func generateLumberjack(c *LogConfig) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   c.Logfile,
		MaxSize:    c.MaxSize, // megabytes
		MaxAge:     c.MaxAge,  //days
		MaxBackups: 3,         // files
		Compress:   true,      // disabled by default
	}
}

func (c *LogConfig) SetLogger() (*log.Logger, *log.Logger, stdLogger, stdLogger) {
	fmt.Printf("Sending log messages to: %s\n", c.Logfile)
	reverseLumberjack := generateLumberjack(c)
	reverseFirstLogger = log.New(reverseLumberjack, "", 1)
	reverseFirstLogger.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	ReverseFirstWriter = stdLogger{reverseLumberjack}

	reverseSecondLumberjack := generateLumberjack(c)
	reverseSecondLogger = log.New(reverseSecondLumberjack, "", 1)
	reverseSecondLogger.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	ReverseSecondWriter = stdLogger{reverseSecondLumberjack}

	reverseThirdLumberjack := generateLumberjack(c)
	reverseThirdLogger := log.New(reverseThirdLumberjack, "", 1)
	reverseThirdLogger.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	ReverseThirdWriter = stdLogger{reverseThirdLumberjack}

	forcepostLumberjack := generateLumberjack(c)
	forcepostFirstLogger := log.New(forcepostLumberjack, "", 1)
	forcepostFirstLogger.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	ForcepostFirstWriter = stdLogger{forcepostLumberjack}

	forcepostSecondLumberjack := generateLumberjack(c)
	forcepostSecondLogger := log.New(forcepostSecondLumberjack, "", 1)
	forcepostSecondLogger.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	ForcepostSecondWriter = stdLogger{forcepostSecondLumberjack}

	forcepostThirdLumberjack := generateLumberjack(c)
	forcepostThirdLogger := log.New(forcepostThirdLumberjack, "", 1)
	forcepostThirdLogger.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	ForcepostThirdWriter = stdLogger{forcepostThirdLumberjack}

	return reverseFirstLogger, forcepostFirstLogger, ReverseFirstWriter, ForcepostFirstWriter
}

func main() {
	configLog := LogConfig{
		Logfile: "default.txt",
		MaxSize: 5,
		MaxAge:  3,
	}
	reverseLogger, forcepostLogger, reverseWriter, forcepostWriter := configLog.SetLogger()

	fileName := "app/kafka_error_logs/reverse_rotate_%s.txt"
	txStr := "test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file end"
	podID := "1243"
	for i := 1; i <= 105000; i++ {
		LogKafkaMessageToFile(fileName, txStr, podID, reverseWriter, reverseLogger)
	}

	fileName = "app/kafka_error_logs/force_post_rotate_2_%s.txt"
	txStr = "2test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file end"
	podID = "1245"
	for i := 1; i <= 105000; i++ {
		LogKafkaMessageToFile(fileName, txStr, podID, forcepostWriter, forcepostLogger)
	}

	t := time.Now()
	rounded := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	fmt.Println(t)
	fmt.Println(rounded)
	if true {
		reverseWriter = ReverseSecondWriter
		reverseLogger = reverseSecondLogger
		ReverseFirstWriter.Close()
	}

	// Can we use switch case to switching log writer?
	switch time.Now().Weekday() {
	case time.Monday:
		fmt.Println("Use 1 writer")
	case time.Tuesday:
		fmt.Println("Use 2 writer")
	case time.Wednesday:
		fmt.Println("Use 3 writer")
	case time.Thursday:
		fmt.Println("Use 4 writer")
	case time.Friday:
		fmt.Println("Use 5 writer")
	case time.Saturday:
		fmt.Println("Use 6 writer")
	case time.Sunday:
		fmt.Println("Use 7 writer")
	}

	fileName = "app/kafka_error_logs/reverse_rotate_tmr_%s.txt"
	txStr = "3test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file end"
	podID = "1243"
	for i := 1; i <= 105000; i++ {
		LogKafkaMessageToFile(fileName, txStr, podID, reverseWriter, reverseLogger)
	}

	if true {
		reverseWriter = ReverseFirstWriter
		reverseLogger = reverseFirstLogger
	}

	fileName = "app/kafka_error_logs/reverse_rotate_3_%s.txt"
	txStr = "4test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file end"
	podID = "1243"
	for i := 1; i <= 105000; i++ {
		LogKafkaMessageToFile(fileName, txStr, podID, reverseWriter, reverseLogger)
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
