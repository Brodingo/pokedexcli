package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	// for stat := range pokemon.Stats {
	// 	fmt.Printf("%s: %d\n", pokemon.Stats[stat].Stat.Name, pokemon.Stats[stat].BaseStat)
	// 	fmt.Printf("Effort: %d\n", pokemon.Stats[stat].Effort)
	// }
	// fmt.Printf("------------\n")
	// fmt.Printf("Base exp: %d\n", pokemon.BaseExperience)

	statTotal := 0
	for _, stat := range pokemon.Stats {
		if stat.Effort > 0 {
			statTotal += (stat.BaseStat * stat.Effort)
			// fmt.Printf("Boosted %s: %d -> %d\n", stat.Stat.Name, stat.BaseStat, statTotal)
		} else {
			statTotal += stat.BaseStat
			// fmt.Printf("%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
	}
	// fmt.Printf("Stat total: %d\n", statTotal)

	// res := rand.Intn(pokemon.BaseExperience)

	rate := pokemon.BaseExperience + statTotal
	res := rand.Intn(rate)
	const maxCatchRate = 100

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	fmt.Printf("Luck: %d\n", maxCatchRate-res)

	// rate := 1.0 / math.Pow(math.Log(float64(pokemon.BaseExperience+statTotal)+1), 1.1)
	// fmt.Printf("Catch rate calc: %v\n", rate)
	// randNum := rand.Float64()
	// fmt.Printf("Random number: %f\n", randNum)
	if res > maxCatchRate {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)

	cfg.caughtPokemon[pokemon.Name] = pokemon
	return nil
}
