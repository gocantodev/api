package cli

import (
	"errors"
	"fmt"
	"github.com/voyago/environment"
)

type Input struct {
	version  string
	commands []string //the allowed commands to be executed.
	args     []string
	err      error
	env      environment.Env
}

func Parse(args []string) (Input, error) {
	env, err := environment.Make("server")

	if err != nil {
		return Input{}, err
	}

	return Input{
		version:  "0.0.1",
		commands: []string{"new"},
		args:     args,
		err:      nil,
		env:      env,
	}, nil
}

func (i Input) ShouldRunMigrations() bool {
	return len(i.args) == 1
}

func (i Input) GetFileName() string {
	return i.args[0]
}

func (i *Input) ShouldReject() bool {
	if len(i.args) == 2 {
		i.err = errors.New(fmt.Sprintf("Expected [3] args, but only [2] were given."))

		return true
	}

	if len(i.args) > 3 {
		i.err = errors.New(fmt.Sprintf("Expected 3 args, but [%v] were given.", len(i.args)))

		return true
	}

	if len(i.args) == 3 && i.isCommandNotSupported(i.args[1]) {
		i.err = errors.New(fmt.Sprintf("The given command [%v] is invalid. ", i.args[1]))

		return true
	}

	return false
}

func (i Input) GetError() error {
	return i.err
}

func (i Input) GetValue() string {
	return i.args[2]
}

func (i Input) GetEnv() environment.Env {
	return i.env
}

func (i Input) isCommandNotSupported(command string) bool {
	for _, v := range i.commands {
		if v == command {
			return false
		}
	}

	return true
}
