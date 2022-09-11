package main

import (
	"bufio"
	"fmt"
	"github.com/MCDevKit/jsonte/jsonte"
	"github.com/MCDevKit/jsonte/jsonte/functions"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/gammazero/deque"
	"os"
	"strings"
)

func main() {
	functions.Init()
	repl()
}

func repl() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	for true {
		read, _ := reader.ReadString('\n')
		text := strings.TrimRight(read, "\n\r")
		if text == "exit" {
			break
		}
		eval, err := jsonte.Eval(text, deque.Deque[interface{}]{}, "#/")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(utils.ToString(eval.Value))
		}
		fmt.Print("> ")
	}
}
