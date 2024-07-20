package main

import (
	"errors"
	"fmt"
	"io"
)

func commandExplore(out io.Writer, in io.Reader, cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("no location provided")
	}
	if len(args) > 1 {
		return errors.New("too many locations provided")
	}
	_, err := fmt.Fprintf(out, "Exploring %s...", args[0])
	if err != nil {
		return err
	}

	location, err := cfg.pokeapiClient.DetailLocation(&args[0])
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(out, "Found Pokemon:")
	if err != nil {
		return err
	}

	for _, enc := range location.PokemonEncounters {
		_, err := fmt.Fprintf(out, "\n - %s", enc.Pokemon.Name)
		if err != nil {
			return err
		}
	}
	_, err = fmt.Fprintf(out, "\n")
	if err != nil {
		return err
	}
	return nil

}
