// Package cli provides the shared building blocks for the command-line interface
package cli

import (
	"fmt"
	"sort"
)

// Handler executes a command
type Handler func(*State, Command) error

// Spec describes a command
type Spec struct {
	Name        string
	Category    string
	Usage       string
	Description string
	Handler     Handler
}

// Commands is the registry of all available commands
type Commands struct {
	specs map[string]Spec
}

var categoryOrder = []string{"Util", "Users"}

func NewCommands() *Commands {
	return &Commands{
		specs: make(map[string]Spec),
	}
}

// Run executes command by name, returning ErrUsage-wrapped error if it doesn't exist
func (c *Commands) Run(s *State, cmd Command) error {
	spec, ok := c.Lookup(cmd.Name)
	if !ok {
		return fmt.Errorf("%w: command '%s' does not exist", ErrUsage, cmd.Name)
	}
	return spec.Handler(s, cmd)
}

// Register adds a command to the registry
func (c *Commands) Register(spec Spec) {
	c.specs[spec.Name] = spec
}

// Lookup returns the spec of the command under name, if it exists
func (c *Commands) Lookup(name string) (Spec, bool) {
	spec, ok := c.specs[name]
	return spec, ok
}

// Specs returns all commands ordered by category, then name
func (c *Commands) Specs() []Spec {
	out := make([]Spec, 0, len(c.specs))
	for _, spec := range c.specs {
		out = append(out, spec)
	}
	sort.Slice(out, func(i, j int) bool {
		ri := categoryRank(out[i].Category)
		rj := categoryRank(out[j].Category)
		if ri != rj {
			return ri < rj
		}
		return out[i].Name < out[j].Name
	})

	return out
}

// categoryRank returns a category's position in categoryOrder (unknown ranks last)
func categoryRank(category string) int {
	for i, name := range categoryOrder {
		if name == category {
			return i
		}
	}
	return len(categoryOrder)
}
