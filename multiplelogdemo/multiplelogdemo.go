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
	ReverseFirstWriter     stdLogger
	ReverseSecondWriter    stdLogger
	ReverseThirdWriter     stdLogger
	ReverseForthWriter     stdLogger
	ReverseFifthWriter     stdLogger
	ReverseSixthWriter     stdLogger
	ReverseSeventhWriter   stdLogger
	ForcepostFirstWriter   stdLogger
	ForcepostSecondWriter  stdLogger
	ForcepostThirdWriter   stdLogger
	ForcepostForthWriter   stdLogger
	ForcepostFifthWriter   stdLogger
	ForcepostSixthWriter   stdLogger
	ForcepostSeventhWriter stdLogger
	ReverseWriter          stdLogger
	ForcepostWriter        stdLogger

	reverseFirstLogger     *log.Logger
	reverseSecondLogger    *log.Logger
	reverseThirdLogger     *log.Logger
	reverseForthLogger     *log.Logger
	reverseFifthLogger     *log.Logger
	reverseSixthLogger     *log.Logger
	reverseSeventhLogger   *log.Logger
	forcepostFirstLogger   *log.Logger
	forcepostSecondLogger  *log.Logger
	forcepostThirdLogger   *log.Logger
	forcepostForthLogger   *log.Logger
	forcepostFifthLogger   *log.Logger
	forcepostSixthLogger   *log.Logger
	forcepostSeventhLogger *log.Logger
	ReverseLogger          *log.Logger
	ForcepostLogger        *log.Logger
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

func CheckKafkaLogWriter() {
	// Can we use switch case to switching log writer?
	switch time.Now().Weekday() {
	case time.Monday:
		fmt.Println("Use 1 writer")
		ReverseWriter = ReverseFirstWriter
		ReverseLogger = reverseFirstLogger
		ForcepostWriter = ForcepostFirstWriter
		ForcepostLogger = forcepostFirstLogger
		ReverseSeventhWriter.Close()
		ForcepostSeventhWriter.Close()
	case time.Tuesday:
		fmt.Println("Use 2 writer")
		ReverseWriter = ReverseSecondWriter
		ReverseLogger = reverseSecondLogger
		ForcepostWriter = ForcepostSecondWriter
		ForcepostLogger = forcepostSecondLogger
		ReverseFirstWriter.Close()
		ForcepostFirstWriter.Close()
	case time.Wednesday:
		fmt.Println("Use 3 writer")
		ReverseWriter = ReverseThirdWriter
		ReverseLogger = reverseThirdLogger
		ForcepostWriter = ForcepostThirdWriter
		ForcepostLogger = forcepostThirdLogger
		ReverseSecondWriter.Close()
		ForcepostSecondWriter.Close()
	case time.Thursday:
		fmt.Println("Use 4 writer")
		ReverseWriter = ReverseForthWriter
		ReverseLogger = reverseForthLogger
		ForcepostWriter = ForcepostForthWriter
		ForcepostLogger = forcepostForthLogger
		ReverseThirdWriter.Close()
		ForcepostThirdWriter.Close()
	case time.Friday:
		fmt.Println("Use 5 writer")
		ReverseWriter = ReverseFifthWriter
		ReverseLogger = reverseFifthLogger
		ForcepostWriter = ForcepostFifthWriter
		ForcepostLogger = forcepostFifthLogger
		ReverseForthWriter.Close()
		ForcepostForthWriter.Close()
	case time.Saturday:
		fmt.Println("Use 6 writer")
		ReverseWriter = ReverseSixthWriter
		ReverseLogger = reverseSixthLogger
		ForcepostWriter = ForcepostSixthWriter
		ForcepostLogger = forcepostSixthLogger
		ReverseFifthWriter.Close()
		ForcepostFifthWriter.Close()
	case time.Sunday:
		fmt.Println("Use 7 writer")
		ReverseWriter = ReverseSeventhWriter
		ReverseLogger = reverseSeventhLogger
		ForcepostWriter = ForcepostSeventhWriter
		ForcepostLogger = forcepostSeventhLogger
		ReverseSixthWriter.Close()
		ForcepostSixthWriter.Close()
	}
}

func (c *LogConfig) SetLogger() {
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
	reverseThirdLogger = log.New(reverseThirdLumberjack, "", 1)
	reverseThirdLogger.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	ReverseThirdWriter = stdLogger{reverseThirdLumberjack}

	forcepostLumberjack := generateLumberjack(c)
	forcepostFirstLogger = log.New(forcepostLumberjack, "", 1)
	forcepostFirstLogger.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	ForcepostFirstWriter = stdLogger{forcepostLumberjack}

	forcepostSecondLumberjack := generateLumberjack(c)
	forcepostSecondLogger = log.New(forcepostSecondLumberjack, "", 1)
	forcepostSecondLogger.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	ForcepostSecondWriter = stdLogger{forcepostSecondLumberjack}

	forcepostThirdLumberjack := generateLumberjack(c)
	forcepostThirdLogger = log.New(forcepostThirdLumberjack, "", 1)
	forcepostThirdLogger.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	ForcepostThirdWriter = stdLogger{forcepostThirdLumberjack}

	ReverseLogger = reverseFirstLogger
	ForcepostLogger = forcepostFirstLogger
	ReverseWriter = ReverseFirstWriter
	ForcepostWriter = ForcepostFirstWriter
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
