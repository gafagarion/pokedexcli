package main

import (
	"errors"
	"fmt"
	"io"
)

func commandInspect(out io.Writer, in io.Reader, cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("must provide exactly one pokemon")
	}

	if pokemon, ok := cfg.caughtPokemon[args[0]]; ok {
		fmt.Fprintf(out, "Name: %s\n", pokemon.Name)
		fmt.Fprintf(out, "Height: %d\n", pokemon.Height)
		fmt.Fprintf(out, "Weight: %d\n", pokemon.Weight)
		fmt.Fprintf(out, "Stats:\n")
		for _, stat := range pokemon.Stats {
			fmt.Fprintf(out, "  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Fprintf(out, "Types:\n")
		for _, typ := range pokemon.Types {
			fmt.Fprintf(out, "  -%s\n", typ.Type.Name)
		}
	} else {
		fmt.Fprintf(out, "you have not caught that pokemon\n")
	}

	return nil
}
