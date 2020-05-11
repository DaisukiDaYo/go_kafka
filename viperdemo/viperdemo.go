package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

func GetFileNameFormat(podID, fileName string) string {
	now := time.Now()

	fileFormat := now.Format("20060201") + "_" + podID
	if strings.Contains(fileName, "reverse") == true {
		fileFormat = strings.ToLower(now.Format("Mon")) + "_" + fileFormat
	}

	cleanFileName := fmt.Sprintf(fileName, fileFormat)
	return cleanFileName
}

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath("../")
	viper.AutomaticEnv()

	// If a config file is found, read it in
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(viper.Get("Kafka.Brokers"))
	fmt.Println(viper.GetString("PODID"))
	fmt.Println(viper.GetString("podid"))
	fmt.Println(viper.GetString("underissexy"))
	// forcePostFileName := viper.GetString("Kafka.ForcePostModule.MessageFile")

	// now := time.Now()

	// fileName := fmt.Sprintf(
	// 	forcePostFileName,
	// 	now.Format("20060101"),
	// 	viper.GetString("PODID"),
	// )

	// fmt.Println(fileName)

	podID := viper.GetString("PODID")
	// fileName := viper.GetString("Kafka.ReverseModule.MessageFile")
	fileName := viper.GetString("Kafka.ForcePostModule.MessageFile")

	cleanFileName := GetFileNameFormat(podID, fileName)

	fmt.Println(cleanFileName)
}
