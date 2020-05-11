package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

// ยังงงอยู่ แต่ก็เข้าใจ 55555 เราจะไปหาวิธีอื่นพลาง
// ขอลองแปบบบ อ่าเช
// ลอง add method เพิ่มใน interface

type MyLogger lumberjack.Logger

func (l *MyLogger) SetFileName() {
	fmt.Println(l.Filename)
}

func dump() {
	now := time.Now().UTC()
	serverEnv := "dev"
	podID := "podID"

	fileName := fmt.Sprintf(
		serverEnv+"_reverse_%s_%s_"+podID+".txt",
		strings.ToLower(now.Format("Mon")),
		now.Format("20060101"),
	)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)

	go func() {
		for {
			<-c
			l.Rotate()
		}
	}()

	// w := bufio.NewWriter(l)

	fmt.Println(l.MaxSize)

	for i := 0; i <= 100000000; i++ {
		// text := "buffer lumberjack"
		log.Println("buffer lumberjack")
	}
}

func initialLogRotate() {
	l := &lumberjack.Logger{
		MaxSize:    1,    // megabytes
		MaxBackups: 3,    // files
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	}
	fmt.Println(&l)
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	log.SetOutput(l)
}
func main() {
	initialLogRotate()
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	parent := filepath.Dir(wd)
	fmt.Println(wd)
	fmt.Println(parent)
	fmt.Println(parent + "/app/kafka_error_logs")

	fmt.Println(os.Getwd())
	now := time.Now().UTC()
	serverEnv := "dev"
	podID := "podID"

	fileName := fmt.Sprintf(
		serverEnv+"_reverse_%s_%s_"+podID+".txt",
		strings.ToLower(now.Format("Mon")),
		now.Format("20060101"),
	)
	txStr := "halo woohoo!!!!"
	WriteLogFile(fileName, string(txStr))
}

func WriteLogFile(fileName string, txStr string) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	parent := filepath.Dir(wd)
	filePath := parent + "/app/kafka_error_logs/"
	fmt.Println(filePath)
	fmt.Println("********************************")
	l := MyLogger()
	l.SetFileName()

	fmt.Println(l)
	fmt.Println("********************************")

}

func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = defaultValue
	}
	return value
}
