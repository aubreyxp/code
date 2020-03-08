package main

import "fmt"

// 岛屿数量问题
func main() {
	fmt.Println("vim-go")

	var a = [][]byte{
		/*
			{1, 1, 1, 1, 0},
			{1, 1, 0, 1, 0},
			{1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0},
		*/

		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Printf("%d ", a[i][j])
		}
		fmt.Printf("\n")
	}

	count := numIslands(a)
	fmt.Println("count:", count)
}

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				count++
				dfs(grid, i, j)
			}
		}
	}

	return count
}

func dfs(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) {
		return
	}

	if j < 0 || j >= len(grid[i]) {
		return
	}

	grid[i][j] += 1

	if (i-1) >= 0 && (i-1) < len(grid) {
		if grid[i-1][j] == 1 {
			grid[i-1][j] += 1
			dfs(grid, i-1, j)
		}
	}

	if (i+1) >= 0 && (i+1) < len(grid) {
		if grid[i+1][j] == 1 {
			grid[i+1][j] += 1
			dfs(grid, i+1, j)
		}
	}

	if (j+1) >= 0 && (j+1) < len(grid[i]) {
		if grid[i][j+1] == 1 {
			grid[i][j+1] += 1
			dfs(grid, i, j+1)
		}
	}

	if (j-1) >= 0 && (j-1) < len(grid[i]) {
		if grid[i][j-1] == 1 {
			grid[i][j-1] += 1
			dfs(grid, i, j-1)
		}
	}
}
