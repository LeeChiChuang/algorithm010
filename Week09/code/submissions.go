package str

func lengthOfLastWord(s string) int {
	ans := make([]string, 0)
	left, right := 0, 0
	for idx, char := range s {
		right++
		if char == ' ' && s[left] != 0 {
			ans = append(ans, s[left:right])
			left = right
		}
		if (idx == len(s)-1) && char != ' ' {
			ans = append(ans, s[left:right])
		}
	}

	return len(ans[len(ans)-1])
}