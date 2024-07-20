package main

import (
	"fmt"
	"io"
)

func commandHelp(out io.Writer, in io.Reader, cfg *config, args ...string) error {
	_, err := out.Write([]byte("This is the display for the help command\n"))
	if err != nil {
		return err
	}

	for _, cmd := range getCommands() {
		_, err := fmt.Fprintf(out, "%s: %s\n", cmd.name, cmd.description)
		if err != nil {
			return err
		}
	}
	fmt.Fprintf(out, "\n")
	return nil
}
