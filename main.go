package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "doro",
	Short: "a cli pomodoro tool",
}

var startCmd = &cobra.Command{
	Use:   "start [duration]",
	Short: "start a new timer e.g. doro new 30m",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println(args[0])
		startTimer(args[0])
	},
	Args: cobra.ExactArgs(1),
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

			fmt.Println("Time left: ", remaining.Truncate(time.Second))
		case <-timer.C:
			fmt.Println("Timer is up")
			return
		}
	}

}

func main() {
	cmd.AddCommand(startCmd)
	cmd.Execute()

}
