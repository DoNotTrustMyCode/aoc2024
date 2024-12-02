package main
import (
	"fmt"
	"sort"
	// "math",
	"bufio"
	"os"
	"strconv"
	"strings"
)

func absInt(n int) int {
    if n < 0 {
        return -n
    }
    return n
}
func countOccurrences (value int, list []int) int {
	var count int = 0;
	for index := range list {
		if (list[index] == value) {
			count+=1;
		}		

	}
	return count
	
}
func createSlicesFromFile(filename string) ([]int, []int, error) { 

	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var slice1, slice2 []int

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		// Split the line into two numbers
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid line format: %s", line)
		}

		// Convert strings to integers
		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number: %s", parts[0])
		}
		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number: %s", parts[1])
		}

		// Append to slices
		slice1 = append(slice1, num1)
		slice2 = append(slice2, num2)
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return slice1, slice2, nil
}
func solve(listA []int, listB []int) int {

	sort.Ints(listA)
	sort.Ints(listB)
	sum:=0;
	for index := range listA {
		a := listA[index]
		b := listB[index]

		var diff = absInt(a-b);
		sum += diff;
		

	}
	return sum;
}

func solveB(listA []int, listB []int) int {
	var res int = 0;
	for index := range listA {
		occ := countOccurrences (listA[index], listB);
		res += listA[index] * occ;
		

	}
	return res
}
func main(){
	slice1, slice2, err := createSlicesFromFile("day1/input");
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	res:= solve(slice1, slice2);

	fmt.Println(res)
	fmt.Println(solveB(slice1, slice2))
	
}

