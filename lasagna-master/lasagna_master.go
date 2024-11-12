package lasagna

const (
	defaultTimePerLayer = 2
	noodlesQtyPerLayer  = 50
	sauceQtyPerLayer    = 0.2
)

func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = defaultTimePerLayer
	}
	return len(layers) * timePerLayer
}

func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for i := 0; i < len(layers); i++ {
		switch layers[i] {
		case "sauce":
			sauce += sauceQtyPerLayer
		case "noodles":
			noodles += noodlesQtyPerLayer
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, numberOfPortions int) []float64 {
	result := make([]float64, len(quantities))
	for i := 0; i < len(quantities); i++ {
		result[i] = quantities[i] * float64(numberOfPortions) / 2.0
	}
	return result
}
