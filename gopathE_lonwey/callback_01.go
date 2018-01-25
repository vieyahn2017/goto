package main
import(
    "fmt"
)

func main() {
	s := make([]int, 10)
	fmt.Println(s)
	for i, _ := range s {
		s[i] = i
	}
	fmt.Println(s)
	fmt.Println(sum(s, isDouble))
	fmt.Println(sum(s, isSingle))
}

func isDouble(a int) bool {
	if a%2 == 0 {
		return true
	} else {
		return false
	}
}

func isSingle(a int) bool {
	if a%2 == 0 {
		return false
	} else {
		return true
	}
}

func sum(s []int, f func(int) bool) int {
	sum := 0
	for _, v := range s {
		if f(v) {
			sum += v
		}
	}
	return sum
}