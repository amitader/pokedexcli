package main
import (
	"fmt"
	"errors"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}
	name := args[0]
	locationResp, err := cfg.pokeapiClient.ListPokemons(name)
	if err != nil {
		return err
	}
	for _, val := range locationResp.PokemonEncounters {
		fmt.Println(val.Pokemon.Name)
	}
	return nil
}