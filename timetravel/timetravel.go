package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().UTC()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Date())
	fmt.Println(now.Day())
	fmt.Println(now.Weekday())
	fmt.Println(now.Weekday())
	fmt.Println(now.Format("%a"))
	fmt.Println(now.Format("%Y%m%d"))

	// {env}_reverse_mon_{yyyymmdd}_{podid}.txt
	// t := time.Now() formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",         t.Year(), t.Month(), t.Day(),         t.Hour(), t.Minute(), t.Second())
	// reverse_format := "dev_reverse_" + now.Format("Mon") + "_" + now.Year() + now.Month() + now.Day()
	// fmt.Println(reverse_format)
	reverse := fmt.Sprintf("dev_reverse_%s_%d%s%d_podid", now.Format("Mon"), now.Year(), now.Format("01"), now.Day())
	fmt.Println(reverse)

}
