package main
import (
	"fmt"
	"errors"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide pokemon name")
	}
	name := args[0]
	if val, ok := cfg.pokedex[name]; ok {
		fmt.Printf("Name: %s\n", val.Name)
		fmt.Printf("Height: %d\n", val.Height)
		fmt.Printf("Weight: %d\n", val.Weight)
		fmt.Printf("Stats:\n")
		for _, stats := range val.Stats {
			fmt.Printf("  -%s %d\n", stats.Stat.Name, stats.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, typ := range val.Types {
			fmt.Printf("  - %s\n", typ.Type.Name)
		}
		return nil
	}
	return nil
}