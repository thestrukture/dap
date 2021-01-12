package methods

func GetTotal(traits map[string]int) int {

	var result = 0

	for _, v := range traits {

		result += v
	}

	return result
}
