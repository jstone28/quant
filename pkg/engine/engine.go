package engine

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	oai "github.com/jstone28/quant/pkg/llm/openai"
)

var runMap = map[string]func(){
	"exit":  exit,
	"help":  help,
	"about": about,
}

func Run() {
	green := color.New(color.FgGreen).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Fprint(os.Stdout, green("🤖  How can I help?\nYou> "))
		prompt, _ := reader.ReadString('\n')
		prompt = blue(strings.TrimSpace(prompt))

		if f, ok := runMap[prompt]; ok {
			f() // check for exit, help, about
		} else {
			// send to GPT-3
			resp := oai.Prompt(prompt)
			// receive and print response from GPT-3

			fmt.Println("🤖  ", resp)
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
