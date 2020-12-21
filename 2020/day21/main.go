package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

type food struct {
	ingredients []string
	allergens   []string
}

func part1(input []string) (int, map[string][]string) {

	ingredientsRE := regexp.MustCompile(`(.*)\(contains`)
	allergenRE := regexp.MustCompile(`\(contains (.*?)\)`)

	foods := []food{}
	allergenToIngredients := make(map[string][]string)

	for _, rawFood := range input {

		ingredientsMatch := ingredientsRE.FindStringSubmatch(rawFood)
		ingredients := strings.Fields(ingredientsMatch[1])

		allergenMatch := allergenRE.FindStringSubmatch(rawFood)
		allergensRAW := strings.Split(allergenMatch[1], ",")
		allergens := []string{}

		for _, a := range allergensRAW {
			allergen := strings.TrimSpace(a)
			allergens = append(allergens, allergen)
			if _, ok := allergenToIngredients[allergen]; ok {
				allergenToIngredients[allergen] = intersection(allergenToIngredients[allergen], ingredients)
			} else {
				allergenToIngredients[allergen] = ingredients
			}
		}

		foods = append(foods, food{ingredients, allergens})
	}

	ingredientsWithAllergens := []string{}
	for _, allergenIngredients := range allergenToIngredients {
		for _, allergenIngredient := range allergenIngredients {
			if !util.StringInSlice(allergenIngredient, ingredientsWithAllergens) {
				ingredientsWithAllergens = append(ingredientsWithAllergens, allergenIngredient)
			}
		}
	}

	ingredientsFreeFromAllergens := []string{}
	for _, food := range foods {
		for _, ingredient := range food.ingredients {
			if !util.StringInSlice(ingredient, ingredientsWithAllergens) && !util.StringInSlice(ingredient, ingredientsFreeFromAllergens) {
				ingredientsFreeFromAllergens = append(ingredientsFreeFromAllergens, ingredient)
			}
		}
	}

	count := 0
	for _, food := range foods {
		for _, ingredient := range food.ingredients {
			if util.StringInSlice(ingredient, ingredientsFreeFromAllergens) {
				count++
			}
		}
	}

	return count, allergenToIngredients

}

func part2(allergenToIngredients map[string][]string) string {

	allergenMapping := make(map[string]string)

	for {

		if len(allergenToIngredients) == 0 {
			break
		}

		for allergen, ingredients := range allergenToIngredients {
			if len(allergenToIngredients[allergen]) == 1 {
				allergenMapping[allergen] = ingredients[0]
				for k, v := range allergenToIngredients {
					allergenToIngredients[k] = util.RemoveStringByValue(ingredients[0], v)
				}
				delete(allergenToIngredients, allergen)
				break
			}
		}
	}

	sortedKeys := []string{}
	for k := range allergenMapping {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	sortedIngredients := []string{}
	for _, k := range sortedKeys {
		sortedIngredients = append(sortedIngredients, allergenMapping[k])
	}

	return strings.Join(sortedIngredients, ",")

}

// https://github.com/juliangruber/go-intersect/blob/master/intersect.go#L43
func intersection(a []string, b []string) []string {
	set := make([]string, 0)
	hash := make(map[string]bool)

	for i := 0; i < len(a); i++ {
		hash[a[i]] = true
	}

	for i := 0; i < len(b); i++ {
		if _, found := hash[b[i]]; found {
			set = append(set, b[i])
		}
	}

	return set
}

func main() {

	rawInput := util.FetchInputForDay("2020", "21")
	parsedInput := util.ParseInputByLine(rawInput)
	fmt.Println("Done parsing input.")
	fmt.Println()

	// PART 1
	s := time.Now()
	ans1, allergenIngredients := part1(parsedInput)
	t := time.Now()
	e := t.Sub(s)
	fmt.Println("Answer for first question: ", ans1)
	fmt.Println("First answer retrieved in: ", e)
	fmt.Println()

	s = time.Now()
	ans2 := part2(allergenIngredients)
	t = time.Now()
	e = t.Sub(s)
	fmt.Println("Answer for second question: ", ans2)
	fmt.Println("Second answer retrieved in: ", e)

}
