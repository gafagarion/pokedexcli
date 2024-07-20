package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
)

func commandCatch(out io.Writer, in io.Reader, cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("must provide exactly one pokemon name")
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(&args[0])
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "Throwing a Pokeball at %s...\n", pokemon.Name)

	res := rand.Intn(pokemon.BaseExperience)

	if res > 40 {
		cfg.caughtPokemon[pokemon.Name] = pokemon
		fmt.Fprintf(out, "%s was caught!\n", pokemon.Name)
	} else {
		fmt.Fprintf(out, "%s escaped!\n", pokemon.Name)
	}

	return nil

}
