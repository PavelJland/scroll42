package main

import (
	"os"
	"fmt"
	"time"
	"strconv"
	"github.com/go-vgo/robotgo"
)

func main() {
	ss, err := strconv.Atoi(os.Args[1])
	if (err != nil) {
		fmt.Println("How large is scroll?");
	}
	to, err2 := strconv.ParseInt(os.Args[2], 10, 32)
	if (err2 != nil) {
		fmt.Println("What is timeout time?");	
	}
	for {
		time.Sleep(time.Duration((int32(to))) * time.Second);
		robotgo.ScrollMouse(ss, "down")
	}
}
