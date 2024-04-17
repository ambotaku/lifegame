package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Conway's "Game of Life" (cellular automata)
// exercise of Lesson 20 of book "Get Programming with Go"

const (
	width  int = 80 // tty columns
	height int = 25 // tty lines
)

// the cell's playground
// (alive cells become true, dead cells false)
type Universe [][]bool

func main() {
	a := NewUniverse()
	a.Seed()
	for {
		a.Show()
		a = a.Render()
		time.Sleep(time.Second / 2)
	}
}

// create new, empty playground
func NewUniverse() Universe {
	u := make(Universe, height)
	for i := range u {
		u[i] = make([]bool, width)
	}
	return u
}

// fill playground with random cells
func (u Universe) Seed() {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r := rand.Intn(100)
			if r < 25 {
				u[y][x] = false // dead
			} else {
				u[y][x] = true // alive
			}
		}
	}
}

// display playground on screen
func (u Universe) Show() {
	fmt.Print("\033c") // clear Xterm screen

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			c := 32 // utf8 space
			if u[y][x] {
				c = 0x2588 // utf8 block rune
			}
			fmt.Print(string(rune(c)))
		}
		fmt.Println()
	}
}

// get neighbor cell on wrapped around playground
func (u Universe) Alive(x, y int) bool {
	if x < 0 {
		x += width
	}
	if x >= width {
		x = 0
	}
	if y < 0 {
		y += height
	}
	if y >= height {
		y = 0
	}
	return u[y][x]
}

// test cell's 8 neighbors
func (u Universe) Neighbors(x, y int) bool {
	type offset struct {
		x, y int
	}
	offsets := []offset{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, -0}, {-1, 1}, {0, 1}, {1, 1}}

	count := 0
	// count alive neighbors
	for _, ofs := range offsets {
		if u.Alive(x+ofs.x, y+ofs.y) {
			count++
		}
	}

	// cell needs 2 or 3 neighbors to survive
	return count == 2 || count == 3
}

// render and return next grid
func (u Universe) Render() Universe {
	b := NewUniverse()

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if u.Neighbors(x, y) {
				b[y][x] = true
			} else {
				b[y][x] = false
			}
		}
	}
	return b
}
