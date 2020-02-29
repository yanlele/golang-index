package main

import "fmt"

type Command struct {
	Name    string
	Var     *int
	Comment string
}

func newCommand(name string, varref *int, comment string) *Command {
	return &Command{
		Name:    name,
		Var:     varref,
		Comment: comment,
	}
}

func main() {
	version := 12
	cmd := newCommand("version", &version, "show  version")

	cmd2 := &Command{
		Name:    "version",
		Var:     &version,
		Comment: "show version",
	}

	fmt.Println(cmd.Name)
	fmt.Println(cmd2.Name)
}
