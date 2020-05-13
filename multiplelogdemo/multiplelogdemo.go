package multiplelogdemo

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
	ForcepostFirstWriter  stdLogger
	ForcepostSecondWriter stdLogger
	ReverseWriter         stdLogger
	ForcepostWriter       stdLogger

	reverseFirstLogger    *log.Logger
	reverseSecondLogger   *log.Logger
	forcepostFirstLogger  *log.Logger
	forcepostSecondLogger *log.Logger
	ReverseLogger         *log.Logger
	ForcepostLogger       *log.Logger
)

type LogConfig struct {
	Logfile string
	MaxSize int
	MaxAge  int
}

func generateLumberjack(c *LogConfig) (*log.Logger, stdLogger) {
	l := &lumberjack.Logger{
		Filename:   c.Logfile,
		MaxSize:    c.MaxSize, // megabytes
		MaxAge:     c.MaxAge,  //days
		MaxBackups: 3,         // files
		Compress:   true,      // disabled by default
	}
	logger := log.New(l, "", 1)
	logger.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	writer := stdLogger{l}
	return logger, writer
}

func (c *LogConfig) SetLogger() {
	reverseFirstLogger, ReverseFirstWriter = generateLumberjack(c)
	reverseSecondLogger, ReverseSecondWriter = generateLumberjack(c)

	forcepostFirstLogger, ForcepostFirstWriter = generateLumberjack(c)
	forcepostSecondLogger, ForcepostSecondWriter = generateLumberjack(c)

	ReverseLogger = reverseFirstLogger
	ForcepostLogger = forcepostFirstLogger
	ReverseWriter = ReverseFirstWriter
	ForcepostWriter = ForcepostFirstWriter
}

func CheckKafkaLogWriter() {
	if ReverseWriter == ReverseFirstWriter && ReverseLogger == reverseFirstLogger && ForcepostWriter == ForcepostFirstWriter && ForcepostLogger == forcepostFirstLogger {
		fmt.Println("Using first logger")
		ReverseWriter = ReverseSecondWriter
		ReverseLogger = reverseSecondLogger
		ForcepostWriter = ForcepostSecondWriter
		ForcepostLogger = forcepostSecondLogger
		ReverseFirstWriter.Close()
		ForcepostFirstWriter.Close()
	} else {
		fmt.Println("Using second logger")
		ReverseWriter = ReverseFirstWriter
		ReverseLogger = reverseFirstLogger
		ForcepostWriter = ForcepostFirstWriter
		ForcepostLogger = forcepostFirstLogger
		ReverseSecondWriter.Close()
		ForcepostSecondWriter.Close()
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
	WriteLogFile(cleanFileName, txStr, stdLogger, log)
}
