package engine

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

var runMap = map[string]func(){
	"exit": exit,
	"help": help,
	"about": about,
}


func Run() {
	green := color.New(color.FgGreen).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Fprint(os.Stdout, green("🤖  How can I help?\nYou> "))
		message, _ := reader.ReadString('\n')
		message = blue(strings.TrimSpace(message))


		if f, ok := runMap[message]; ok {
			f()
		} else {
			// send to GPT-3
			fmt.Print("sending message: " + message + "\n")
			// receive and print response from GPT-3
			fmt.Print("GPT-3: <response> \n")
			// assess response for correctness
			// use user feedback to decide whether to google search
				// if google search, send to google
			// search plugins for additional functionality
		}
	}
}

func exit() {
	os.Exit(0)
}

func help() {
	fmt.Println("help")
}

func about() {
	fmt.Println("about")
}
