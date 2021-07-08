/*

Remove Black Pixels

Given a square matrix with 0 and 1 values where 0 means white pixel and 1 black pixel.
Design an algorithm to remove all the black pixels where they are not connected to the
border matrix pixels. Connected means that 2 or more pixels are horizontally or vertically
adjacent, no matter diagonally.

Example
Input:
[ 1, 0, 1, 0, 0 ]
[ 0, 1, 0, 1, 1 ]
[ 0, 0, 0, 1, 0 ]
[ 0, 1, 0, 1, 0 ]
[ 0, 0, 0, 0, 1 ]

Pixels to remove
[ 1, 0, 1, 0, 0 ] The pixels with * are not connected so they are removed.
[ 0, 1*, 0, 1', 1 ] The three pixels with ' are connected with a border pixel
[ 0, 0, 0, 1', 0 ] they can not be removed.
[ 0, 1*, 0, 1', 0 ]
[ 0, 0, 0, 0, 1 ]

Output:
[ 1, 0, 1, 0, 0 ]
[ 0, 0, 0, 1, 1 ]
[ 0, 0, 0, 1, 0 ]
[ 0, 0, 0, 1, 0 ]
[ 0, 0, 0, 0, 1 ]

*/

package main

import (
	"strconv"
)

func contains(res map[string]bool, i int, j int) bool {
	_, found := res[strconv.Itoa(i)+strconv.Itoa(j)]
	return found
}

func removeBlackPixels(a [][]int) [][]int {
	n := len(a)    // number of rows
	m := len(a[0]) // number of columns

	res := make(map[string]bool) // map with coors as key and value the bool if hold the pixel or not

	// go through the matrix to get the edges and with value 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if a[i][j] == 1 && (i == 0 || i == n-1 || j == 0 || j == m-1) {
				// Call my recursive function
				move(i, j, a, res, n, m)
			}
		}
	}

	// go through the matrix except the borders to create new matrix with black pixels removed
	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			if a[i][j] == 1 && !contains(res, i, j) {
				a[i][j] = 0
			}
		}
	}

	return a
}

func move(x int, y int, a [][]int, res map[string]bool, n int, m int) {
	// left
	if x-1 > 0 { // avoid overflow
		if a[x-1][y] == 1 && !contains(res, x-1, y) {
			res[strconv.Itoa(x-1)+strconv.Itoa(y)] = true
			move(x-1, y, a, res, n, m)
		}
	}
	// right
	if x+1 < n-1 { // avoid overflow
		if a[x+1][y] == 1 && !contains(res, x+1, y) {
			res[strconv.Itoa(x+1)+strconv.Itoa(y)] = true
			move(x+1, y, a, res, n, m)
		}
	}
	// up
	if y-1 > 0 { // avoid overflow
		if a[x][y-1] == 1 && !contains(res, x, y-1) {
			res[strconv.Itoa(x)+strconv.Itoa(y-1)] = true
			move(x, y-1, a, res, n, m)
		}
	}
	// abajo
	if y+1 < m-1 { // avoid overflow
		if a[x][y+1] == 1 && !contains(res, x, y+1) {
			res[strconv.Itoa(x)+strconv.Itoa(y+1)] = true
			move(x, y+1, a, res, n, m)
		}
	}
}

func main() {
}
