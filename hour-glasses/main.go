package hour_glasses

// Complete the hourglassSum function below.
// https://www.hackerrank.com/challenges/2d-array/problem
func hourglassSum(arr [][]int32) int32 {

	rows := len(arr)
	columns := len(arr[0])

	var maxSum int32 = 0
	for i := 0; i < rows-2; i++ {
		for j := 0; j < columns-2; j++ {
			a1 := arr[i][j]
			a2 := arr[i][j+1]
			a3 := arr[i][j+2]

			b2 := arr[i+1][j+1]

			c1 := arr[i+2][j]
			c2 := arr[i+2][j+1]
			c3 := arr[i+2][j+2]

			sum := a1 + a2 + a3 + b2 + c1 + c2 + c3

			if sum > maxSum {
				maxSum = sum
			}
		}
	}

	return maxSum

}
