package main

import (
	"bufio"
	"fmt"
	"jsonte/jsonte"
	"jsonte/jsonte/functions"
	"os"
	"strings"
)

func main() {
	functions.Init()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	for true {
		read, _ := reader.ReadString('\n')
		text := strings.TrimRight(read, "\n\r")
		if text == "exit" {
			break
		}
		jsonte.Eval(text)
		fmt.Print("> ")
	}
}
