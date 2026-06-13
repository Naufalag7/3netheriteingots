```markdown
# 🍳 ResepKu - Culinary Recipe Management App

A console-based culinary recipe management application built with Go. It manages food recipes, records ingredients and cooking steps, tracks the most searched menus, and demonstrates classic searching and sorting algorithms in a real-world use case.

## 👥 Member Group

* **Naufal Abdurrahman Gumelar** * **Role:** Lead Developer, Logic Architecture, and Code Review.
  * **Features:** Array Structuring, Binary/Sequential Search logic, Edit & Delete constraints, View Statistics logic.
* **Achmad Ramadhanu Putra Abdi**
  * **Role:** Co-Developer, UI/UX Console Formatting, and Algorithm Tester.
  * **Features:** Selection/Insertion Sort implementation, Menu Navigation loop, Display formatting, Documentation.

> *Created for the "Algoritma Pemrograman 2" Final Project at Telkom University.*

---

## 📖 Description

ResepKu is a data-driven application for managing and finding cooking ideas based on available ingredients. It stores recipe profiles, ingredient lists, and cooking steps in local memory, then processes them using fundamental sorting and searching algorithms.

**Target Users:**
* Cooking enthusiasts looking to save their personal recipes.
* General public who wants to look up meal ideas based on a specific main ingredient they currently have.

## ✨ Features

| Feature | Description |
| :--- | :--- |
| **Recipe Management** | View, add, delete, and update recipe data (Title, Category, Ingredients, Steps, and Time). |
| **Advanced Sorting** | Sort the displayed recipes either alphabetically (A-Z / Z-A) or by cooking duration (Fastest / Longest). |
| **Ingredient Search** | Find specific recipes based on their *Main Ingredient* using different search algorithms. |
| **Data Statistics** | View the total recipes available per category and a leaderboard of the most frequently searched recipes. |

---

## 🚀 Getting Started

### Requirements
* Go 1.20 or later installed on your machine.

### Run the Application
```bash
# Clone the repository
git clone [https://github.com/yourusername/resepku-golang-app.git](https://github.com/yourusername/resepku-golang-app.git)

# Navigate to the directory
cd resepku-golang-app

# Run the program
go run main.go

```

### Main Menu CLI

1. Add recipe
2. Edit recipe
3. Delete recipe
4. Search by ingredient
5. Display all recipes
6. View Statistics
7. Exit

---

## 📂 Project Structure

To comply with academic constraints, the app utilizes a modular monolithic approach inside a single file, divided into distinct functional blocks:

```text
resepku-golang-app/
├── main.go               # Entry point, loads the menu loop and data array
│   ├── (Structs)         # Data structures and array limits
│   ├── (CRUD Logic)      # Functions to add, edit, and delete recipes
│   ├── (Search Logic)    # Sequential & Binary search implementations
│   ├── (Sort Logic)      # Selection & Insertion sort implementations
│   └── (Display Logic)   # UI formatting, statistics, and sub-menus
└── README.md             # Project documentation

```

---

## 🗄️ Data Structures

### Recipe

Stores all information regarding a single cooking recipe.

```go
type Recipe struct {
	name             string      // Title of the recipe
	category         string      // e.g., Dessert, Soup, Main Course
	ingredients      [100]string // Array of ingredients (Index 0 is Main Ingredient)
	steps            [100]string // Array of cooking steps
	countIngredients int         // Total ingredients registered
	countSteps       int         // Total steps registered
	cookingTime      int         // Duration in minutes
	searchCount      int         // Tracker for statistics
}

```

### Recipe List

The main static database simulating memory storage.

```go
const NMAX int = 1000

type recipeList [NMAX]Recipe // Array of 1000 Recipe structs

```

---

## ⚙️ Algorithms

### Searching

| Algorithm | Used when | Searches by |
| --- | --- | --- |
| **Sequential Search** | Searching via the `[S]` menu option. | Main Ingredient (`ingredients[0]`) |
| **Binary Search** | Searching via the `[B]` menu option. Requires the array to be sorted first. | Main Ingredient (`ingredients[0]`) |
| **Binary Search** | Background process during Edit and Delete actions. | Exact Recipe Title (`name`) |

### Sorting

| Algorithm | Used for | Sorts by |
| --- | --- | --- |
| **Selection Sort** | `SortbyNameAscending` & `SortbyNameDescending` | Recipe Title (Alphabetical) |
| **Selection Sort** | `sortByMainIngredientAscending` | Main Ingredient (Prep for Binary Search) |
| **Insertion Sort** | `SortbyTimeAscending` & `SortbyTimeDescending` | Cooking Time (Integers) |

---

## 🔄 Application Flow

1. **Adding a Recipe (`addRecipe`)**
* Checks if the `NMAX` limit is reached.
* Validates duplicate names using a quick sequential sweep.
* Prompts user for details, saving `ingredients[0]` explicitly as the "Main Ingredient".


2. **Editing / Deleting (`editRecipe` / `deleteRecipe`)**
* Displays all available recipes.
* Sorts the array alphabetically automatically.
* Uses Binary Search to pinpoint the exact array index of the requested recipe title.
* Overwrites data (Edit) or shifts the array elements to the left (Delete).


3. **Searching (`search`)**
* Asks the user to input a main ingredient.
* **Sequential:** Sweeps from index 0 to `n`, incrementing the `searchCount` for every match.
* **Binary:** Sorts the array by main ingredient first. Halves the data to find a match, then sweeps slightly to the right to catch duplicate main ingredients across different recipes.


4. **Statistics (`viewStatistics`)**
* Iterates through the data to count categories.
* Duplicates the main array into a temporary `sorted` array.
* Sorts the temporary array based on `searchCount` using Selection Sort to display a leaderboard without altering the original database order.



---

## 📜 Functions Documentation

Every core function in the codebase is documented using block comments (`/* ... */`) explaining its exact purpose and behavior.

**Core Operations**

| Function | Purpose |
| --- | --- |
| `addRecipe` | Validates input and adds a new recipe into the array. |
| `editRecipe` | Retrieves a recipe by exact name and opens a sub-menu to modify specific fields. |
| `deleteRecipe` | Locates a recipe by name and shifts array elements to overwrite it. |
| `viewStatistics` | Tallies category data and ranks recipes based on search frequency. |
| `printRecipeDetails` | A reusable UI helper to print data cleanly with 50-character alignment. |

---

## ⚖️ License

This project was created for academic purposes.

```

```
