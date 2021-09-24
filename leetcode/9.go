package leetcode

func isPalindrome(x int) bool {
	reverseHalfX := 0
	if x < 0 || x%10 == 0 {
		return false
	}

	for x > reverseHalfX {
		digit := x % 10
		reverseHalfX = reverseHalfX*10 + digit
		if x == reverseHalfX {
			return true
		}

		x = x / 10

	}

	if x == reverseHalfX {
		return true
	}

	return false
}
