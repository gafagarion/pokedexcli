package main

import (
	"errors"
	"fmt"
	"io"
)

func commandPokedex(out io.Writer, in io.Reader, cfg *config, args ...string) error {
	if len(args) != 0 {
		return errors.New("pokedex command doesn't need any argument")
	}

	if len(cfg.caughtPokemon) == 0 {
		fmt.Fprintf(out, "Your pokedex is empty\n")
		return nil
	}

	fmt.Fprintf(out, "Your pokedex:\n")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Fprintf(out, " - %s\n", pokemon.Name)
	}

	return nil

}
