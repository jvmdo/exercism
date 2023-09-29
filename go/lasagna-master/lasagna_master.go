package lasagna

func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		return 2 * len(layers)
	}

	return timePerLayer * len(layers)
}

func Quantities(layers []string) (noodles int, sauce float64) {
	for _, ing := range layers {
		if (ing == "noodles") {
			noodles += 50
		}

		if (ing == "sauce") {
			sauce += 0.2
		}
	}

	return 
}

func AddSecretIngredient(friendList, ownList []string) {
	ownList[len(ownList) - 1] = friendList[len(friendList) - 1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	scale := make([]float64, len(quantities))

	copy(scale, quantities)

	for i := range scale {
		scale[i] *= float64(portions) / 2.0
	}

	return scale
}
