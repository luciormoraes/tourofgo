// exercise slice

package main

//	"golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy, dy)
	for i := range result {
		result[i] = make([]uint8, dx, dx)
		for j := range result[i] {
			result[i][j] = uint8(i ^ j)
		}
	}

	return result
}

// func main() {
// 	// pic.Show(Pic)

// }
