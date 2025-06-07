package main  

//Question
// Consider a rat placed at (0,0) in a square matrix of order N * N. It has to reach the destination at (N-1, N - 1). 
// Find all possible paths that the rat can take to reach from source to destination. 
// The directions in which the rat can move are 'U' (up), 'D'(down), 'L' (left), 'R' (right).
//  Value 0 at a cell in the matrix represents that it is blocked and rat cannot move to it while value 1 at a cell in 
// the matrix represents that rat can be travel through it.

// Note: In a path, no cell can be visited more than one time.

func findPath(maze [][]int, n int) []string {
	ans :=[] string {}
	vis := make([][]bool,n)
	for i := range vis{
		vis[i]= make([]bool, n)
	}

	dir := "DLRU"
	di := []int{1,0,0,-1}
	dj := []int{0,-1,1,0}

	if maze[0][0] == 1{
		solvee(0, 0, maze, n, &ans, "", vis, di, dj, dir)
	}
	return ans
}
func solvee(i, j int, maze [][]int, n int, ans *[]string, move string, vis [][]bool, di, dj []int, dir string) {
	if i == n-1 && j == n-1{
		*ans = append(*ans,move)
		return
	}

	for ind := 0;ind<4;ind++{
		nexti := i + di[ind]
		nextj := j + dj[ind]

		if nexti >= 0 && nextj >= 0 && nexti < n && nextj < n && !vis[nexti][nextj] && maze[nexti][nextj] == 1 {
			vis[i][j] = true
			solvee(nexti,nextj,maze,n,ans,move+string(dir[ind]),vis,di,dj,dir)
			vis[i][j] = false
		}
	}
}