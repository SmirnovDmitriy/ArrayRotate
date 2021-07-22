package Smirnov

import "fmt"

func ArrayRotate(arr []int, count int) []int {

	var num int
	for j := 0; j < count; j++ {
		//fmt.Println(len(arr) - 1)
		for i := (len(arr) - 1); i >= 0; i-- {
			if i == (len(arr) - 1) {
				num = arr[i]
			}

			if i > 0 {
				arr[i] = arr[i-1]
			}
			if i == 0 {
				arr[i] = num
			}
		}

	}
	return arr
}

func main() {
	fmt.Println(ArrayRotate([]int{1, 2, 3, 4, 5}, 2))
}
