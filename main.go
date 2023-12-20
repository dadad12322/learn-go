package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 200; i++ {
		// Use time.Now() for the current date
		currentTime := time.Now()
		// Subtract 'i' days
		d := currentTime.AddDate(0, 0, -i).Format("2006-01-02")
		randNum := rand.Intn(11) + 1

		// Write to test.txt
		file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()
		if _, err := file.WriteString(d + "\n"); err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

		// Git commands
		cmdAdd := exec.Command("git", "add", "test.txt")
		cmdAdd.Run()

		cmdCommit := exec.Command("git", "commit", "--date", fmt.Sprintf("%sT%02d:00:00", d, randNum), "-m", "1")
		cmdCommit.Run()
	}

	// Git push
	cmdPush := exec.Command("git", "push", "-u", "origin", "main")
	cmdPush.Run()
}
