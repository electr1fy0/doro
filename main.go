package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "doro",
	Short: "a cli clock tool",
}

var startTimerCmd = &cobra.Command{
	Use:   "timer [duration]",
	Short: "start a timer e.g. doro timer 30m",
	Run: func(cmd *cobra.Command, args []string) {
		startTimer(args[0])
	},
	Args: cobra.ExactArgs(1),
}

var startStopwatchCmd = &cobra.Command{
	Use:   "stopwatch",
	Short: "start a stopwatch",
	Run: func(cmd *cobra.Command, args []string) {
		startStopWatch()
	},
}

func startTimer(duration string) {
	diff, _ := time.ParseDuration(duration)
	end := time.Now().Add(diff)
	ticker := time.NewTicker(1 * time.Second)
	timer := time.NewTimer(diff)
	defer ticker.Stop()

	fmt.Println("Starting a timer for", diff)
	for {
		select {
		case <-ticker.C:
			remaining := time.Until(end)
			remaining = max(0, remaining)

			fmt.Printf("\rTime left: %-5v", remaining.Truncate(time.Second))
		case <-timer.C:
			fmt.Println("Timer is up")
			return
		}
	}
}

func startStopWatch() {
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	start := time.Now()
	fmt.Println("Stopwatch is running... (<ctrl-c> to stop)")

	for {
		<-ticker.C
		elapsed := time.Since(start)

		fmt.Printf("\rElapsed: %-8v", elapsed.Truncate(time.Millisecond))
	}
}

func main() {
	cmd.AddCommand(startTimerCmd)
	cmd.AddCommand(startStopwatchCmd)

	cmd.Execute()
}
