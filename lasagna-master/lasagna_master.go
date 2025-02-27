package lasagna

func PreparationTime(layers []string, avgTime int) int {
	if avgTime == 0 {
		avgTime = 2
	}
	return len(layers) * avgTime
}

func Quantities(layers []string) (int, float64) {
	nBase := 50
	sBase := 0.2
	nCount := 0
	sCount := 0
	for _, v := range layers {
		switch v {
		case "noodles":
			nCount++
		case "sauce":
			sCount++
		}
	}
	return nCount * nBase, float64(sCount) * sBase
}

func AddSecretIngredient(input, inputN []string) []string {
	recipe := make(map[string]bool)
	for _, v := range inputN {
		recipe[v] = true
	}
	for _, v := range input {
		if !recipe[v] {
			inputN[len(inputN)-1] = v
		}
	}

	return inputN
}

func ScaleRecipe(recipe []float64, portions int) []float64 {
	var scaledQuant []float64
	for _, v := range recipe {
		scaledQuant = append(scaledQuant, v/2.0*float64(portions))
	}
	return scaledQuant
}

// Your first steps could be to read through the tasks, and create
// these fuNctions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
