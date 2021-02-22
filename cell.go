package main

type Cell struct {
	Image    int
	Bomb     bool
	Adjecent int
	Check    bool
}

type Cells []*Cell
