package main

import (
	"fmt"
	"github.com/42races/tree/bstree"
	"time"
)

var GridSize = 5 // for printing purpose

type Grid struct {
	cells *bstree.BStree
}

// key logic won't work for gridsize > 9
func (g *Grid) newKey(x, y int) int {
	key := y
	for y != 0 {
		x = x * 10
		y = y / 10
	}

	key = key + x
	return key
}

func (g *Grid) makeAlive(x, y int) {
	key := g.newKey(x, y)
	bstree.Insert(g.cells, key, true)
}

func (g *Grid) kill(x, y int) {
	key := g.newKey(x, y)
	bstree.Delete(g.cells, key)
}

func (g *Grid) display() {
	bstree.Display(g.cells)
}

func (g *Grid) init() {
	g.cells = bstree.New()
}

func (g *Grid) get(x, y int) bool {
	_, ok, _ := bstree.Get(g.cells, g.newKey(x, y))
	return ok
}

func (g *Grid) printGrid() {
	for i := 0; i < GridSize; i++ {
		for j := 0; j < GridSize; j++ {
			ok := g.get(i, j)
			if ok {
				fmt.Print(" 1 ")
			} else {
				fmt.Print(" 0 ")
			}
		}
		fmt.Println("")
	}
}

func (g *Grid) getNeighbourerCount(x, y int) int {
	var c int
	for i := x - 1; (i <= (x + 1)) && (i >= 0); i++ {
		for j := y - 1; (j <= (y + 1)) && (j >= 0); j++ {
			if (x == i) && (y == j) {
				continue
			}
			if ok := g.get(i, j); ok {
				c++
			}
		}
	}

	return c
}

func (g *Grid) tick() Grid {
	var grid Grid
	grid.init()

	for i := 0; i < GridSize; i++ {
		for j := 0; j < GridSize; j++ {
			c := g.getNeighbourerCount(i, j)
			ok := g.get(i, j)
			if ((c < 2) || (c > 3)) && ok {
				grid.kill(i, j)
			} else if (c == 3) && !ok {
				grid.makeAlive(i, j)
			} else {
				if ok {
					grid.makeAlive(i, j)
				}
			}
		}
	}
	return grid
}

func (g *Grid) count() int {
	return g.cells.Count
}

func main() {
	var grid Grid
	grid.init()
	grid.makeAlive(1, 2)
	grid.makeAlive(2, 2)
	grid.makeAlive(3, 2)

	fmt.Println("Initial alive cell count:", grid.count())

	for i := 0; ; i++ {
		grid.printGrid()
		grid = grid.tick()
		time.Sleep(time.Second)
		fmt.Println("********** tick", i, " ********** alive cells:", grid.count())
	}
}
