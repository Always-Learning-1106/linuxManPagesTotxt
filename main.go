package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	commands := []string{}
	file, err := os.Open("commands.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		commands = append(commands, scanner.Text())

	}
	for _, command := range commands {
		cache := make(map[string]int)

		if cache[command] < 1 {
			cache[command]++
			commandString := fmt.Sprintf("mkdir %[1]v && cd %[1]v && man %[1]v >> %[1]v.txt && cd ..", command)

			cmd1 := exec.Command("bash", "-c", commandString)
			cmd1.Run()
		}
	}
}
