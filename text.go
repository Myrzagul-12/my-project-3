package main

import (
	"fmt"
	"time"
)

func FormatDate(t time.Time) string {
	return t.Format("02.01.2006 15:04")
}

func main() {

	
	now := time.Now()
	fmt.Println(FormatDate(now)) 
}

