// Package cli provides the shared building blocks for the command-line interface
package cli

import "fmt"

type Commands struct {
	handlers map[string]Handler
}

type Handler func(*State, Command) error

func NewCommands() *Commands {
	return &Commands{
		handlers: make(map[string]Handler),
	}
}

func (c *Commands) Run(s *State, cmd Command) error {
	handler, exists := c.handlers[cmd.Name]
	if !exists {
		return fmt.Errorf("command '%s' does not exist", cmd.Name)
	}
	return handler(s, cmd)
}

func (c *Commands) Register(name string, h Handler) {
	c.handlers[name] = h
}
