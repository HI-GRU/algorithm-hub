package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	n      int
	w      int
	dp     [][]int
	pts    [][]int
	memo   [][][]int
	resArr [][]int
)

const INF = 3_000_000

func read() {
	fmt.Scan(&n, &w)
	reader := bufio.NewReader(os.Stdin)
	dp = make([][]int, w+1)
	pts = make([][]int, w+1)
	memo = make([][][]int, w+1)
	resArr = make([][]int, w+1)

	for i := 1; i <= w; i++ {
		pts[i] = make([]int, 2)
		resArr[i] = make([]int, 2)

		line, _ := reader.ReadString('\n')
		fields := strings.Fields(line)

		for j, v := range fields {
			pts[i][j], _ = strconv.Atoi(v)
		}
	}

	for i := 0; i <= w; i++ {
		dp[i] = make([]int, w+1)
		memo[i] = make([][]int, w+1)

		for j := 0; j <= w; j++ {
			dp[i][j] = INF
			memo[i][j] = make([]int, 2)
		}
	}

	pts[0] = make([]int, 2)
	dp[0][0] = 0
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	read()

	for i := 0; i < w; i++ {
		for j := 0; j < w; j++ {
			if i == j && i != 0 {
				continue
			}

			if dp[i][j] == INF {
				continue
			}

			next := max(i, j) + 1

			k := dp[i][j] + distance(i, next, 1)
			if dp[next][j] > k {
				dp[next][j] = k
				memo[next][j] = []int{i, j}
			}
			k = dp[i][j] + distance(j, next, 2)
			if dp[i][next] > k {
				dp[i][next] = k
				memo[i][next] = []int{i, j}
			}
		}
	}

	ans := INF

	for i := 0; i < w; i++ {
		if ans > dp[i][w] {
			ans = dp[i][w]
			resArr[w] = []int{i, w}
		}

		if ans > dp[w][i] {
			ans = dp[w][i]
			resArr[w] = []int{w, i}
		}
	}

	for i := w; i >= 1; i-- {
		resArr[i-1] = memo[resArr[i][0]][resArr[i][1]]
	}

	writer.WriteString(strconv.Itoa(ans) + "\n")
	for i := 1; i <= w; i++ {

		if resArr[i][0] == i {
			writer.WriteString("1\n")
		} else {
			writer.WriteString("2\n")
		}
	}
}

func distance(a, b, t int) int {

	x1 := pts[a][0]
	y1 := pts[a][1]
	x2 := pts[b][0]
	y2 := pts[b][1]

	if t == 1 {
		if a == 0 {
			x1 = 1
			y1 = 1
		} else if b == 0 {
			x2 = 1
			y2 = 1
		}
	}

	if t == 2 {
		if a == 0 {
			x1 = n
			y1 = n
		} else if b == 0 {
			x2 = n
			y2 = n
		}
	}

	return int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
}
