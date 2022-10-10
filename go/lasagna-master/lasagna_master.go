package lasagna

// gramsOfNoodles represents the amount of noodles in grams for each noodle layer.
const gramsOfNoodles = 50

// litersOfSauce represents the amount of sauce in liters for each layer of souce.
const litersOfSauce = 0.2

// PreparationTime returns the estimated time to prepare the lasagna based on the layers and timePerLayer.
func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = 2
	}
	return len(layers) * timePerLayer
}

// Quantities returns the needed quantities of sauce and noodles based on the layers in the lasagna.
func Quantities(layers []string) (amountOfNoodles int, amountOfSauce float64) {
	for _, layer := range layers {
		if layer == "sauce" {
			amountOfSauce += litersOfSauce
		}
		if layer == "noodles" {
			amountOfNoodles += gramsOfNoodles
		}
	}
	return
}

// AddSecretIngredient updates your recipe list with the secret ingredient from your friend's recipe.
func AddSecretIngredient(friendsIngredients []string, myIngredients []string) {
	myIngredients[len(myIngredients)-1] = friendsIngredients[len(friendsIngredients)-1]
}

// ScaleRecipe returns a list of ingredient amounts for the desired number of portions
// based on the ingredient amount for 2 portions.
func ScaleRecipe(amtTwoPortions []float64, numberOfPortions int) []float64 {
	properPortions := make([]float64, len(amtTwoPortions))
	for i, amount := range amtTwoPortions {
		properPortions[i] = amount * float64(numberOfPortions) / 2
	}
	return properPortions
}
