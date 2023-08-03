package cake

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

type Cake struct {
	Commands []*Command
	parser   *Parser
}

func NewCake(parser *Parser) *Cake {

	return &Cake{
		Commands: []*Command{},
		parser:   parser,
	}
}

func (c *Cake) Init() error {
	cmds, err := c.parser.Parse()
	if err != nil {
		return err
	}

	c.Commands = cmds
	return nil
}

func (c *Cake) Execute(name string) error {

	cmd := findCommand(c.Commands, name)

	if cmd == nil {
		return errors.New("Command not found")
	}

	if cmd.Preq != "" {
		err := c.Execute(cmd.Preq)
		if err != nil {
			return err
		}
	}

	for _, v := range cmd.Commands {
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
		if v.Name == cmd {
			return v
		}
	}
	return nil
}
