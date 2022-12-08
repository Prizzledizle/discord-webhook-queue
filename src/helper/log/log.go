package log

import (
	"fmt"

	"github.com/fatih/color"
)

func Error(message string, alias string) {
	red := color.New(color.FgRed).SprintFunc()

	if alias != "" {
		fmt.Println(red("["+alias+"]	", message))
	} else {
		fmt.Println(red(message))
	}
}

func Info(message string, alias string) {
	cyan := color.New(color.FgCyan).SprintFunc()

	if alias != "" {
		fmt.Println(cyan("["+alias+"]	", message))
	} else {
		fmt.Println(cyan(message))
	}
}

func Success(message string, alias string) {
	green := color.New(color.FgGreen).SprintFunc()

	if alias != "" {
		fmt.Println(green("["+alias+"]	", message))
	} else {
		fmt.Println(green(message))
	}
}

func Warning(message string, alias string) {
	yellow := color.New(color.FgYellow).SprintFunc()

	if alias != "" {
		fmt.Println(yellow("["+alias+"]	", message))
	} else {
		fmt.Println(yellow(message))
	}
}
