package leetcode

import "fmt"

func PacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 {
		return [][]int{}
	}

	// store cells that can flow both side
	var res [][]int

	// cache to store if certain cell can flow all the way or not
	canFlow := make([][]bool, len(heights))
	for i := range canFlow {
		canFlow[i] = make([]bool, len(heights[0]))
	}

	// loop over each cell
	for r := range heights {
		for c := range heights[r] {

			var dfs func(int, int) bool
			dfs = func(r, c int) bool {
				if canFlow[r][c] {
					return true
				}

				// if at border return true
				if r <= 0 || c <= 0 || r >= len(heights)-1 || c >= len(heights[r])-2 {
					return true
				}

				canFlow[r][c] = (dfs(r-1, c) || dfs(r, c-1)) && (dfs(r, c+1) || dfs(r+1, c))
				return true
			}

			if dfs(r, c) {
				res = append(res, []int{r, c})
			}
		}
	}

	fmt.Println(canFlow)
	return res
}
