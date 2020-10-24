package Week_07

type UnionFind struct {
	count  int
	parent []int
}

func NewUnionFind(grid [][]byte) *UnionFind {
	ret := new(UnionFind)
	ret.count = 0
	m := len(grid)
	n := len(grid[0])
	parent := make([]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				parent[i*n+j] = i*n + j
				ret.count++
			}
		}
	}
	ret.parent = parent
	return ret
}

func (uf *UnionFind) union(elem1, elem2 int) {
	parent1 := uf.find(elem1)
	parent2 := uf.find(elem2)
	if parent1 != parent2 {
		uf.parent[parent1] = parent2
		uf.count--
	}
}

func (uf *UnionFind) find(elem int) int {
	if uf.parent[elem] != elem {
		uf.parent[elem] = uf.find(uf.parent[elem])
	}
	return uf.parent[elem]
}

func (uf *UnionFind) getCount() int {
	return uf.count
}

func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	row := len(grid)
	col := len(grid[0])
	uf := NewUnionFind(grid)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				grid[i][j] = '0'
				cur := i*col + j
				if i+1 < row && grid[i+1][j] == '1' {
					uf.union(cur, (i+1)*col+j)
				}
				if j+1 < col && grid[i][j+1] == '1' {
					uf.union(cur, i*col+j+1)
				}
			}
		}
	}

	return uf.getCount()
}
