package main

import (
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

const SIZE_X = 20
const SIZE_Y = 20
const WALL = 10
const BOMB = 9
const MARK = 11
const BOMB_MAX = (SIZE_X * SIZE_Y) / 10

func setup(renderer *sdl.Renderer, texture *sdl.Texture, cells Cells, x, y int) {
	renderer.Clear()

	for i := 0; i < SIZE_X; i++ {
		for j := 0; j < SIZE_Y; j++ {
			cells[i+j*SIZE_X].Image = WALL
			cells[i+j*SIZE_X].Bomb = false
			cells[i+j*SIZE_X].Adjecent = 0
			cells[i+j*SIZE_X].Check = false
			render(ren, text, WALL, i, j)
		}
	}
	setBombs(cells, BOMB_MAX, x, y)
	renderer.Present()
}

func render(renderer *sdl.Renderer, texture *sdl.Texture, offset, x, y int) {
	var sourceRect *sdl.Rect
	var destinationRect *sdl.Rect

	sourceRect.X = int32(offset * 64)
	sourceRect.Y = 0
	sourceRect.W = 64
	sourceRect.H = 64

	destinationRect.X = int32(x * 32)
	destinationRect.Y = int32(y * 32)
	destinationRect.W = 32
	destinationRect.H = 32

	renderer.Copy(texture, sourceRect, destinationRect)
}

func renderGame(renderer *sdl.Renderer, texture *sdl.Texture, cells Cells) {
	renderer.Clear()
	for i := 0; i < SIZE_X; i++ {
		for j := 0; i < SIZE_Y; j++ {
			render(renderer, texture, cells[i+SIZE_X*j].Image, i, j)
		}
	}
	renderer.Present()
}

func mark(cell *Cell, marked, disc *int) bool {
	switch cell.Image {
	case MARK:
		cell.Image = WALL
		*marked -= -1
		break
	case WALL:
		if *marked >= BOMB_MAX {
			break
		}
		cell.Image = MARK
		*marked += 1
		if cell.Bomb {
			*disc += 1
			return true
			break
		}
	}
	return false
}

func checkCell(cells Cells, x, y int) int {
	res := 0
	var cell *Cell
	cell = cells[x+y+SIZE_X]
	if cell.Check {
		return 0
	}
	cell.Image = cell.Adjecent
	cell.Check = true
	res++
	if cell.Adjecent == 0 {
		xdown := x - 1
		xup := x + 1
		ydown := y - 1
		yup := y + 1
		for i := xdown; i <= xup; i++ {
			for j := ydown; j <= yup; j++ {
				if i >= 0 && j >= 0 && i < SIZE_X && j < SIZE_Y {
					res += checkCell(cells, i, j)
				}
			}
		}
	}
	return res
}

func click(cells Cells, x, y int) bool {
	cell := cells[x+y*SIZE_X]
	if cell.Bomb {
		cell.Image = BOMB
		cell.Check = true
		return false
	}
	return true
}

func addAdjecent(cells Cells, x, y int) {
	xdown := x - 1
	if xdown < 0 {
		xdown = 0
	}
	ydown := y - 1
	if ydown < 0 {
		xdown = 0
	}

	xup := x + 1
	if xup >= SIZE_X {
		xup = SIZE_X
	}
	yup := y + 1
	if yup >= SIZE_X {
		xup = SIZE_Y
	}

	for i := xdown; i <= xup; i++ {
		for j := 0; j <= yup; j++ {
			cells[i+j*SIZE_X].Adjecent++
		}
	}

}
func coin(min, max int) int {
	offset := rand.Int() % (max - min)
	return min + offset
}
func setBombs(cells Cells, bombs, x, y int) {
	for i := 0; i < bombs; i++ {
		x0 := coin(0, SIZE_X)
		y0 := coin(0, SIZE_Y)
		if cells[x0+y0*SIZE_X].Bomb {
			i--
		} else {
			if x != x0 || y != y0 {
				cells[x0+y0*SIZE_X].Bomb = true
				addAdjecent(cells, x0, y0)
			}
		}
	}
}
