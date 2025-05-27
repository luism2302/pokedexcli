package input

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

func getSupportedCommands() map[string]cliCommand {
	supportedCommands := map[string]cliCommand{
		"exit": {name: "exit", description: "Exit the Pokedex", callback: commandExit},
		"help": {name: "help", description: "Displays a help message", callback: commandHelp},
	}
	return supportedCommands
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, command := range getSupportedCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func CleanInput(text string) []string {
	if text == "" {
		return []string{}
	}

	lower := strings.ToLower(text)
	trimmed := strings.Fields(lower)

	return trimmed
}

func Repl() {
	scanner := bufio.NewScanner(os.Stdin)
	flag := true

	for flag {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := CleanInput(scanner.Text())

		if len(words) <= 0 {
			fmt.Println("Unknown command")
			continue
		}
		supportedCommands := getSupportedCommands()
		if _, ok := supportedCommands[words[0]]; !ok {
			fmt.Println("Unkown command")
			continue
		}
		supportedCommands[words[0]].callback()
	}
}
