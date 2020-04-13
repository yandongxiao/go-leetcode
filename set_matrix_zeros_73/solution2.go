package set_matrix_zeros_73

const maxUint = ^uint(0)
const maxInt = int(maxUint >> 1)

func setZeroes(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		row := matrix[i]
		for j := 0; j < len(row); j++ {
			if row[j] == 0 {
				for x := 0; x < len(matrix); x++ {
					if x != i && matrix[x][j] != 0 {
						matrix[x][j] = maxInt
					}
				}

				doBreak := true
				for y := 0; y < len(row); y++ {
					if row[y] == 0 {
						if y != j {
							doBreak = false
						}
					} else {
						row[y] = maxInt
					}
				}
				if doBreak {
					break
				}
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		row := matrix[i]
		for j := 0; j < len(row); j++ {
			if row[j] == maxInt {
				row[j] = 0
			}
		}
	}
}
