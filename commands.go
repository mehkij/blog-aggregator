package main

import "errors"

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

// Register new handler function for a given command name
func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

// Run the given command with the provided state if it exists
func (c *commands) run(s *state, cmd command) error {
	f, ok := c.registeredCommands[cmd.Name]

	if !ok {
		return errors.New("command not found")
	}

	return f(s, cmd)
}
