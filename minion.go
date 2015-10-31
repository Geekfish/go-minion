package main

import (
	//"fmt"
	"time"
	"github.com/everdev/mack"
	"strconv"
)

func main() {
	response, err := mack.Dialog("Enter a count down", "Count down", "10")
	if err != nil {
		panic(err)
	}
	if response.Clicked == "Cancel" {
		return
	} else {
		mack.Notify("Will notify every " + response.Text + " seconds")
		interval, _ := strconv.Atoi(response.Text)
		sleepTime := time.Duration(interval) * time.Second
		for {
			time.Sleep(sleepTime)
			mack.Notify("Posture check? Looking good!")
		}
	}
}
