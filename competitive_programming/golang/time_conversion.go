package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var time string
	var separatedTime []string
	fmt.Scan(&time)
	separatedTime = strings.Split(time, ":")
	hour, _ := strconv.Atoi(separatedTime[0])
	minute := separatedTime[1]
	second := separatedTime[2][0:2]
	AmPm := regexp.MustCompile(`PM`)
	if AmPm.MatchString(separatedTime[2]) && hour < 12 {
		hour += 12
	} else if AmPm.MatchString(separatedTime[2]) && hour == 12 {
		hour = 12
	} else if hour == 12 {
		hour = 0
	}
	fmt.Printf("%02d:%s:%s\n", hour, minute, second)
}
