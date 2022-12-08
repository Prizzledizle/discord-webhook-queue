package cli

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"webhook-queue/src/types"

	"github.com/fatih/color"
)

func ClearCLI() {
	var cmd *exec.Cmd

	switch os := runtime.GOOS; os {
	case "darwin":
		cmd = exec.Command("clear")
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		panic("OS not supported")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func PrintHeader() {
	blue := color.New(color.FgBlue).SprintFunc()

	fmt.Println(
		blue(
			` _       __      ____                      
| |     / /     / __ \__  _____  __  _____ 
| | /| / /_____/ / / / / / / _ \/ / / / _ \
| |/ |/ /_____/ /_/ / /_/ /  __/ /_/ /  __/
|__/|__/      \___\_\__,_/\___/\__,_/\___/ `))
	fmt.Print("\n")
}

func RenameCLI(queues []types.QueueObject) {
	var cmd *exec.Cmd

	nameString := ""

	if len(queues) == 0 {
		nameString = " | No Queues"
	} else {
		for i := range queues {
			nameString += " | " + queues[i].Alias + ": " + strconv.Itoa(queues[i].Length)
		}
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		cmd = exec.Command("osascript", "-e", `tell application "Terminal" to set custom title of first window to "W-Queue`+nameString+`"`)
	case "windows":
		cmd = exec.Command("cmd", "/c", "title", "W-Queue"+nameString)
	default:
		panic("OS not supported")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
