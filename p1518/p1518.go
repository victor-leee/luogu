package main

import (
	"fmt"
)

const (
	cellEmpty    byte = '.'
	cellObstacle byte = '*'
	cellCow      byte = 'C'
	cellFarmer   byte = 'F'
)

type direction int

func (d *direction) turn() {
	*d = (*d + 1) % 4
}

func (d *direction) move(curX, curY int) (int, int) {
	x, y := 0, 0
	if *d == 0 {
		x--
	}
	if *d == 1 {
		y++
	}
	if *d == 2 {
		x++
	}
	if *d == 3 {
		y--
	}

	return x + curX, y + curY
}

type boardMap [10][10]byte

func (b *boardMap) isObstacle(x, y int) bool {
	return b[x][y] == cellObstacle
}

func (b *boardMap) isWithinRange(x, y int) bool {
	return 0 <= x && x < len(b) && 0 <= y && y < len(b[0])
}

func (b *boardMap) canMove(x, y int) bool {
	if x > 0 && b[x-1][y] != cellObstacle {
		return true
	}
	if x < len(b) - 1 && b[x+1][y] != cellObstacle {
		return true
	}
	if y > 0 && b[x][y-1] != cellObstacle {
		return true
	}
	if y < len(b[0]) - 1 && b[x][y+1] != cellObstacle {
		return true
	}

	return false
}

func main() {
	var board boardMap
	var xFarmer, yFarmer, xCow, yCow int
	var dirFarmer, dirCow direction
	var time int
	state := make(map[int]struct{})

	for i := 0; i < 10; i++ {
		var row string
		fmt.Scanf("%s\n", &row)
		for j := 0; j < 10; j++ {
			board[i][j] = row[j]
			if board[i][j] == cellCow {
				xCow, yCow = i, j
			}
			if board[i][j] == cellFarmer {
				xFarmer, yFarmer = i, j
			}
		}
	}

	if !board.canMove(xCow, yCow) || !board.canMove(xFarmer, yFarmer) {
		fmt.Println(0)
		return
	}

	for {
		if isStateExist(state, flat(xCow, yCow), dirCow, flat(xFarmer, yFarmer), dirFarmer) {
			fmt.Println(0)
			return
		}

		nxFarmer, nyFarmer := dirFarmer.move(xFarmer, yFarmer)
		if !board.isWithinRange(nxFarmer, nyFarmer) || board.isObstacle(nxFarmer, nyFarmer) {
			dirFarmer.turn()
		} else {
			setState(state, flat(xCow, yCow), dirCow, flat(xFarmer, yFarmer), dirFarmer)
			xFarmer, yFarmer = nxFarmer, nyFarmer
		}

		nxCow, nyCow := dirCow.move(xCow, yCow)
		if !board.isWithinRange(nxCow, nyCow) || board.isObstacle(nxCow, nyCow) {
			dirCow.turn()
		} else {
			setState(state, flat(xCow, yCow), dirCow, flat(xFarmer, yFarmer), dirFarmer)
			xCow, yCow = nxCow, nyCow
		}

		time++
		if xCow == xFarmer && yCow == yFarmer {
			fmt.Println(time)
			return
		}
	}
}

func flat(x, y int) int {
	return x * 10 + y
}

func isStateExist(state map[int]struct{}, posCow int, cowDir direction, posFarmer int, farmerDir direction) bool{
	cow := int(cowDir) * 100 + posCow
	farmer := int(farmerDir) * 100 + posFarmer
	stateNum := cow * 1000 + farmer
	_, ok := state[stateNum]

	return ok
}

func setState(state map[int]struct{}, posCow int, cowDir direction, posFarmer int, farmerDir direction) {
	cow := int(cowDir) * 100 + posCow
	farmer := int(farmerDir) * 100 + posFarmer
	stateNum := cow * 1000 + farmer
	state[stateNum] = struct{}{}
}