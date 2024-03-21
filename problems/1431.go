package problems

func kidsWithCandies(candies []int, extraCandies int) []bool {
	// get max of candies
	maxCandy := 0
	for _, candy := range candies {
		if candy > maxCandy {
			maxCandy = candy
		}
	}
	result := make([]bool, len(candies))
	for i, candy := range candies {
		if candy+extraCandies >= maxCandy {
			result[i] = true
		}
	}
	return result
}
