package main 

// isSafe checks if it's safe to color the node with 'col'
func isSafe(node int, G [][]int, color []int, col int) bool {
	for _, it := range G[node] {
		if color[it] == col {
			return false
		}
	}
	return true
}

func solve(node int,G [][]int,color []int,n int,m int) bool{
	if node == n{
		return true
	}

	for i := 1;i<=m;i++{
		if isSafe(node,G,color,i){
			color[node] = i
			if solve(node+1,G,color,n,m){
				return true
			}

			color[node]=0
		}
	}
	return false
}

func graphColoring(G[][]int,m int) bool{
	n := len(G)
	color := make([]int,n)
	return solve(0,G,color,n,m)

}