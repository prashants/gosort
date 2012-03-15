package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"strconv"
	"container/list"
)

var data list.List

func bubble_sort() {
	last_element := data.Len()
	last_element -= 1
	for counter1 := last_element; counter1 >= 1; counter1 -= 1 {
		/* single pass through the list and swap elements */
		counter2 := counter1
		for e := data.Front(); counter2 >= 1; e = e.Next() {
			next := e.Next()
			if e.Value.(int) >= next.Value.(int) {
				/* swap current and next element */
				temp := e.Value
				e.Value = next.Value
				next.Value = temp
			}
			counter2 -= 1
		}
	}
}

func main() {
	fmt.Println("Welcome to Bubble Sort")

	data.Init()

	/* Read unsorted data from file and store it into array */
	f, err := os.Open("./data1000.txt")
	if err == nil {
		r := bufio.NewReader(f)
		for {
			line, _, err := r.ReadLine()
			if err == io.EOF {
				break
			}
			i, err := strconv.Atoi(string(line))
			data.PushBack(i)
		}
	} else {
		fmt.Println("Error reading data file")
		return
	}
	err = f.Close();
	fmt.Println("Read ", data.Len(), " integers")

	bubble_sort()

	fmt.Println("Sorted List: ")
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

