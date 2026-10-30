package cli

import (
	"fmt"
	"strings"
)

type Command struct {
	Name string
	Args []string
}

func ParseInput(input []string) (Command, error) {
	if len(input) < 1 {
		return Command{}, fmt.Errorf("no command specified")
		// TODO: call help/usage command here once created
	}

	// WARN: This will ruin URLs if we will be taking those as arguments
	for i, str := range input {
		input[i] = strings.ToLower(str)
	}

	return Command{
		Name: input[0],
		Args: input[1:],
	}, nil
}
