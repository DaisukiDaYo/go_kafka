package main-bak

import (
	"fmt"
	"go-kafka/multiplelogdemo"

	"github.com/robfig/cron/v3"
)

func bak() {
	c := cron.New()
	//c.AddFunc("@daily", multiplelogdemo.CheckKafkaLogWriter())
	c.AddFunc("* * * * *", func() { fmt.Println("Every minutes.") })
	c.AddFunc("* * * * *", multiplelogdemo.CheckKafkaLogWriter)
	c.Start()
	fmt.Println(c.Entries())

	configLog := multiplelogdemo.LogConfig{
		Logfile: "default.txt",
		MaxSize: 5,
		MaxAge:  3,
	}
	configLog.SetLogger()

	fileName := "kafka_error_logs/reverse_rotate_%s.txt"
	txStr := "test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file end"
	podID := "1243"
	for i := 1; i <= 105000; i++ {
		multiplelogdemo.LogKafkaMessageToFile(
			fileName,
			txStr,
			podID,
			multiplelogdemo.ReverseWriter,
			multiplelogdemo.ReverseLogger)
	}

	fileName = "kafka_error_logs/force_post_rotate_2_%s.txt"
	txStr = "2test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file end"
	podID = "1245"
	for i := 1; i <= 105000; i++ {
		multiplelogdemo.LogKafkaMessageToFile(
			fileName,
			txStr,
			podID,
			multiplelogdemo.ForcepostWriter,
			multiplelogdemo.ForcepostLogger)
	}

	multiplelogdemo.CheckKafkaLogWriter()

	fileName = "kafka_error_logs/reverse_rotate_tmr_%s.txt"
	txStr = "3test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file end"
	podID = "1243"
	for i := 1; i <= 105000; i++ {
		multiplelogdemo.LogKafkaMessageToFile(
			fileName,
			txStr,
			podID,
			multiplelogdemo.ReverseWriter,
			multiplelogdemo.ReverseLogger)
	}

	fileName = "kafka_error_logs/force_post_rotate_tmr_%s.txt"
	txStr = "4test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file test rolling file end"
	podID = "1245"
	for i := 1; i <= 105000; i++ {
		multiplelogdemo.LogKafkaMessageToFile(
			fileName,
			txStr,
			podID,
			multiplelogdemo.ForcepostWriter,
			multiplelogdemo.ForcepostLogger)
	}

	for {}
}
