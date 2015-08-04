package main

import "fmt"

var currentId int

var meals Meals

// Give us some seed data
func init() {
	RepoCreateMeal(Meal{Name: "Lochmuir Salmon"})
	RepoCreateMeal(Meal{Name: "Chicken"})
	RepoCreateMeal(Meal{Name: "Beef"})
	RepoCreateMeal(Meal{Name: "Melba Toast"})
	RepoCreateMeal(Meal{Name: "Rice"})
	RepoCreateMeal(Meal{Name: "Pasta"})
	RepoCreateMeal(Meal{Name: "Carrot Cake"})
	RepoCreateMeal(Meal{Name: "Digestive Biscuit"})
	RepoCreateMeal(Meal{Name: "Carrot"})
}

func RepoFindMeal(id int) Meal {
	for _, t := range meals {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return Meal{}
}

func RepoCreateMeal(t Meal) Meal {
	currentId += 1
	t.Id = currentId
	meals = append(meals, t)
	return t
}

func RepoDestroyMeal(id int) error {
	for i, t := range meals {
		if t.Id == id {
			meals = append(meals[:i], meals[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Meals with id of %d to delete", id)
}
