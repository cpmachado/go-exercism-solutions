package resistorcolorduo

// Colors returns the list of all colors.
func Colors() []string {
	return []string{
		"black",
		"brown",
		"red",
		"orange",
		"yellow",
		"green",
		"blue",
		"violet",
		"grey",
		"white",
	}
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	for i, v := range Colors() {
		if v == color {
			return i
		}
	}
	return -1
}

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	sum := 0
	for i, color := range colors {
		sum = sum*10 + ColorCode(color)
		if i == 1 {
			break
		}
	}
	return sum
}
