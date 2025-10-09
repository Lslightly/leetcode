package main

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */
// @lc code=start

var parenCache map[int][]string = make(map[int][]string)

func generateParenthesis(n int) []string {
	if strs, ok := parenCache[n]; ok {
		return strs
	}
	if n == 0 {
		tmp := []string{""}
		parenCache[n] = tmp
		return tmp
	}
	if n == 1 {
		tmp := []string{"()"}
		parenCache[n] = tmp
		return tmp
	}
	/*
		(n-1)
		(n-2) 1
		(n-3) 2
		...
		() n-1
	*/
	res := make([]string, 0)
	for i := range n {
		leftStrs := generateParenthesis(n - 1 - i)
		rightStrs := generateParenthesis(i)
		for _, left := range leftStrs {
			for _, right := range rightStrs {
				res = append(res, "("+left+")"+right)
			}
		}
	}
	parenCache[n] = res
	return res
}

// @lc code=end
