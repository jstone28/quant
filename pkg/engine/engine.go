package engine

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var runMap = map[string]func(){
	"exit": exit,
	"help": help,
	"about": about,
}


func Run() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("You: ")
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

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
