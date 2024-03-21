package problems

func move(nums []int, target, start int) {
	l := len(nums)
	for i := start; i < l; i++ {
		nums[target] = nums[i]
		target++
		nums[i] = 99999
	}
}

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l < 2 {
		return l
	}
	n := nums[0]
	nIndex := 1
	res := 1
	next := 1

	for next < l {
		nowN := nums[next]
		if nowN == 99999 {
			break
		}
		if nowN == n {
			if nIndex < 2 {
				nIndex += 1
				next += 1
				res += 1
			} else {
				target := next
				for next < l && nowN == n {
					next += 1
				}
				move(nums, target, next)
			}
		} else {
			n = nowN
			nIndex = 1
			res += 1
			next += 1
		}
	}
	return res
}
