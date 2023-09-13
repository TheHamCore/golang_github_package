package golang_github_package

func ReadMap(data map[string]int) []int {
	result := make([]int, 0)

	for _, value := range data {
		result = append(result, value)
	}

	return result
}
