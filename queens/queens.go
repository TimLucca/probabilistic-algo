package queens

import (
	"math/rand"
	"time"
)

type queen struct {
	x int32
	y int32
}

func solve(queens []queen, col int32) []queen {
	for i := range queens {
		if queens[i].x == col {
			if col == 7 {
				return queens
			}
			return solve(queens, col+1)
		}
	}
	for i := 0; i < 8; i++ {
		valid := true
		for j := range queens {
			if queens[j].y == int32(i) || queens[j].y == (queens[j].x-col)+int32(i) || queens[j].y == -(queens[j].x-col)+int32(i) {
				valid = false
			}
		}
		if valid && col == 7 {
			return append(queens, queen{x: col, y: int32(i)})
		} else if valid {
			p := solve(append(queens, queen{x: col, y: int32(i)}), col+1)
			if p != nil {
				return p
			}
		}
	}
	return nil
}

func RQueens(k int) (bool, int64) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	start := time.Now()
	var queens []queen
	for i := 0; i < k; i++ {
		j := 0
		for {
			// Limiter for random queen placement
			if j > 1000 {
				return false, 0
			}
			valid := true
			tmp := queen{
				x: r.Int31n(8),
				y: r.Int31n(8),
			}
			for j := 0; j < i; j++ {
				//fmt.Printf("Checking (%d, %d) vs (%d, %d)\n", tmp.x, tmp.y, queens[j].x, queens[j].y)
				if queens[j].x == tmp.x || queens[j].y == tmp.y || queens[j].y == (queens[j].x-tmp.x)+tmp.y || queens[j].y == -(queens[j].x-tmp.x)+tmp.y {
					valid = false
					break
				}
			}
			if valid {
				// fmt.Printf("Queen added at (%d, %d)\n", tmp.x, tmp.y)
				queens = append(queens, tmp)
				break
			}
			j++
		}
	}

	queens = solve(queens, 0)
	elapsed := time.Since(start)
	if queens == nil {
		return false, 0
	}
	return true, elapsed.Nanoseconds()
	// var board [8][8]string
	// for i := 0; i < 8; i++ {
	// 	for j := 0; j < 8; j++ {
	// 		board[i][j] = "O"
	// 	}
	// }
	// for _, q := range queens {
	// 	board[q.x][q.y] = "X"
	// }

	// for _, row := range board {
	// 	fmt.Println(row)
	// }

}
