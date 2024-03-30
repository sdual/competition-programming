package chap02

var a []int
var n, k int

func dfs(i int, sum int) bool {
	if i == n {
		return sum == k
	}
	if dfs(i+1, sum) {
		return true
	}

	if dfs(i+1, sum+a[i]) {
		return true
	}
	return false
}
