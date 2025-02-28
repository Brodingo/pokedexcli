package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "List all available commands",
			callback:    commandHelp,
		},
	}
}

func startRepl() {
	// Scanner for user input
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// Prompt user for input
		fmt.Print("Pokedex > ")
		// Scan for user input
		scanner.Scan()
		// Store user input
		input := scanner.Text()
		// Clean user input
		words := cleanInput(input)
		// Store the first word of the input
		commandName := words[0]
		// args := words[1:]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Printf("Command %s not found\n", commandName)
			continue
		}

	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
