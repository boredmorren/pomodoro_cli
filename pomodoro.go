package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/schollz/progressbar/v3"
)

func main() {
	workTimeMin := flag.Int("wm", 0, "Set minutes of work session")
	restTimeMin := flag.Int("rm", 0, "Set minutes of rest session")

	workTimeHour := flag.Int("wh", 0, "Set hours of work session")
	restTimeHour := flag.Int("rh", 0, "Set hours of rest session")

	flag.Parse()

	var workTime int
	var restTime int

	if *workTimeHour != 0 {
		workTime = *workTimeHour * 60
	}

	if *restTimeHour != 0 {
		restTime = *restTimeHour * 60
	}

	workTime += *workTimeMin
	restTime += *restTimeMin

	barWork := progressbar.NewOptions(workTime,
		progressbar.OptionSetDescription("WORK"),
		progressbar.OptionSetWriter(os.Stdout),
		progressbar.OptionSetRenderBlankState(true),
		progressbar.OptionShowCount(),
		progressbar.OptionSetPredictTime(false),
	)
	for i := 0; i < workTime; i++ {
		time.Sleep(60 * time.Second)
		barWork.Add(1)
	}

	fmt.Print("\a")

	time.Sleep(5 * time.Second)

	barWork.Finish()

	barRest := progressbar.NewOptions(restTime,
		progressbar.OptionSetDescription("REST"),
		progressbar.OptionSetWriter(os.Stdout),
		progressbar.OptionSetRenderBlankState(true),
		progressbar.OptionShowCount(),
		progressbar.OptionSetPredictTime(false),
	)

	for i := 0; i < restTime; i++ {
		time.Sleep(60 * time.Second)
		barRest.Add(1)
	}

	barRest.Finish()
}
