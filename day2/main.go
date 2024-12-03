package main
import (
	"fmt"
	// "sort"
	// "math",
	"bufio"
	"os"
	"strconv"
	"strings"
)

func loadInput(filename string) ([][]int, error) { 

	file, err := os.Open(filename)
	if err != nil {
		return  nil, err
	}
	defer file.Close()

	var arr [][]int

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var row []int
		line := scanner.Text()
		// Split the line into two numbers
		parts := strings.Fields(line)
		
		for index := range parts {
			num, err := strconv.Atoi(parts[index])
			row = append(row, num)
			if err != nil {
				return nil,  fmt.Errorf("invalid number: %s", parts[0])
			}

		}

		arr =append(arr, row)
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return arr, nil
}
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
func removeElement(slice []int, i int) []int {
    if i < 0 || i >= len(slice) {
        fmt.Println("Index out of range")
        return slice // Return the original slice if index is invalid
    }
	res := make([]int, 0, len(slice)-1)
	res = append(res, slice[:i]...)
    res = append(res, slice[i+1:]...)
	return res
}

func isConditionallySafe(slice []int) bool {
	fmt.Println("slice: ", slice)	
	for index := range(slice) {
		var partial = removeElement(slice, index) 
		fmt.Println("partial: ", index, partial)	
		if isSafe(partial) {return true}
	}
	return false

}
func solveA ()  {
	arr, err := loadInput("day2/input");
	
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var res int = 0;
	for _, row := range arr {
		if isSafe(row) {
			res+=1;
		}
	}
	fmt.Println(res)
}
func solveB ()  {
	arr, err := loadInput("day2/input");
	
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var res int = 0;
	for _, row := range arr {
		if isConditionallySafe(row) {
			res+=1;
		}
	}
	fmt.Println(res)



}
func main() {
	// solveA();
	solveB();
	// levels1 := []int{7,6,4,2,1}
	// levels2 := []int{1,2,7,8,9}
	// levels3 := []int{9,7,6,2,1}
	// levels4 := []int{1,3,2,4,5}
	// levels5 := []int{8,6,4,4,1}
	// levels6 := []int{1,3,6,7,9}
	// //
	// fmt.Println(isConditionallySafe(levels1));
	// fmt.Println(isConditionallySafe(levels2));
	// fmt.Println(isConditionallySafe(levels3));
	// fmt.Println(isConditionallySafe(levels4));
	// fmt.Println(isConditionallySafe(levels5));
	// fmt.Println(isConditionallySafe(levels6));
}
