package main

import (
	"io"
	"os"
)

func commandExit(out io.Writer, in io.Reader, cfg *config, args ...string) error {
	_, err := out.Write([]byte("Thank you for using Pokedex CLI\n"))
	if err != nil {
		return err
	}
	os.Exit(1)
	return nil
}
