package main

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

type Cake struct {
	commands []*Command
	parser   *Parser
}

func NewCake(parser *Parser) *Cake {

	return &Cake{
		commands: []*Command{},
		parser:   parser,
	}
}

func (c *Cake) Init() error {
	cmds, err := c.parser.Parse()
	if err != nil {
		return err
	}

	c.commands = cmds
	return nil
}

func (c *Cake) Execute(name string) error {

	cmd := findCommand(c.commands, name)

	if cmd == nil {
		return errors.New("Command not found")
	}

	if cmd.preq != "" {
		err := c.Execute(cmd.preq)
		if err != nil {
			return err
		}
	}

	for _, v := range cmd.commands {
		n := strings.TrimSpace(v)
		if n == "" {
			continue
		}
		cmdname := strings.Split(n, " ")[0]
		args := n[len(cmdname):]
		cmd := exec.Command(cmdname, args)
		out, err := cmd.Output()
		if err != nil {
			return err
		}
		os.Stdout.Write(out)
	}

	return nil
}

func findCommand(cmds []*Command, cmd string) *Command {
	for _, v := range cmds {
		if v.name == cmd {
			return v
		}
	}
	return nil
}
