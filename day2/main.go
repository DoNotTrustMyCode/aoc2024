package main
import (
	"fmt"
	// "sort"
	// "math",
	// "bufio"
	// "os"
	// "strconv"
	// "strings"
)

func isSafe(levels []int) bool {
	
	var increased bool = false;
	var decreased bool = false;

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		if (diff > 0) {
			increased = true
		} else if (diff < 0) {
			decreased = true;
		}
		if (increased && decreased) {return false}
		
		if (diff < 0) { diff = -diff};
		if (diff < 1 || diff > 3) {return false}
    }
	return true;
}


func main() {

	levels1 := []int{7,6,4,2,1}
	levels2 := []int{1,2,7,8,9}
	levels3 := []int{9,7,6,2,1}
	levels4 := []int{1,3,2,4,5}
	levels5 := []int{8,6,4,4,1}
	levels6 := []int{1,3,6,7,9}

	fmt.Println(isSafe(levels1));
	fmt.Println(isSafe(levels2));
	fmt.Println(isSafe(levels3));
	fmt.Println(isSafe(levels4));
	fmt.Println(isSafe(levels5));
	fmt.Println(isSafe(levels6));
}
