package main

import (
	"fmt"
	"errors"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide pokemon name")
	}
	name := args[0]
	pokemonResp, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	res := rand.Intn(pokemonResp.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	if res > 40 {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}

	fmt.Printf("%s was caught!\n", name)

	cfg.pokedex[name] = pokemonResp
	return nil
}