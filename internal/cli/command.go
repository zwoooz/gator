package cli

import (
	"fmt"
	"strings"
)

type Command struct {
	Name string
	Args []string
}

// ParseInput turns raw CLI arguments into a Command
// Only the name is lowercased, args are untouched
func ParseInput(input []string) (Command, error) {
	if len(input) < 1 {
		return Command{}, fmt.Errorf("%w: no command specified", ErrUsage)
	}

	return Command{
		Name: strings.ToLower(input[0]),
		Args: input[1:],
	}, nil
}
