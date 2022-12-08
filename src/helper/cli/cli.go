package cli

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

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
