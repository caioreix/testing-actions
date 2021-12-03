package main

import (
	"errors"
	"strings"
	"time"
)

func Sum(x, y int) int {
	return x + y
}

func Sub(x, y int) int {
	return x - y
}

var erroDeMultiplasChamadas = errors.New("deu ruim aqui")

type Command struct {
	ID          string
	Name        string
	Aliases     []string
	Subcommands []string
	Args        []string
	Help        string

	RootOnly bool

	Cooldown time.Duration

	Errs []error
}

func (c *Command) SetArgs(args ...string) *Command {
	var jaexiste bool
	for _, arg := range args {
		if jaexiste {
			c.Errs = []error{erroDeMultiplasChamadas}
			return c
		}

		if strings.HasPrefix(arg, "...") {
			jaexiste = true
		}
	}
	c.Args = args
	return c
}
