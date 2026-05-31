package main

import "fmt"

const NMAX int = 1000

type Recipe struct {
	name             string
	category         string
	ingredients      [99]string
	steps            string
	countIngredients int
	cookingTime      int
	searchCount      int
}

type recipeList [NMAX]Recipe

func main() {
	var recipes recipeList
	var choice int
	var amount int

	for {
		fmt.Println("Culinary Recipe Management and Search Application (MyRecipe)")
		fmt.Println("1. Add recipe")
		fmt.Println("2. Edit recipe")
		fmt.Println("3. Delete recipe")
		fmt.Println("4. Search recipe")
		fmt.Println("5. Display all recipes")
		fmt.Println("6. View Statistics")
		fmt.Println("7. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addRecipe(&recipes, &amount)
		case 2:
			editRecipe(&recipes, amount)
		case 3:
			deleteRecipe(&recipes, &amount)
		case 4:
			searchRecipe(&recipes, amount)
		case 5:
			displayRecipes(&recipes, amount)
		case 6:
			viewStatistics(&recipes, amount)
		case 7:
			fmt.Println("Exiting the application. Goodbye! 😘😘😘")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func addRecipe(recipes *recipeList, n *int) {
	if *n >= NMAX {
		fmt.Println("Recipe list is full. Cannot add more recipes.")
		return
	}
	fmt.Print("Enter recipe name: ")
	fmt.Scan(&recipes[*n].name)
	fmt.Print("Enter category: ")
	fmt.Scan(&recipes[*n].category)
	fmt.Print("Enter cooking time (in minutes): ")
	fmt.Scan(&recipes[*n].cookingTime)
	fmt.Print("Enter number of ingredients: ")
	fmt.Scan(&recipes[*n].countIngredients)
	for i := 0; i < recipes[*n].countIngredients; i++ {
		fmt.Printf("Ingredient %d: ", i+1)
		fmt.Scan(&recipes[*n].ingredients[i])
	}

	fmt.Print("Cooking steps: ")
	fmt.Scan(&recipes[*n].steps)
	recipes[*n].searchCount = 0
	(*n)++
	fmt.Println("Recipe added successfully!")
}

func editRecipe(recipes *recipeList, n int) {
	var title string
	var choice int
	var ingredientNumber int

	fmt.Print("Enter the name of the recipe to edit: ")
	fmt.Scan(&title)
	index := findIndexRecipe(&recipes, n, title)
	if index == -1 {
		fmt.Println("Recipe not found.")
		return
	}

	for {
		fmt.Println("Edit recipe: ", recipes[index].name)
		fmt.Println("1. Change title")
		fmt.Println("2. Change ingredient")
		fmt.Println("3. Change cooking time")
		fmt.Println("4. Change cooking steps")
		fmt.Println("5. Done")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter new title: ")
			fmt.Scan(&recipes[index].name)
			fmt.Println("Title changed.")
		case 2:
			if recipes[index].countIngredients == 0 {
				fmt.Println("This recipe has no ingredients.")
			} else {
				fmt.Println("Current ingredients:")
				for i := 0; i < recipes[index].countIngredients; i++ {
					fmt.Printf("%d. %s\n", i+1, recipes[index].ingredients[i])
				}
				fmt.Print("Enter ingredient number to change: ")
				fmt.Scan(&ingredientNumber)
				if ingredientNumber < 1 || ingredientNumber > recipes[index].countIngredients {
					fmt.Println("Invalid ingredient number.")
				} else {
					fmt.Print("Enter the new ingredient: ")
					fmt.Scan(&recipes[index].ingredients[ingredientNumber-1])
					fmt.Println("Ingredient updated.")
				}
			}
		case 3:
			fmt.Print("Enter new cooking time (in minutes): ")
			fmt.Scan(&recipes[index].cookingTime)
			fmt.Println("Cooking time updated.")
		case 4:
			fmt.Print("Enter new cooking steps: ")
			fmt.Scan(&recipes[index].steps)
			fmt.Println("Cooking steps updated.")
		case 5:
			fmt.Println("Finished editing recipe.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func findIndexRecipe(recipes *recipeList, n int, title string) int {
	for i := 0; i < n; i++ {
		if recipes[i].name == title {
			return i
		}
	}
	return -1
}
func deleteRecipe(recipes *recipeList, n *int) {

}
func searchRecipe(recipes recipeList, n int) {

}
func displayRecipes(recipes *recipeList, n int) {

}
func viewStatistics(recipes recipeList, n int) {

}
