package cake

import (
	"os"
	"strings"
)

type Command struct {
	Comment  string
	Name     string
	Cursor   int
	Preq     string
	Commands []string
}

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse() ([]*Command, error) {
	b, err := p.readFile()
	if err != nil {
		return nil, err
	}
	f := strings.Split(string(b), "\n")

	cmds := []*Command{}

	i := 0
	for i < len(f) { // find names, preqs and main location of titles
		v := f[i]
		if s := strings.Split(v, ":"); len(s) > 1 && s[0] != v { // if the content has no :, it returns a slice with v

			name := s[0]
			preq := ""

			if len(s) >= 2 {
				preq = strings.TrimSpace(s[1])
			}

			cmds = append(cmds, &Command{
				Name:   name,
				Preq:   preq,
				Cursor: i,
			})
		}
		i++
	}
	for i, v := range cmds { // pack command lines between names
		if i == len(cmds)-1 {
			v.Commands = f[v.Cursor+1:]
			break
		}
		v.Commands = f[v.Cursor+1 : cmds[i+1].Cursor]
	}

	return cmds, nil
}

func (p *Parser) readFile() ([]byte, error) {
	b, err := os.ReadFile("Cakefile")
	if err != nil {
		return nil, err
	}
	return b, nil
}
