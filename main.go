package main

import (
	"bufio"
	"fmt"
	"github.com/luism2302/pokedexcli/text"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		hasInput := scanner.Scan()
		if hasInput {
			userInput := scanner.Text()
			command := text.CleanInput(userInput)[0]
			fmt.Printf("Your command was: %s\n", command)
		}
	}

}
