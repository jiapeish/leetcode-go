package main

// cut every square to 4 parts, and mark them
//      \top/
//       \0/
// left 3 X 1 right
//       /2\
//     /down\

func regionsBySlashes(grid []string) int {
	n := len(grid)
	u := newUnion(4*n*n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			baseIdx := 4*(n*i+j)
			top := baseIdx + 0
			right := baseIdx + 1
			down := baseIdx + 2
			left := baseIdx + 3

			switch grid[i][j] {
			case '\\':
				u.Union(top, right)
				u.Union(left, down)
			case '/':
				u.Union(top, left)
				u.Union(right, down)
			case ' ':
				u.Union(top, right)
				u.Union(right, down)
				u.Union(down, left)
			}

			// right union with right square's left
			if j < n - 1{
				rsl := baseIdx + 4 + 3
				u.Union(right, rsl)
			}

			// down union with down square's top
			if i < n - 1{
				dst := baseIdx + 4*n
				u.Union(down, dst)
			}
		}
	}

	return u.count
}



type Union struct {
	parent []int
	rank []int
	count int
}

func newUnion(size int) *Union {
	parent := make([]int, size)
	rank := make([]int, size)
	for i := range parent {
		parent[i] = i
	}

	return &Union{
		parent: parent,
		rank: rank,
		count:  size,
	}
}

func (u *Union) Find1(i int) int {
	if u.parent[i] != i {
		u.parent[i] = u.Find1(u.parent[i])
	}
	return u.parent[i]
}

func (u *Union) Find(i int) int {
	root := i
	for u.parent[root] != root {
		root = u.parent[root]
	}

	for u.parent[i] != i {
		tmp := u.parent[i]
		u.parent[i] = root
		i = tmp
	}
	return root
}


func (u *Union) Union(i, j int) {
	iRoot := u.Find(i)
	jRoot := u.Find(j)
	if iRoot == jRoot {
		return
	}

	if u.rank[iRoot] < u.rank[jRoot] {
		u.parent[iRoot] = jRoot
	} else if u.rank[iRoot]	> u.rank[jRoot] {
		u.parent[jRoot] = iRoot
	} else {
		u.parent[jRoot] = iRoot
		u.rank[iRoot] = u.rank[iRoot] + 1
	}
	u.count--
}