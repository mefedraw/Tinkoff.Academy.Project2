package DataStructures

import "errors"

type UnionFind struct {
	parent map[int]int
}

func NewUnionFind() *UnionFind {
	return &UnionFind{
		parent: make(map[int]int),
	}
}

func (uf *UnionFind) Add(x int) error {
	if _, exists := uf.parent[x]; exists {
		return errors.New("element already exists")
	}
	uf.parent[x] = x
	return nil
}

func (uf *UnionFind) Find(x int) (int, error) {
	if _, exists := uf.parent[x]; !exists {
		return -1, errors.New("element does not exist")
	}
	if uf.parent[x] != x {
		uf.parent[x], _ = uf.Find(uf.parent[x])
	}
	return uf.parent[x], nil
}

func (uf *UnionFind) Union(x, y int) error {
	rootX, err := uf.Find(x)
	if err != nil {
		return err
	}
	rootY, err := uf.Find(y)
	if err != nil {
		return err
	}
	if rootX == rootY {
		return nil
	}
	uf.parent[rootY] = rootX
	return nil
}
