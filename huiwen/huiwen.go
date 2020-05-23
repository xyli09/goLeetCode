package huiwen

import "fmt"

func max(nums ...int) int {
	m := nums[0]
	for _, n := range nums {
		if m < n {
			m = n
		}
	}
	return m
}

//Palindrome 回文

func LongestPalindromeSubseq(s string) int {
	n := len(s)

	if n < 2 {
		return n
	}

	dp := [1000][1000]int{}
	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}

	}

	return dp[0][n-1]
}


func TestHuiWen(){
	s :="ejaajeq"
	l:=LongestPalindromeSubseq(s)
	fmt.Println(l)
}