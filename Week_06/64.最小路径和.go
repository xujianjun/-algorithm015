package Week_06

// dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
func minPathSum(grid [][]int) int {
	l := len(grid)
	if l < 1 {
		return 0
	}

	for i := 0; i < l; i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j > 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if i > 0 && j == 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else if i > 0 && j > 0 {
				grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
			}
		}
	}
	return grid[l-1][len(grid[l-1])-1]
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
